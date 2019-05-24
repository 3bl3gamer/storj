// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: uplink.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type IDVersion struct {
	Number               uint32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	NewPrivateKey        uint64   `protobuf:"varint,2,opt,name=new_private_key,json=newPrivateKey,proto3" json:"new_private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDVersion) Reset()         { *m = IDVersion{} }
func (m *IDVersion) String() string { return proto.CompactTextString(m) }
func (*IDVersion) ProtoMessage()    {}
func (*IDVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6cead23ed820d0c, []int{0}
}
func (m *IDVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDVersion.Unmarshal(m, b)
}
func (m *IDVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDVersion.Marshal(b, m, deterministic)
}
func (m *IDVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDVersion.Merge(m, src)
}
func (m *IDVersion) XXX_Size() int {
	return xxx_messageInfo_IDVersion.Size(m)
}
func (m *IDVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_IDVersion.DiscardUnknown(m)
}

var xxx_messageInfo_IDVersion proto.InternalMessageInfo

func (m *IDVersion) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *IDVersion) GetNewPrivateKey() uint64 {
	if m != nil {
		return m.NewPrivateKey
	}
	return 0
}

type TLSConfig struct {
	SkipPeerCaWhitelist  bool     `protobuf:"varint,1,opt,name=skip_peer_ca_whitelist,json=skipPeerCaWhitelist,proto3" json:"skip_peer_ca_whitelist,omitempty"`
	PeerCaWhitelistPath  string   `protobuf:"bytes,2,opt,name=peer_ca_whitelist_path,json=peerCaWhitelistPath,proto3" json:"peer_ca_whitelist_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSConfig) Reset()         { *m = TLSConfig{} }
