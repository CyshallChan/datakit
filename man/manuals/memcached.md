{{.CSS}}
# Memcached
---

- 操作系统支持：{{.AvailableArchs}}

Memcached 采集器可以从 Memcached 实例中采集实例运行状态指标，并将指标采集到观测云，帮助监控分析 Memcached 各种异常情况

![](imgs/input-memcached-1.png)

## 前置条件

- Memcached 版本 >= 1.5.0

## 配置

进入 DataKit 安装目录下的 `conf.d/{{.Catalog}}` 目录，复制 `{{.InputName}}.conf.sample` 并命名为 `{{.InputName}}.conf`。示例如下：

```toml
{{.InputSample}}
```

配置好后，重启 DataKit 即可。

## 指标预览

![](imgs/input-memcached-1.png)

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

<场景 - 新建仪表板 - 内置模板库 - Memcached 监控视图>

## 异常检测

<监控 - 模板新建 - Memcached 检测库>
