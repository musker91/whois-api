# Whois 查询

## 部署教程
> 需要具备go语言环境, version >= 1.18.3

1.编译代码

```text
$ // 下载源码
$ go env -w GOPROXY=https://goproxy.io,direct
$ cd whois-api

> 编译可执行文件
$ make clean       // 清理
$ make build       // 编译二进制

// 或者使用下面的手动编译
$ GOOS=linux GOARCH=amd64 go build -o whois-api main.go
```

2.修改配置文件

修改项目下 `config.yml` 文件

3.运行服务

```
$ ./whois-api
```

## 接口文档

### - 接口地址
`http://ip:port/api`

### - 请求参数

字段名称 | 类型 | 必填 | 说明
--- | --- | --- | ---
domain | String | 是 | 域名
type | String | 否 | whois数据返回类型(text 文本串/json json格式数据)
standard | Bool | 否 | 是否按照标准固定格式输出json字段，默认是按原whois信息中的所有字段返回，只对返回json格式有效

### - 请求示例

`http://ip:port/api?domain=devopsclub.cn&type=json&standard=true`

### - 返回参数说明

字段名称 | 类型 | 说明
--- | --- | ---
status | Int | 域名查询状态(0 获取域名whois信息成功/1 域名解析失败/2 域名未注册/3 暂不支持此域名后缀查询/4 域名查询失败/5 请求参数错误)
data | Map/String | 域名whois详细数据
msg | String | 消息