{{.CSS}}

- DataKit 版本：{{.Version}}
- 文档发布日期：{{.ReleaseDate}}
- 操作系统支持：`{{.AvailableArchs}}`

# {{.InputName}}

## Jaeger 文档

- [getting started](https://www.jaegertracing.io/docs/1.27/getting-started/)
- [source code](https://github.com/jaegertracing/jaeger)
- [client download](https://github.com/jaegertracing/jaeger-client-go/releases)

## 配置

进入 DataKit 安装目录下的 `conf.d/{{.Catalog}}` 目录，复制 `{{.InputName}}.conf.sample` 并命名为 `{{.InputName}}.conf`。示例如下：

```toml
{{.InputSample}}
```

### 配置 Jaeger HTTP Agent

endpoint 代表 Jaeger HTTP Agent 路由

```toml
[[inputs.jaeger]]
  # Jaeger endpoint for receiving tracing span over HTTP.
  # Default value set as below. DO NOT MODIFY THE ENDPOINT if not necessary.
  endpoint = "/apis/traces"
```

- 修改 Jaeger Client 的 Agent Host Port 为 Datakit Port（默认为 9529）
- 修改 Jaeger Client 的 Agent endpoint 为上面配置中指定的 endpoint

### 配置 Jaeger UDP Agent

```toml
[[inputs.jaeger]]
  # Jaeger agent host:port address for UDP transport.
  address = "127.0.0.1:6831"
```

- 修改 Jaeger Client 的 Agent UDP Host:Port 为上面配置中指定的 address

有关数据采样，数据过滤，关闭资源等配置请参考[Datakit Tracing](datakit-tracing)

## Golang 示例

```golang
package main

import (
	"log"
	"time"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func main() {
	jgcfg := jaegercfg.Configuration{
		ServiceName: "jaeger_sample_code",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			BufferFlushInterval: time.Second,
			LocalAgentHostPort:  "127.0.0.1:6831",
			// CollectorEndpoint:   "http://localhost:9529/jaeger/traces",
			LogSpans: true,
			// HTTPHeaders:         map[string]string{"Content-Type": "application/vnd.apache.thrift.binary"},
		},
	}

	tracer, closer, err := jgcfg.NewTracer(jaegercfg.Logger(jaegerlog.StdLogger))
	defer func() {
		if err := closer.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	if err != nil {
		log.Panicln(err.Error())
	}

	for {
		span := tracer.StartSpan("test_start_span")
		log.Println("start new span")
		span.SetTag("key", "value")
		span.Finish()
		log.Println("new span finished")

		time.Sleep(time.Second)
	}
}

```
