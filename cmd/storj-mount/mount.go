// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

// usage: storj-sim sj://bucket_name ~/target-directory
//
// Currently storj-mount is using 'uplink' configuration

// +build linux darwin

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"storj.io/storj/internal/fpath"
	libuplink "storj.io/storj/lib/uplink"
	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/process"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/uplink"
)

// UplinkFlags test
type UplinkFlags struct {
	uplink.Config
}

var (
	cfg UplinkFlags

	rootCmd = &cobra.Command{
		Use:   "mount",
		Short: "Storj mount utility",
		Args:  cobra.OnlyValidArgs,
	}

	mountCmd = &cobra.Command{
		Use:   "run",
		Short: "to run mount utility",
		RunE:  mountBucket,
		Args:  cobra.OnlyValidArgs,
	}
)

func main() {
	process.Exec(rootCmd)
}

func init() {
	rootCmd.AddCommand(mountCmd)

	defaultConfDir := fpath.ApplicationDir("storj", "uplink")

	confDirParam := cfgstruct.FindConfigDirParam()
	if confDirParam != "" {
		defaultConfDir = confDirParam
	}

	var confDir string
	cfgstruct.SetupFlag(zap.L(), mountCmd, &confDir, "config-dir", defaultConfDir, "main directory for mount configuration")

	defaults := cfgstruct.DefaultsFlag(mountCmd)
	cfgstruct.Bind(mountCmd.Flags(), &cfg, defaults, cfgstruct.ConfDir(confDir))
}

func mountBucket(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		return fmt.Errorf("No bucket specified for mounting")
	}
	if len(args) == 1 {
		return fmt.Errorf("No destination specified")
	}

	ctx := process.Ctx(cmd)

	if err := process.InitMetrics(ctx, nil, ""); err != nil {
		zap.S().Error("Failed to initialize telemetry batcher: ", err)
	}

	src, err := fpath.New(args[0])
	if err != nil {
		return err
	}
	if src.IsLocal() {
		return fmt.Errorf("No bucket specified. Use format sj://bucket/")
	}

	access, err := useOrLoadEncryptionAccess(cfg.Enc.EncryptionKey, cfg.Enc.KeyFilepath)
	if err != nil {
		return err
	}

	project, bucket, err := cfg.GetProjectAndBucket(ctx, src.Bucket(), access)
	if err != nil {
		return fmt.Errorf("Unable to get bucket: %v", err)
	}
	defer closeProjectAndBucket(project, bucket)

	nfs := pathfs.NewPathNodeFs(newStorjFS(ctx, bucket), nil)
	conn := nodefs.NewFileSystemConnector(nfs.Root(), nil)

	// workaround to avoid async (unordered) reading
	mountOpts := fuse.MountOptions{MaxBackground: 1}
	server, err := fuse.NewServer(conn.RawFS(), args[1], &mountOpts)
	if err != nil {
		return fmt.Errorf("Mount failed: %v", err)
	}

	go func() {
		<-ctx.Done()

		if err := server.Unmount(); err != nil {
			fmt.Printf("Unmount failed: %v", err)
		}
	}()

	server.Serve()
	return nil
}

type storjFS struct {
	ctx          context.Context
	bucket       *libuplink.Bucket
	createdFiles map[string]*storjFile
	nodeFS       *pathfs.PathNodeFs
	pathfs.FileSystem
}

func newStorjFS(ctx context.Context, bucket *libuplink.Bucket) *storjFS {
	return &storjFS{
		ctx:          ctx,
		bucket:       bucket,
		createdFiles: make(map[string]*storjFile),
		FileSystem:   pathfs.NewDefaultFileSystem(),
	}
}

func (sf *storjFS) OnMount(nodeFS *pathfs.PathNodeFs) {
	sf.nodeFS = nodeFS
}