func (m *TLSConfig) String() string { return proto.CompactTextString(m) }
func (*TLSConfig) ProtoMessage()    {}
func (*TLSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6cead23ed820d0c, []int{1}
}
func (m *TLSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSConfig.Unmarshal(m, b)
}
func (m *TLSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSConfig.Marshal(b, m, deterministic)
}
func (m *TLSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSConfig.Merge(m, src)
}
func (m *TLSConfig) XXX_Size() int {
	return xxx_messageInfo_TLSConfig.Size(m)
}
func (m *TLSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TLSConfig proto.InternalMessageInfo

func (m *TLSConfig) GetSkipPeerCaWhitelist() bool {
	if m != nil {
		return m.SkipPeerCaWhitelist
	}
	return false
}

func (m *TLSConfig) GetPeerCaWhitelistPath() string {
	if m != nil {
		return m.PeerCaWhitelistPath
	}
	return ""
}

type UplinkConfig struct {
	Tls                  *TLSConfig `protobuf:"bytes,1,opt,name=tls,proto3" json:"tls,omitempty"`
	IdentityVersion      *IDVersion `protobuf:"bytes,2,opt,name=identity_version,json=identityVersion,proto3" json:"identity_version,omitempty"`
	PeerIdVersion        string     `protobuf:"bytes,3,opt,name=peer_id_version,json=peerIdVersion,proto3" json:"peer_id_version,omitempty"`
	MaxInlineSize        int64      `protobuf:"varint,4,opt,name=max_inline_size,json=maxInlineSize,proto3" json:"max_inline_size,omitempty"`
	MaxMemory            int64      `protobuf:"varint,5,opt,name=max_memory,json=maxMemory,proto3" json:"max_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UplinkConfig) Reset()         { *m = UplinkConfig{} }
func (m *UplinkConfig) String() string { return proto.CompactTextString(m) }
func (*UplinkConfig) ProtoMessage()    {}
func (*UplinkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6cead23ed820d0c, []int{2}
}
func (m *UplinkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkConfig.Unmarshal(m, b)
}
func (m *UplinkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkConfig.Marshal(b, m, deterministic)
}
func (m *UplinkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkConfig.Merge(m, src)
}
func (m *UplinkConfig) XXX_Size() int {
	return xxx_messageInfo_UplinkConfig.Size(m)
}
func (m *UplinkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkConfig proto.InternalMessageInfo

func (m *UplinkConfig) GetTls() *TLSConfig {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *UplinkConfig) GetIdentityVersion() *IDVersion {
	if m != nil {
		return m.IdentityVersion
	}
	return nil
}

func (m *UplinkConfig) GetPeerIdVersion() string {
	if m != nil {
		return m.PeerIdVersion
	}
	return ""
}

func (m *UplinkConfig) GetMaxInlineSize() int64 {
	if m != nil {
		return m.MaxInlineSize
	}
	return 0
}

func (m *UplinkConfig) GetMaxMemory() int64 {
	if m != nil {
		return m.MaxMemory
	}
	return 0
}

type EncryptionParameters struct {
	CipherSuite          uint32   `protobuf:"varint,1,opt,name=cipher_suite,json=cipherSuite,proto3" json:"cipher_suite,omitempty"`
	BlockSize            int32    `protobuf:"varint,2,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptionParameters) Reset()         { *m = EncryptionParameters{} }
func (m *EncryptionParameters) String() string { return proto.CompactTextString(m) }
func (*EncryptionParameters) ProtoMessage()    {}
func (*EncryptionParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6cead23ed820d0c, []int{3}
}
func (m *EncryptionParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptionParameters.Unmarshal(m, b)
}
func (m *EncryptionParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptionParameters.Marshal(b, m, deterministic)
}
func (m *EncryptionParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionParameters.Merge(m, src)
}
func (m *EncryptionParameters) XXX_Size() int {
	return xxx_messageInfo_EncryptionParameters.Size(m)
}
func (m *EncryptionParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionParameters.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionParameters proto.InternalMessageInfo

func (m *EncryptionParameters) GetCipherSuite() uint32 {
	if m != nil {
		return m.CipherSuite
	}
	return 0
}

func (m *EncryptionParameters) GetBlockSize() int32 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

type RedundancyScheme struct {
	Algorithm            uint32   `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	ShareSize            int32    `protobuf:"varint,2,opt,name=share_size,json=shareSize,proto3" json:"share_size,omitempty"`
	RequiredShares       int32    `protobuf:"varint,3,opt,name=required_shares,json=requiredShares,proto3" json:"required_shares,omitempty"`
	RepairShares         int32    `protobuf:"varint,4,opt,name=repair_shares,json=repairShares,proto3" json:"repair_shares,omitempty"`
	OptimalShares        int32    `protobuf:"varint,5,opt,name=optimal_shares,json=optimalShares,proto3" json:"optimal_shares,omitempty"`
	TotalShares          int32    `protobuf:"varint,6,opt,name=total_shares,json=totalShares,proto3" json:"total_shares,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedundancyScheme) Reset()         { *m = RedundancyScheme{} }
func (m *RedundancyScheme) String() string { return proto.CompactTextString(m) }
func (*RedundancyScheme) ProtoMessage()    {}
func (*RedundancyScheme) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6cead23ed820d0c, []int{4}
}
func (m *RedundancyScheme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedundancyScheme.Unmarshal(m, b)
}
func (m *RedundancyScheme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedundancyScheme.Marshal(b, m, deterministic)
}
func (m *RedundancyScheme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedundancyScheme.Merge(m, src)
}
func (m *RedundancyScheme) XXX_Size() int {
	return xxx_messageInfo_RedundancyScheme.Size(m)
}
func (m *RedundancyScheme) XXX_DiscardUnknown() {
	xxx_messageInfo_RedundancyScheme.DiscardUnknown(m)
}

var xxx_messageInfo_RedundancyScheme proto.InternalMessageInfo

func (m *RedundancyScheme) GetAlgorithm() uint32 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *RedundancyScheme) GetShareSize() int32 {
	if m != nil {
		return m.ShareSize
	}
	return 0
}

func (m *RedundancyScheme) GetRequiredShares() int32 {
	if m != nil {
		return m.RequiredShares
	}
	return 0
}

func (m *RedundancyScheme) GetRepairShares() int32 {
	if m != nil {
		return m.RepairShares
	}
	return 0
}

func (m *RedundancyScheme) GetOptimalShares() int32 {
	if m != nil {
		return m.OptimalShares
	}
	return 0
}

func (m *RedundancyScheme) GetTotalShares() int32 {
	if m != nil {
		return m.TotalShares
	}
	return 0
}

type BucketConfig struct {
	EncryptionParameters *EncryptionParameters `protobuf:"bytes,1,opt,name=encryption_parameters,json=encryptionParameters,proto3" json:"encryption_parameters,omitempty"`
	PathCipher           uint32                `protobuf:"varint,3,opt,name=path_cipher,json=pathCipher,proto3" json:"path_cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BucketConfig) Reset()         { *m = BucketConfig{} }
func (m *BucketConfig) String() string { return proto.CompactTextString(m) }
func (*BucketConfig) ProtoMessage()    {}
func (*BucketConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6cead23ed820d0c, []int{5}
}
func (m *BucketConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketConfig.Unmarshal(m, b)
}
func (m *BucketConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketConfig.Marshal(b, m, deterministic)
}
func (m *BucketConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketConfig.Merge(m, src)
}
func (m *BucketConfig) XXX_Size() int {
	return xxx_messageInfo_BucketConfig.Size(m)
}
func (m *BucketConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BucketConfig proto.InternalMessageInfo

func (m *BucketConfig) GetEncryptionParameters() *EncryptionParameters {
	if m != nil {
		return m.EncryptionParameters
	}
	return nil
}

func (m *BucketConfig) GetPathCipher() uint32 {
	if m != nil {
		return m.PathCipher
	}
	return 0
}

type Bucket struct {
	EncryptionParameters *EncryptionParameters `protobuf:"bytes,1,opt,name=encryption_parameters,json=encryptionParameters,proto3" json:"encryption_parameters,omitempty"`
	RedundancyScheme     *RedundancyScheme     `protobuf:"bytes,2,opt,name=redundancy_scheme,json=redundancyScheme,proto3" json:"redundancy_scheme,omitempty"`
	Name                 string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              int64                 `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	PathCipher           uint32                `protobuf:"varint,5,opt,name=path_cipher,json=pathCipher,proto3" json:"path_cipher,omitempty"`
	SegmentSize          int64                 `protobuf:"varint,6,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6cead23ed820d0c, []int{6}
}
func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetEncryptionParameters() *EncryptionParameters {
	if m != nil {
		return m.EncryptionParameters
	}
	return nil
}

func (m *Bucket) GetRedundancyScheme() *RedundancyScheme {
	if m != nil {
		return m.RedundancyScheme
	}
	return nil
}

func (m *Bucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bucket) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Bucket) GetPathCipher() uint32 {
	if m != nil {
		return m.PathCipher
	}
	return 0
}

func (m *Bucket) GetSegmentSize() int64 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func init() {
	proto.RegisterType((*IDVersion)(nil), "storj.libuplink.IDVersion")
	proto.RegisterType((*TLSConfig)(nil), "storj.libuplink.TLSConfig")
	proto.RegisterType((*UplinkConfig)(nil), "storj.libuplink.UplinkConfig")
	proto.RegisterType((*EncryptionParameters)(nil), "storj.libuplink.EncryptionParameters")
	proto.RegisterType((*RedundancyScheme)(nil), "storj.libuplink.RedundancyScheme")
	proto.RegisterType((*BucketConfig)(nil), "storj.libuplink.BucketConfig")
	proto.RegisterType((*Bucket)(nil), "storj.libuplink.Bucket")
}

func init() { proto.RegisterFile("uplink.proto", fileDescriptor_c6cead23ed820d0c) }

var fileDescriptor_c6cead23ed820d0c = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x65, 0x48, 0xf2, 0x7d, 0xb9, 0x89, 0x09, 0x35, 0x14, 0x45, 0x55, 0xab, 0x82, 0x2b,
	0x28, 0x8b, 0x2a, 0x8b, 0xf2, 0x06, 0x50, 0x16, 0x88, 0xb6, 0x8a, 0x9c, 0xfe, 0x13, 0x9b, 0xd1,
	0xc4, 0xbe, 0xc5, 0xd3, 0xd8, 0xe3, 0xe9, 0x78, 0x0c, 0x84, 0x57, 0xe8, 0x7b, 0xf4, 0xd5, 0xba,
	0xef, 0x13, 0x54, 0x73, 0x3d, 0x0e, 0x25, 0x64, 0xdb, 0x5d, 0xfc, 0x9b, 0x33, 0xf7, 0xde, 0x39,
	0xf7, 0x28, 0xd0, 0xaf, 0x54, 0x26, 0xe4, 0x6c, 0xa4, 0x74, 0x61, 0x8a, 0x60, 0x50, 0x9a, 0x42,
	0x7f, 0x1b, 0x65, 0x62, 0x5a, 0xe3, 0xf0, 0x1c, 0xba, 0x67, 0x6f, 0x3e, 0xa1, 0x2e, 0x45, 0x21,
	0x83, 0x1d, 0xe8, 0xc8, 0x2a, 0x9f, 0xa2, 0x1e, 0x7a, 0xbb, 0xde, 0xa1, 0x1f, 0xb9, 0xaf, 0xe0,
	0x00, 0x06, 0x12, 0xaf, 0x99, 0xd2, 0xe2, 0x8a, 0x1b, 0x64, 0x33, 0x9c, 0x0f, 0xd7, 0x76, 0xbd,
	0xc3, 0x56, 0xe4, 0x4b, 0xbc, 0x1e, 0xd7, 0xf4, 0x1c, 0xe7, 0x61, 0x05, 0xdd, 0x0f, 0x6f, 0x27,
	0x27, 0x85, 0xfc, 0x2a, 0x2e, 0x83, 0x23, 0xd8, 0x29, 0x67, 0x42, 0x31, 0x85, 0xa8, 0x59, 0xcc,
	0xd9, 0x75, 0x2a, 0x0c, 0x66, 0xa2, 0x34, 0x54, 0xfc, 0xff, 0x68, 0xcb, 0x9e, 0x8e, 0x11, 0xf5,
	0x09, 0xff, 0xdc, 0x1c, 0xd9, 0x4b, 0x0f, 0xf4, 0x4c, 0x71, 0x93, 0x52, 0xc3, 0x6e, 0xb4, 0xa5,
	0xee, 0x5f, 0x18, 0x73, 0x93, 0x86, 0xbf, 0x3d, 0xe8, 0x7f, 0xa4, 0xe7, 0xb8, 0xd6, 0xaf, 0x60,
	0xdd, 0x64, 0x25, 0xf5, 0xe9, 0xbd, 0x7e, 0x32, 0x5a, 0x7a, 0xf3, 0x68, 0x31, 0x63, 0x64, 0x65,
	0xc1, 0x29, 0x6c, 0x8a, 0x04, 0xa5, 0x11, 0x66, 0xce, 0xae, 0x6a, 0x27, 0xa8, 0xdb, 0xaa, 0xab,
	0x0b, 0xaf, 0xa2, 0x41, 0x73, 0xa7, 0x31, 0xef, 0x00, 0x06, 0x34, 0xba, 0x48, 0x16, 0x55, 0xd6,
	0x69, 0x66, 0xdf, 0xe2, 0xb3, 0xe4, 0x2f, 0x5d, 0xce, 0x6f, 0x98, 0x90, 0x99, 0x90, 0xc8, 0x4a,
	0x71, 0x8b, 0xc3, 0xd6, 0xae, 0x77, 0xb8, 0x1e, 0xf9, 0x39, 0xbf, 0x39, 0x23, 0x3a, 0x11, 0xb7,
	0x18, 0x3c, 0x03, 0xb0, 0xba, 0x1c, 0xf3, 0x42, 0xcf, 0x87, 0x6d, 0x92, 0x74, 0x73, 0x7e, 0xf3,
	0x8e, 0x40, 0xf8, 0x05, 0xb6, 0x4f, 0x65, 0xac, 0xe7, 0xca, 0x88, 0x42, 0x8e, 0xb9, 0xe6, 0x39,
	0x1a, 0xd4, 0x65, 0xb0, 0x07, 0xfd, 0x58, 0xa8, 0x14, 0x35, 0x2b, 0x2b, 0x61, 0xd0, 0x6d, 0xb2,
	0x57, 0xb3, 0x89, 0x45, 0xb6, 0xf2, 0x34, 0x2b, 0xe2, 0x59, 0xdd, 0xdc, 0x3e, 0xb5, 0x1d, 0x75,
	0x89, 0xd8, 0xc6, 0xe1, 0x2f, 0x0f, 0x36, 0x23, 0x4c, 0x2a, 0x99, 0x70, 0x19, 0xcf, 0x27, 0x71,
	0x8a, 0x39, 0x06, 0x4f, 0xa1, 0xcb, 0xb3, 0xcb, 0x42, 0x0b, 0x93, 0xe6, 0xae, 0xe6, 0x1d, 0xb0,
	0x15, 0xcb, 0x94, 0x6b, 0xbc, 0x57, 0x91, 0x08, 0x3d, 0xe5, 0x25, 0x0c, 0x34, 0x7e, 0xaf, 0x84,
	0xc6, 0x84, 0x11, 0x2d, 0xc9, 0x9a, 0x76, 0xb4, 0xd1, 0xe0, 0x09, 0xd1, 0xe0, 0x05, 0xf8, 0x1a,
	0x15, 0x17, 0xba, 0x91, 0xb5, 0x48, 0xd6, 0xaf, 0xa1, 0x13, 0xed, 0xc3, 0x46, 0xa1, 0x8c, 0xc8,
	0x79, 0xd6, 0xa8, 0xda, 0xa4, 0xf2, 0x1d, 0x75, 0xb2, 0x3d, 0xe8, 0x9b, 0xc2, 0xdc, 0x89, 0x3a,
	0x24, 0xea, 0x11, 0xab, 0x25, 0xe1, 0x0f, 0x0f, 0xfa, 0xc7, 0x55, 0x3c, 0x43, 0xe3, 0x82, 0x73,
	0x01, 0x8f, 0x71, 0x61, 0x2a, 0x53, 0x0b, 0x57, 0x5d, 0x94, 0xf6, 0x1f, 0xe4, 0x61, 0xd5, 0x0a,
	0xa2, 0x6d, 0x5c, 0xb5, 0x98, 0xe7, 0xd0, 0xb3, 0x41, 0x66, 0xf5, 0x26, 0xc8, 0x00, 0x3f, 0x02,
	0x8b, 0x4e, 0x88, 0x84, 0x3f, 0xd7, 0xa0, 0x53, 0x4f, 0xf3, 0x4f, 0xe7, 0x78, 0x0f, 0x8f, 0xf4,
	0x62, 0xbb, 0xac, 0xa4, 0xf5, 0xba, 0xbc, 0xef, 0x3d, 0xa8, 0xbb, 0x9c, 0x83, 0x68, 0x53, 0x2f,
	0x27, 0x23, 0x80, 0x96, 0xe4, 0x39, 0xba, 0xb0, 0xd3, 0xef, 0x60, 0x08, 0xff, 0xc5, 0x1a, 0xb9,
	0xc1, 0xc4, 0x65, 0xbb, 0xf9, 0x5c, 0x76, 0xa1, 0xbd, 0xec, 0x82, 0x5d, 0x5b, 0x89, 0x97, 0x39,
	0x4a, 0x53, 0x87, 0xa9, 0x43, 0xf7, 0x7b, 0x8e, 0xd9, 0x38, 0x1d, 0xb7, 0x2e, 0xd6, 0xd4, 0x74,
	0xda, 0xa1, 0x7f, 0xb4, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x25, 0xf1, 0x0a, 0xe1,
	0x04, 0x00, 0x00,
}
