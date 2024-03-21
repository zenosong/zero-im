SERVICE=Zero

SERVICE_STYLE=zero


# ---- You may not need to modify the codes below | 下面的代码大概率不需要更改 ----

# Swagger type, support yml,json | Swagger 文件类型，支持yml,json
SWAGGER_TYPE=json

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w


.PHONY: fmt
fmt: # Format the codes | 格式化代码
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint: # Run go linter | 运行代码错误分析
	golangci-lint run -D staticcheck

.PHONY: tools
tools: # Install the necessary tools | 安装必要的工具
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: gen-api
gen-api: # Generate API files | 生成 API 的代码
	goctl api go --api ./api/docs/all.api --dir ./api
	@echo "Generate API files successfully"

.PHONY: gen-rpc
gen-rpc: # Generate RPC files from proto | 生成 RPC 的代码
	goctl rpc protoc ./rpc/$(SERVICE_STYLE).proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc
	@echo "Generate RPC files successfully"


.PHONY: build-linux
build-linux: # Build project for Linux | 构建Linux下的可执行文件
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_rpc ./rpc/$(SERVICE_STYLE).go
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_api ./api/$(SERVICE_STYLE).go
	@echo "Build project for Linux successfully"