func (sf *storjFS) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	zap.S().Debug("GetAttr: ", name)

	if name == "" {
		return &fuse.Attr{Mode: fuse.S_IFDIR | 0755}, fuse.OK
	}

	// special case for just created files e.g. while coping into directory
	createdFile, ok := sf.createdFiles[name]
	if ok {
		attr := &fuse.Attr{}
		status := createdFile.GetAttr(attr)
		return attr, status
	}

	node := sf.nodeFS.Node(name)
	if node != nil && node.IsDir() {
		return &fuse.Attr{Mode: fuse.S_IFDIR | 0755}, fuse.OK
	}

	object, err := sf.bucket.OpenObject(sf.ctx, name)
	if err != nil && !storj.ErrObjectNotFound.Has(err) {
		return nil, fuse.EIO
	}

	// file not found so maybe it's a prefix/directory
	if err != nil {
		list, err := sf.bucket.ListObjects(sf.ctx, &storj.ListOptions{
			Prefix:    name,
			Direction: storj.After,
			Limit:     1,
		})
		if err != nil {
			return nil, fuse.EIO
		}

		// if exactly one element has this prefix then it's directory
		if len(list.Items) == 1 {
			return &fuse.Attr{Mode: fuse.S_IFDIR | 0755}, fuse.OK
		}

		return nil, fuse.ENOENT
	}

	return &fuse.Attr{
		Owner: *fuse.CurrentOwner(),
		Mode:  fuse.S_IFREG | 0644,
		Size:  uint64(object.Meta.Size),
		Mtime: uint64(object.Meta.Modified.Unix()),
	}, fuse.OK
}

func (sf *storjFS) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	zap.S().Debug("OpenDir: ", name)

	var entries []fuse.DirEntry
	err := sf.listObjects(sf.ctx, name, false, func(items []storj.Object) error {
		for _, item := range items {
			path := item.Path

			mode := fuse.S_IFREG
			if item.IsPrefix {
				path = strings.TrimSuffix(path, "/")
				mode = fuse.S_IFDIR
			}
			entries = append(entries, fuse.DirEntry{Name: path, Mode: uint32(mode)})
		}
		return nil
	})
	if err != nil {
		zap.S().Errorf("error during opening directory: %v", err)
		return nil, fuse.EIO
	}

	return entries, fuse.OK
}

func (sf *storjFS) Mkdir(name string, mode uint32, context *fuse.Context) fuse.Status {
	zap.S().Debug("Mkdir: ", name)

	emptyReader := bytes.NewReader([]byte{})
	err := sf.bucket.UploadObject(sf.ctx, name+"/", emptyReader, &libuplink.UploadOptions{
		ContentType: "application/directory",
	})
	if err != nil {
		return fuse.EIO
	}

	return fuse.OK
}

func (sf *storjFS) Rmdir(name string, context *fuse.Context) (code fuse.Status) {
	zap.S().Debug("Rmdir: ", name)

	err := sf.listObjects(sf.ctx, name, true, func(items []storj.Object) error {
		for _, item := range items {
			err := sf.bucket.DeleteObject(sf.ctx, storj.JoinPaths(name, item.Path))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		zap.S().Errorf("error during removing directory: %v", err)
		return fuse.EIO
	}

	return fuse.OK
}

func (sf *storjFS) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	zap.S().Debug("Open: ", name)
	return newStorjFile(sf.ctx, name, sf.bucket, false, sf), fuse.OK
}

func (sf *storjFS) Create(name string, flags uint32, mode uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	zap.S().Debug("Create: ", name)

	return sf.addCreatedFile(name, newStorjFile(sf.ctx, name, sf.bucket, true, sf)), fuse.OK
}

func (sf *storjFS) addCreatedFile(name string, file *storjFile) *storjFile {
	sf.createdFiles[name] = file
	return file
}

func (sf *storjFS) removeCreatedFile(name string) {
	delete(sf.createdFiles, name)
}

func (sf *storjFS) listObjects(ctx context.Context, name string, recursive bool, handler func([]storj.Object) error) error {
	startAfter := ""

	for {
		list, err := sf.bucket.ListObjects(sf.ctx, &storj.ListOptions{
			Direction: storj.After,
			Cursor:    startAfter,
			Prefix:    name,
			Recursive: recursive,
		})
		if err != nil {
			return err
		}

		err = handler(list.Items)
		if err != nil {
			return err
		}

		if !list.More {
			break
		}

		startAfter = list.Items[len(list.Items)-1].Path
	}

	return nil
}

func (sf *storjFS) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	zap.S().Debug("Unlink: ", name)

	err := sf.bucket.DeleteObject(sf.ctx, name)
	if err != nil {
		if storj.ErrObjectNotFound.Has(err) {
			return fuse.ENOENT
		}
		return fuse.EIO
	}

	return fuse.OK
}

