{{.CSS}}
# 内存
---

- 操作系统支持：{{.AvailableArchs}}

mem 采集器用于收集系统内存信息，一些通用的指标如主机总内存、用的内存、已使用的内存等  

![](imgs/input-mem-1.png)

## 前置条件

暂无

## 配置

进入 DataKit 安装目录下的 `conf.d/{{.Catalog}}` 目录，复制 `{{.InputName}}.conf.sample` 并命名为 `{{.InputName}}.conf`。示例如下：

```toml
{{.InputSample}}
```

配置好后，重启 DataKit 即可。

支持以环境变量的方式修改配置参数（只在 DataKit 以 K8s daemonset 方式运行时生效，主机部署的 DataKit 不支持此功能）：

| 环境变量名               | 对应的配置参数项 | 参数示例                                                     |
| :---                     | ---              | ---                                                          |
| `ENV_INPUT_MEM_TAGS`     | `tags`           | `tag1=value1,tag2=value2` 如果配置文件中有同名 tag，会覆盖它 |
| `ENV_INPUT_MEM_INTERVAL` | `interval`       | `10s`                                                        |

## 指标预览

![](imgs/input-mem-3.png)

## 指标集

以下所有数据采集，默认会追加名为 `host` 的全局 tag（tag 值为 DataKit 所在主机名），也可以在配置中通过 `[inputs.{{.InputName}}.tags]` 指定其它标签：

``` toml
 [inputs.{{.InputName}}.tags]
  # some_tag = "some_value"
  # more_tag = "some_other_value"
  # ...
```

{{ range $i, $m := .Measurements }}

### `{{$m.Name}}`

-  标签

{{$m.TagsMarkdownTable}}

- 指标列表

{{$m.FieldsMarkdownTable}}

{{ end }}

## 场景视图

<场景 - 新建仪表板 - 内置模板库 - Memory>

## 异常检测

<监控 - 模板新建 - 主机检测库>
