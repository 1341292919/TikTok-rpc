# 辅助工具安装列表
# 执行 go install github.com/cloudwego/hertz/cmd/hz@latest
# 执行 go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
# 访问 https://golangci-lint.run/welcome/install/ 以查看安装 golangci-lint 的方法

# 项目 MODULE 名
MODULE = TikTok-rpc
# 服务名
SERVICES := gateway user video interact websocket
# 当前架构
ARCH := $(shell uname -m)

# 目录相关
DIR = $(shell pwd)
IDL_PATH = $(DIR)/idl

.PHONY: tidy
tidy:
	go mod tidy
	go mod verify

#启动相应服务
.PHONY: $(addprefix run-,$(SERVICES))
$(addprefix run-,$(SERVICES)): run-%:
	@echo "Building $* service..."
	@go build -o bin/$* cmd/$*/main.go && ./bin/$* &
# 默认运行所有服务
run-all: $(addprefix run-,$(SERVICES))

#停止相应服务
.PHONY: $(addprefix stop-,$(SERVICES))
$(addprefix stop-,$(SERVICES)): stop-%:
	 -pkill -f "bin/$*"
stop-all :$(addprefix stop-,$(SERVICES))

# 生成基于 Hertz 的脚手架
.PHONY: hz-%
hz-%:
	hz update -idl ${IDL_PATH}/api/$*.thrift

# 基于 idl 生成相关的 go 语言描述文件
.PHONY: kitex-gen-%
kitex-gen-%:
	@ kitex -module "${MODULE}" \
		${IDL_PATH}/$*.thrift
	@ if [ "$*" != "model" ]; then \
        mkdir -p rpc/$* && \
        cd rpc/$* && kitex -module "${MODULE}" \
            -thrift no_default_serdes \
            -service $* \
            ${IDL_PATH}/$*.thrift; \
   	fi
	@ go mod tidy

# 启动必要的环境，比如 etcd、mysql
.PHONY: env-up
env-up:
	@ docker compose -f ./docker/docker-compose.yml up -d

# 关闭必要的环境，但不清理 data（位于 docker/data 目录中）
.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

# 清除所有的构建产物
.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print

# 清除所有构建产物、compose 环境和它的数据
.PHONY: clean-all
clean-all: clean
	@echo "$(PREFIX) Checking if docker-compose services are running..."
	@docker compose -f ./docker/docker-compose.yml ps -q | grep '.' && docker compose -f ./docker/docker-compose.yml down || echo "$(PREFIX) No services are running."
	@echo "$(PREFIX) Removing docker data..."
	rm -rf ./docker/data

# 优化 import 顺序结构
.PHONY: import
import:
	goimports -w -local Tiktok-rpc .

# 检查可能的错误
.PHONY: vet
vet:
	go vet ./...

# 代码格式校验
.PHONY: lint
lint:
	golangci-lint run --config=./.golangci.yml --tests --allow-parallel-runners --sort-results --show-stats --print-resources-usage

# 检查依赖漏洞
.PHONY: vulncheck
vulncheck:
	govulncheck ./...

.PHONY: build-%
build-%:
	@read -p "Confirm service name to push (type '$*' to confirm): " CONFIRM_SERVICE; \
	if [ "$$CONFIRM_SERVICE" != "$*" ]; then \
		echo "Confirmation failed. Expected '$*', but got '$$CONFIRM_SERVICE'."; \
		exit 1; \
	fi; \
	if echo "$(SERVICES)" | grep -wq "$*"; then \
		if [ "$(ARCH)" = "x86_64" ] || [ "$(ARCH)" = "amd64" ]; then \
			echo "Building and pushing $* for amd64 architecture..."; \
			docker build --build-arg SERVICE=$* -t tiktok:$* -f docker/Dockerfile .; \
		else \
			echo "Building and pushing $* using buildx for amd64 architecture..."; \
			docker buildx build --platform linux/amd64 --build-arg SERVICE=$* -t :$* -f docker/Dockerfile --push .; \
		fi; \
	else \
		echo "Service '$*' is not a valid service. Available: [$(SERVICES)]"; \
		exit 1; \
	fi


