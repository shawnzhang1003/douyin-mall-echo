.PHONY: gen-service
gen-service: ## gen service code of {svc}. example: make gen-service svc=product ,需要提前安装protoc
ifeq ($(svc),user)
	@cd rpc_gen && kitex --module github.com/MakiJOJO/douyin-mall-echo/rpc_gen -type protobuf -no-fast-api -I ../idl/ ../idl/${svc}.proto
else
	@cd rpc_gen && kitex --module github.com/MakiJOJO/douyin-mall-echo/rpc_gen -type protobuf -I ../idl/ ../idl/${svc}.proto
		
endif

.PHONY: gen-handler
gen-handler: ## gen handler code of {svc}. example: make gen-handler svc=product ,需要提前安装protoc
	@mkdir -p app/${svc}/rpc && cd app/${svc}/rpc && kitex -module github.com/MakiJOJO/douyin-mall-echo/app/${svc} -type protobuf -I ../../../idl/ -service ${svc} -use rpc_gen/kitex_gen ../../../idl/${svc}.proto


# 定义伪目标
.PHONY: tidy_all

# 查找包含 go.mod 的目录列表
GO_MOD_DIRS := $(shell find . -name "go.mod" -exec dirname {} \;)

# tidy 目标
tidy_all:
	for dir in $(GO_MOD_DIRS); do \
	    echo "Running go mod tidy in $$dir"; \
	    (cd $$dir && go mod tidy); \
	done