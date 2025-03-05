# 脚手架生成代码存放处
app/douyin-mall-echo-templete/rpc
该目录下是使用kitex框架生成的rpc服务代码，包含服务定义, rpc handler的实现
命令如下:
kitex -module github.com/MakiJOJO/douyin-mall-echo/rpc_gen -type protobuf -I /Users/jojoman/work/douyin-mall-echo/idl/ -service user -use rpc_gen/kitex_gen /Users/jojoman/work/douyin-mall-echo/idl/user.proto


# 以下为非脚手架生成代码
app/douyin-mall-echo-templete/rpc/client.go 
是rpc服务的客户端初始化代码,用来调用其他服务,如果不需要调用其他服务,可以不用