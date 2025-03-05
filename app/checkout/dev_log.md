# checkout开发日志
## 0217 
1. 安装环境
- kitex：go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
- 设置环境变量： export PATH=$PATH:/home/shawn/go/bin
- 安装protoc编译器，makefile生成rpc客户端代码

2. checkout只有逻辑上的处理，没有涉及数据库，因此无model层

3. 修改官方的接口设计：不在结算中调用支付服务，修改checkoutResp中order_id为uint64类型，与订单服务一致

4. func 结算handler {
    调用购物车，获取物品list
    调用product，获取物品总价
    调用订单，创建订单
}

5. 完成结算：http_handler调用购物车rpc 已过postman测试

6. 初始化rpc服务端，查看已注册的服务etcdctl get --prefix "kitex"

## 0219
1. 重构checkout