type storjFile struct {
	ctx             context.Context
	bucket          *libuplink.Bucket
	created         bool
	name            string
	size            uint64
	mtime           uint64
	reader          io.ReadCloser
	writer          io.WriteCloser
	predictedOffset int64
	FS              *storjFS

	nodefs.File
}

func newStorjFile(ctx context.Context, name string, bucket *libuplink.Bucket, created bool, FS *storjFS) *storjFile {
	return &storjFile{
		name:    name,
		ctx:     ctx,
		bucket:  bucket,
		mtime:   uint64(time.Now().Unix()),
		created: created,
		FS:      FS,
		File:    nodefs.NewDefaultFile(),
	}
}

func (f *storjFile) GetAttr(attr *fuse.Attr) fuse.Status {
	zap.S().Debug("GetAttr file: ", f.name)

	if f.created {
		attr.Mode = fuse.S_IFREG | 0644
		if f.size != 0 {
			attr.Size = f.size
		}
		attr.Mtime = f.mtime
		return fuse.OK
	}
	return fuse.ENOSYS
}

func (f *storjFile) Read(buf []byte, off int64) (res fuse.ReadResult, code fuse.Status) {
	// Detect if offset was moved manually (e.g. stream rev/fwd)
	if off != f.predictedOffset {
		f.closeReader()
	}

	reader, err := f.getReader(off)
	if err != nil {
		if storj.ErrObjectNotFound.Has(err) {
			return nil, fuse.ENOENT
		}
		return nil, fuse.EIO
	}

	n, err := io.ReadFull(reader, buf)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, fuse.EIO
	}

	f.predictedOffset = off + int64(n)

	return fuse.ReadResultData(buf[:n]), fuse.OK
}

func (f *storjFile) Write(data []byte, off int64) (uint32, fuse.Status) {
	writer, err := f.getWriter(off)
	if err != nil {
		return 0, fuse.EIO
	}

	zap.S().Debug("Starts writting: ", f.name)

	written, err := writer.Write(data)
	if err != nil {
		return 0, fuse.EIO
	}

	f.size += uint64(written)

	zap.S().Debug("Stops writting: ", f.name)

	return uint32(written), fuse.OK
}

func (f *storjFile) getReader(off int64) (io.ReadCloser, error) {
	if f.reader == nil {
		object, err := f.bucket.OpenObject(f.ctx, f.name)
		if err != nil {
			return nil, err
		}
		reader, err := object.DownloadRange(f.ctx, off, object.Meta.Size)
		if err != nil {
			return nil, err
		}

		f.reader = reader
	}
	return f.reader, nil
}

func (f *storjFile) getWriter(off int64) (_ io.Writer, err error) {
	if off == 0 {
		f.size = 0
		f.closeWriter()

		f.writer, err = f.bucket.NewWriter(f.ctx, f.name, &libuplink.UploadOptions{})
	}
	return f.writer, err
}

func (f *storjFile) Flush() fuse.Status {
	zap.S().Debug("Flush: ", f.name)

	f.closeReader()
	f.closeWriter()
	return fuse.OK
}

func (f *storjFile) closeReader() {
	if f.reader != nil {
		closeErr := f.reader.Close()
		if closeErr != nil {
			zap.S().Errorf("error closing reader: %v", closeErr)
		}
		f.reader = nil
	}
}

func (f *storjFile) closeWriter() {
	if f.writer != nil {
		closeErr := f.writer.Close()
		if closeErr != nil {
			zap.S().Errorf("error closing writer: %v", closeErr)
		}

		f.FS.removeCreatedFile(f.name)
		f.writer = nil
	}
}

