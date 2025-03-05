# douyin-mall-echo

# 技术栈
http框架 echo/v4 
rpc框架 kitex
配置文件管理 viper
日志管理 slog
监控 prometheus
数据存储 gorm
数据库 mysql
缓存 redis
消息队列 RocketMQ


# 项目架构
```
 _______________________________________
| echo http handler | kitex rpc handler |   #处理参数校验/预处理,返回值封装等
        |                     |
        v                     v
 _______________________________________
|                 logic                 |   #处理具体业务逻辑
        |                     |
        v                     v
 _______________________________________
|     model     |      rpc client       |   #处理数据模型,数据库操作,远程调用等
        |                     |
        v                     v
 _______________________________________
| mysql redis | other rpc server handler|  #数据源
```

# make使用脚手架生成代码说明
生成/更新proto文件后,重新生成/更新相关rpc服务的生成代码
```shell
make gen-service svc=sevicename
```

生成rpc handler代码
```shell
make gen-handler svc=sevicename
```

# 监听端口
#### service port

| service name | api service port(808x) | rpc service port(888x) | other service port |
| ------------ | ---------------------- | ---------------------- | ------------------------ |
| user         | 8080                   | 8880                   |                          |
| order        | 8081                   | 8881                   | mq-xxxx                  |
| payment      | 8082                   | 8882                   |        -                  |
| cart         | 8083                   | 8883                   |         -                 |
| product      | 8084                   | 8884                   |              -            |
| checkout     | 8085                   | 8885                   |              -            |
| auth         | 8086                   | 8886                   |              -             |



#### Prometheus Port


| service name     | prometheus port |
| ---------------- | --------------- |
| user             | 9000            |
| order            | 9001            |
| payment          | 9002            |
| cart             | 9003            |
| product          | 9004            |
| checkout         | 9005            |
| auth             | 9006            |

# 常见问题

## ./app文件夹下的微服务项目的package或者依赖检查报错
在douyin-mall-echo目录下运行go work use ./app/xxx, 然后在./app/xxx下运行go mod tidy

## 项目根目录使用go work,在项目根目录下运行go run ./app/xxx/cmd 能正常运行, 但是在cd ./app/xxx下运行go mod tidy会报错
这是官方包依赖管理问题, 请参考:
https://github.com/golang/go/issues/50750
临时解决方案在./app/xxx/go.mod中加入replace 比如replace github.com/MakiJOJO/douyin-mall-echo/common => ../../common

## 如何把本微服添加进指标监控
在deploy/prometheus/prometheus.yml中添加

```yaml
  - job_name: 'user'
    static_configs:
      - targets: [ 'douyin:9000' ]
        labels:
          job: user
          app: user
          env: dev
```

如果服务没有运行在同个Docker网络中, 请把"douyin"换成对应的服务IP地址

## 如何打印日志
1. 使用原生的log包,但是格式过于简单,需要自行配置日志输出文件, 但是打印内容少方便调试, 所以debug和开发调试阶段可以临时使用

2. 使用common/mtl包的Logger, 它是slog.Logger,可以自动给日志加上时间, source code位置等信息,格式是json格式
作为区分echo输出的请求日志不是json格式

> Tips: 因为mtl.InitLogger的操作放在了server.NewEchoServer里面，在调用server.NewEchoServer前用mtl.Logger会空指针panic. 
*正式服上线时采用这个Logger, 方便从日志文件中找出问题原因*

在config.yaml配置文件中可以配置log存储的位置, 默认是为空就会打印日志到stdout
```json
{"time":"2025-02-19T03:01:27.267559+08:00","level":"INFO","source":{"function":"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/routers.HelloWorldHandler","file":"/Users/jojoman/work/douyin-mall-echo/app/douyin-mall-echo-templete/internal/routers/handler.go","line":17},"msg":"Hello World","service_name":"user"}
```
3. 使用echo.Logger, 它是echo框架的日志, 默认是打印到stdout
```json
{"time":"2025-02-19T03:01:27.267504+08:00","level":"INFO","prefix":"echo","file":"handler.go","line":"16","message":"Hello World"}
```
## 服务启动了但是Prometheus监控服务状态是down
请检查是否在prometheus.yml中添加了对应的job_name, 并且targets的地址端口是否正确

## grafaana监控dashboard没有数据
需要配置data source和添加变量, 请参考:【【新手教程】15可观测性（指标）】 https://www.bilibili.com/video/BV18J4m1T72o/?share_source=copy_web&vd_source=c00b90c0ac58c6bf8d6e017fcbabc320

## .proto注册失败,命名冲突
报错信息: 
```shell
panic: proto: file "auth.proto" is already registered
        See https://protobuf.dev/reference/go/faq#namespace-conflict


goroutine 1 [running]:
google.golang.org/protobuf/reflect/protoregistry.init.func1({0x10630e730?, 0x1059945a0?}, {0x1059c15e0, 0x140004e6950})
        /Users/xxx/go/pkg/mod/google.golang.org/protobuf@v1.36.3/reflect/protoregistry/registry.go:56 +0x1f8
google.golang.org/protobuf/reflect/protoregistry.(*Files).RegisterFile(0x1400000c918, {0x1059dfdf0, 0x1400188f0e0})
        /Users/xxx/go/pkg/mod/google.golang.org/protobuf@v1.36.3/reflect/protoregistry/registry.go:130 +0x9b0
google.golang.org/protobuf/internal/filedesc.Builder.Build({{0x0, 0x0}, {0x140018a0200, 0x1f4, 0x200}, 0x1, 0x4, 0x0, 0x0, {0x1059c8f68, ...}, ...})
        /Users/xxx/go/pkg/mod/google.golang.org/protobuf@v1.36.3/internal/filedesc/build.go:112 +0x1a0
github.com/golang/protobuf/proto.RegisterFile({0x1054ed4fc, 0xa}, {0x10629fc00, 0x152, 0x152})
        /Users/xxx/go/pkg/mod/github.com/golang/protobuf@v1.5.4/proto/registry.go:48 +0x144
go.etcd.io/etcd/api/v3/authpb.init.1()
        /Users/xxx/go/pkg/mod/go.etcd.io/etcd/api/v3@v3.5.12/authpb/auth.pb.go:232 +0x40
exit status 2
```
解决方法:
1. 由于注册中心使用的是etcd, etcd package初始化时自带auth.proto文件和我们的auth.proto文件名字有冲突, 请在proto文件中加入option go_package = "github.com/MakiJOJO/douyin-mall-echo/app/auth"; 来避免命名冲突(没有尝试这个方法)

2. 更改./idl/auth.proto,改为authentication.proto, 然后重新生成代码(成功)