module gitlab.jiagouyun.com/cloudcare-tools/datakit

go 1.16

require (
	code.cloudfoundry.org/bytefmt v0.0.0-20200131002437-cf55d5288a48
	github.com/BurntSushi/toml v0.3.1
	github.com/DataDog/ebpf v0.0.0-20210419131141-ea64821c9793
	github.com/DataDog/gopsutil v0.0.0-20210826200402-bbfc5b0ae6e9 // indirect
	github.com/MichaelMure/go-term-markdown v0.1.3
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/alecthomas/repr v0.0.0-20181024024818-d37bc2a10ba1 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.2.0+incompatible
	github.com/apache/thrift v0.13.0
	github.com/araddon/dateparse v0.0.0-20201001162425-8aadafed4dc4
	github.com/c-bata/go-prompt v0.2.5
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/chromedp/cdproto v0.0.0-20210910012206-68626162910d // indirect
	github.com/chromedp/chromedp v0.7.4 // indirect
	github.com/containerd/cgroups v1.0.1
	github.com/containerd/containerd v1.5.5 // indirect
	github.com/denisenkom/go-mssqldb v0.10.0
	github.com/didip/tollbooth/v6 v6.1.2
	github.com/docker/docker v20.10.8+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/elazarl/goproxy v0.0.0-20210801061803-8e322dfb79c4
	github.com/fatih/color v1.12.0
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gdamore/tcell/v2 v2.4.1-0.20210905002822-f057f0a857a1 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-git/go-git/v5 v5.4.2
	github.com/go-ole/go-ole v1.2.5
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/go-redis/redis/v8 v8.11.3
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gobuffalo/logger v1.0.4 // indirect
	github.com/gobuffalo/packr/v2 v2.8.1
	github.com/gobwas/glob v0.2.3
	github.com/gobwas/ws v1.1.0
	github.com/godror/godror v0.17.0
	github.com/gofrs/flock v0.8.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.4
	github.com/gomarkdown/markdown v0.0.0-20210208175418-bda154fe17d8
	github.com/google/gopacket v1.1.19
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/gorilla/websocket v1.5.0
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/influxdata/telegraf v1.15.2
	github.com/influxdata/toml v0.0.0-20190415235208-270119a8ce65
	github.com/ip2location/ip2location-go v8.3.0+incompatible
	github.com/jessevdk/go-flags v1.5.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kardianos/service v1.0.0
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/lib/pq v1.10.2
	github.com/mdlayher/netlink v1.4.1 // indirect
	github.com/mssola/user_agent v0.5.2
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/pborman/ansi v1.0.0
	github.com/pkg/sftp v1.11.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.32.1
	github.com/prometheus/procfs v0.7.3
	github.com/prometheus/prometheus v2.5.0+incompatible
	github.com/rivo/tview v0.0.0-20220129131435-1f7581b67bd1
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/shirou/gopsutil v3.21.8+incompatible
	github.com/shirou/gopsutil/v3 v3.20.12
	github.com/spf13/cast v1.4.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.12.1
	github.com/tinylib/msgp v1.1.6
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tweekmonster/luser v0.0.0-20161003172636-3fa38070dbd7
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/ugorji/go/codec v1.2.6
	github.com/vishvananda/netlink v1.1.1-0.20210508154835-66ddd91f7ddd // indirect
	github.com/vishvananda/netns v0.0.0-20210104183010-2eb08e3e575f // indirect
	github.com/vjeantet/grok v1.0.0
	github.com/whilp/git-urls v1.0.0
	gitlab.jiagouyun.com/cloudcare-tools/cliutils v0.0.0-20220303071313-73cbbd097512
	gitlab.jiagouyun.com/cloudcare-tools/kodo v0.0.0-20210923054635-8b8d5dd9ba4f
	go.etcd.io/bbolt v1.3.5
	go.opentelemetry.io/otel v1.4.1
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.4.1
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.4.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.4.1
	go.opentelemetry.io/otel/metric v0.27.0
	go.opentelemetry.io/otel/sdk v1.4.1
	go.opentelemetry.io/otel/sdk/metric v0.27.0
	go.opentelemetry.io/otel/trace v1.4.1
	go.opentelemetry.io/proto/otlp v0.12.0
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	golang.org/x/text v0.3.7
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/CodapeWild/dd-trace-go.v1 v1.35.17
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/klog/v2 v2.40.1 // indirect
	k8s.io/metrics v0.20.5
	k8s.io/utils v0.0.0-20211208161948-7d6a63dca704 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/c-bata/go-prompt => github.com/coanor/go-prompt v0.2.6

// added for ddtrace
replace (
	github.com/iovisor/gobpf => github.com/DataDog/gobpf v0.0.0-20210322155958-9866ef4cd22c
	k8s.io/api => k8s.io/api v0.20.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.5
	k8s.io/apiserver => k8s.io/apiserver v0.20.5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.5
	k8s.io/client-go => k8s.io/client-go v0.20.5
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.5
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.5
	k8s.io/code-generator => k8s.io/code-generator v0.20.5
	k8s.io/component-base => k8s.io/component-base v0.20.5
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.5
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.5
	k8s.io/cri-api => k8s.io/cri-api v0.20.5
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.5
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.5
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.5
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.5
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.5
	k8s.io/kubectl => k8s.io/kubectl v0.20.5
	k8s.io/kubelet => k8s.io/kubelet v0.20.5
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.5
	k8s.io/metrics => k8s.io/metrics v0.20.5
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.3-rc.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.5
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.20.5
	k8s.io/sample-controller => k8s.io/sample-controller v0.20.5
)