func convertError(err error, path fpath.FPath) error {
	if storj.ErrBucketNotFound.Has(err) {
		return fmt.Errorf("Bucket not found: %s", path.Bucket())
	}

	if storj.ErrObjectNotFound.Has(err) {
		return fmt.Errorf("Object not found: %s", path.String())
	}

	return err
}

// GetProjectAndBucket returns a *libuplink.Bucket for interacting with a specific project's bucket
func (c *UplinkFlags) GetProjectAndBucket(ctx context.Context, bucketName string, access libuplink.EncryptionAccess) (project *libuplink.Project, bucket *libuplink.Bucket, err error) {
	project, err = c.GetProject(ctx)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		if err != nil {
			if err := project.Close(); err != nil {
				fmt.Printf("error closing project: %+v\n", err)
			}
		}
	}()

	bucket, err = project.OpenBucket(ctx, bucketName, &access)
	if err != nil {
		return nil, nil, err
	}

	return project, bucket, nil
}

// GetProject returns a *libuplink.Project for interacting with a specific project
func (c *UplinkFlags) GetProject(ctx context.Context) (*libuplink.Project, error) {
	apiKey, err := libuplink.ParseAPIKey(c.Client.APIKey)
	if err != nil {
		return nil, err
	}

	satelliteAddr := c.Client.SatelliteAddr

	cfg := &libuplink.Config{}

	cfg.Volatile.TLS = struct {
		SkipPeerCAWhitelist bool
		PeerCAWhitelistPath string
	}{
		SkipPeerCAWhitelist: !c.TLS.UsePeerCAWhitelist,
		PeerCAWhitelistPath: c.TLS.PeerCAWhitelistPath,
	}

	cfg.Volatile.MaxInlineSize = c.Client.MaxInlineSize
	cfg.Volatile.MaxMemory = c.RS.MaxBufferMem

	uplk, err := libuplink.NewUplink(ctx, cfg)
	if err != nil {
		return nil, err
	}

	opts := &libuplink.ProjectOptions{}

	encKey, err := uplink.UseOrLoadEncryptionKey(c.Enc.EncryptionKey, c.Enc.KeyFilepath)
	if err != nil {
		return nil, err
	}

	opts.Volatile.EncryptionKey = encKey

	project, err := uplk.OpenProject(ctx, satelliteAddr, apiKey, opts)

	if err != nil {
		if err := uplk.Close(); err != nil {
			fmt.Printf("error closing uplink: %+v\n", err)
		}
	}

	return project, err
}

func closeProjectAndBucket(project *libuplink.Project, bucket *libuplink.Bucket) {
	if err := bucket.Close(); err != nil {
		fmt.Printf("error closing bucket: %+v\n", err)
	}

	if err := project.Close(); err != nil {
		fmt.Printf("error closing project: %+v\n", err)
	}
}

// loadEncryptionAccess loads the encryption key stored in the file pointed by
// filepath and creates an EncryptionAccess with it.
func loadEncryptionAccess(filepath string) (libuplink.EncryptionAccess, error) {
	key, err := uplink.LoadEncryptionKey(filepath)
	if err != nil {
		return libuplink.EncryptionAccess{}, err
	}

	return libuplink.EncryptionAccess{
		Key: *key,
	}, nil
}

// useOrLoadEncryptionAccess creates an encryption key from humanReadableKey
// when it isn't empty otherwise try to load the key from the file pointed by
// filepath and creates an EnryptionAccess with it.
func useOrLoadEncryptionAccess(humanReadableKey string, filepath string) (libuplink.EncryptionAccess, error) {
	if humanReadableKey != "" {
		key, err := storj.NewKey([]byte(humanReadableKey))
		if err != nil {
			return libuplink.EncryptionAccess{}, err
		}

		return libuplink.EncryptionAccess{
			Key: *key,
		}, nil
	}

	return loadEncryptionAccess(filepath)
}
