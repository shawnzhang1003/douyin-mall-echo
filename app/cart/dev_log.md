# cart开发日志
## 0215
1. 安装环境
- kitex：go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
- 设置环境变量： export PATH=$PATH:/home/shawn/go/bin
- 安装protoc编译器，makefile生成rpc客户端代码

2. model层
- 定义购物车模型
- 实现添加物品、按用户获取购物车信息、清空购物车的方法
- 在main函数中开启模型自动迁移

3. logic层
- 这里是对http请求的handler函数层
- 对route分发的http请求预处理，并调用对应的model层方法操作数据库

4. route层
- 注册路由
```go
	api.POST("/cart", logic.AddItem)
	api.GET("/cart", logic.GetCart)
	api.DELETE("/cart", logic.EmptyCart)
```

5. 修改mysql配置

## 0216
1. 更新了idl中CartItem的数量类型为uint32
2. 初始化cart rpc服务
- rpc_server：拉起rpc服务端，向etcd注册服务
- rpc_handler：处理rpc请求的handler
3. cmd文件夹下 docker compose 拉起etcd服务
- sudo docker exec -it etcd bash 进入docker bash
- etcdctl get --prefix "kitex" 查看已注册的服务
```bash
I have no name!@29332dc8299e:/opt/bitnami/etcd$ etcdctl get --prefix "kitex"
kitex/registry-etcd/cart/127.0.0.1:8890
{"network":"tcp","address":"127.0.0.1:8890","weight":10,"tags":null}
```
4. 编写单测测试cart的rpc调用 rpc_handler_test.go

## 0219 
1. merge main分支，重构购物车服务

2. https://github.com/hertz-contrib/obs-opentelemetry/issues/30 遗留报错待解决
- 2025/02/19 02:42:02 traces export: context deadline exceeded: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing: dial tcp 127.0.0.1:4317: connect: connection refused"

3. 