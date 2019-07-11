module storj.io/storj

go 1.12

// force specific versions for minio
require (
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/garyburd/redigo v1.0.1-0.20170216214944-0d253a66e6e1 // indirect
	github.com/graphql-go/graphql v0.7.9-0.20190403165646-199d20bbfed7
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect

	github.com/minio/minio v0.0.0-20180508161510-54cd29b51c38
	github.com/mitchellh/mapstructure v1.1.1 // indirect
	github.com/segmentio/go-prompt v1.2.1-0.20161017233205-f0d19b6901ad
)

exclude gopkg.in/olivere/elastic.v5 v5.0.72 // buggy import, see https://github.com/olivere/elastic/pull/869

require (
	cloud.google.com/go v0.27.0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.6.0 // indirect
	git.apache.org/thrift.git v0.0.0-20180807212849-6e67faa92827 // indirect
	github.com/Microsoft/go-winio v0.4.11 // indirect
	github.com/Shopify/go-lua v0.0.0-20181106184032-48449c60c0a9
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/Sirupsen/logrus v1.0.6 // indirect
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/alicebob/miniredis v0.0.0-20180911162847-3657542c8629
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/aws/aws-sdk-go v1.15.34 // indirect
	github.com/bitly/go-hostpool v0.0.0-20171023180738-a3a6125de932 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/boltdb/bolt v1.3.1
	github.com/cheggaaa/pb v1.0.5-0.20160713104425-73ae1d68fe0b
	github.com/cockroachdb/cockroach-go v0.0.0-20180212155653-59c0560478b7 // indirect
	github.com/cznic/ql v1.2.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/djherbis/atime v1.0.0 // indirect
	github.com/docker/distribution v0.0.0-20180720172123-0dae0957e5fe // indirect
	github.com/docker/docker v0.0.0-20170502054910-90d35abf7b35 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.3.3 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/eclipse/paho.mqtt.golang v1.1.1 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/fatih/color v1.7.0
	github.com/fatih/structs v1.0.0 // indirect
	github.com/fsouza/fake-gcs-server v1.1.0 // indirect
	github.com/go-ini/ini v1.38.2 // indirect
	github.com/go-redis/redis v6.14.1+incompatible
	github.com/go-sql-driver/mysql v1.4.0 // indirect
	github.com/gocql/gocql v0.0.0-20180913072538-864d5908455a // indirect
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.3.0
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v0.0.0-20170111101155-53e6ce116135 // indirect
	github.com/google/martian v2.0.0-beta.2+incompatible // indirect
	github.com/googleapis/gax-go v2.0.0+incompatible // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/handlers v1.4.0 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/gorilla/rpc v1.1.0 // indirect
	github.com/gorilla/schema v1.1.0
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.3 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/hashicorp/raft v1.0.0 // indirect
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c // indirect
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf // indirect
	github.com/jbenet/go-base58 v0.0.0-20150317085156-6237cf65f3a6
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/jtolds/go-luar v0.0.0-20170419063437-0786921db8c0
	github.com/jtolds/monkit-hw v0.0.0-20190108155550-0f753668cf20
	github.com/klauspost/cpuid v0.0.0-20180405133222-e7e905edc00e // indirect
	github.com/klauspost/reedsolomon v0.0.0-20180704173009-925cb01d6510 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kshvakov/clickhouse v1.3.4 // indirect
	github.com/lib/pq v1.0.0
	github.com/loov/hrtime v0.0.0-20181214195526-37a208e8344e
	github.com/loov/plot v0.0.0-20180510142208-e59891ae1271
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/mattn/go-runewidth v0.0.3 // indirect
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/minio/cli v1.3.0
	github.com/minio/dsync v0.0.0-20180124070302-439a0961af70 // indirect
	github.com/minio/highwayhash v0.0.0-20180501080913-85fc8a2dacad // indirect
	github.com/minio/lsync v0.0.0-20180328070428-f332c3883f63 // indirect
	github.com/minio/mc v0.0.0-20180926130011-a215fbb71884 // indirect
	github.com/minio/minio-go v6.0.3+incompatible
	github.com/minio/sha256-simd v0.0.0-20190328051042-05b4dd3047e5
	github.com/minio/sio v0.0.0-20180327104954-6a41828a60f0 // indirect
	github.com/mitchellh/go-homedir v0.0.0-20180801233206-58046073cbff // indirect
	github.com/nats-io/gnatsd v1.3.0 // indirect
	github.com/nats-io/go-nats v1.6.0 // indirect
	github.com/nats-io/go-nats-streaming v0.4.2 // indirect
	github.com/nats-io/nats v1.6.0 // indirect
	github.com/nats-io/nats-streaming-server v0.12.2 // indirect
	github.com/nats-io/nuid v1.0.0 // indirect
	github.com/nsf/jsondiff v0.0.0-20160203110537-7de28ed2b6e3
	github.com/nsf/termbox-go v0.0.0-20190121233118-02980233997d
	github.com/onsi/ginkgo v1.7.0 // indirect
	github.com/onsi/gomega v1.4.3 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/openzipkin/zipkin-go v0.1.1 // indirect
	github.com/pascaldekloe/goe v0.0.0-20180627143212-57f6aae5913c // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/pkg/profile v1.2.1 // indirect
	github.com/prometheus/client_golang v0.9.3 // indirect
	github.com/prometheus/procfs v0.0.0-20190517135640-51af30a78b0e // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/rs/cors v1.5.0 // indirect
	github.com/sirupsen/logrus v1.3.0 // indirect
	github.com/skyrings/skyring-common v0.0.0-20160929130248-d1c0bb1cbd5e
	github.com/smartystreets/assertions v0.0.0-20180820201707-7c9eb446e3cf // indirect
	github.com/smartystreets/goconvey v0.0.0-20180222194500-ef6db91d284a // indirect
	github.com/spacemonkeygo/errors v0.0.0-20171212215202-9064522e9fd1 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.2.1
	github.com/streadway/amqp v0.0.0-20180806233856-70e15c650864 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/stripe/stripe-go v60.17.0+incompatible
	github.com/tidwall/gjson v1.1.3 // indirect
	github.com/tidwall/match v0.0.0-20171002075945-1731857f09b1 // indirect
	github.com/vivint/infectious v0.0.0-20190108171102-2455b059135b
	github.com/yuin/gopher-lua v0.0.0-20180918061612-799fa34954fb // indirect
	github.com/zeebo/admission v0.0.0-20180821192747-f24f2a94a40c
	github.com/zeebo/errs v1.2.1-0.20190617123220-06a113fed680
	github.com/zeebo/float16 v0.1.0 // indirect
	github.com/zeebo/incenc v0.0.0-20180505221441-0d92902eec54 // indirect
	github.com/zeebo/structs v1.0.2
	go.etcd.io/bbolt v1.3.2 // indirect
	go.opencensus.io v0.15.0 // indirect
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/exp v0.0.0-20190121172915-509febef88a4 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	golang.org/x/tools v0.0.0-20190614152001-1edc8e83c897
	google.golang.org/api v0.0.0-20180910000450-7ca32eb868bf // indirect
	google.golang.org/appengine v1.6.0 // indirect
	google.golang.org/grpc v1.2.1-0.20190710205655-08d23162a9a5
	gopkg.in/Shopify/sarama.v1 v1.18.0 // indirect
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.25 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gopkg.in/ini.v1 v1.38.2 // indirect
	gopkg.in/olivere/elastic.v5 v5.0.76 // indirect
	gopkg.in/spacemonkeygo/monkit.v2 v2.0.0-20190612171030-cf5a9e6f8fd2
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
