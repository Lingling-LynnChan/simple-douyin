GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: install
# 下载依赖
install:
	go mod tidy
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	make ent
	make api

.PHONY: init
# 自动初始化
init:
	go mod tidy
	make api
	make ent
	make wire
	make clean

.PHONY: api
# 生成proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:./api \
	       $(API_PROTO_FILES)

.PHONY: ent
# 自动生成数据库操作代码
ent:
	ent generate --target ./internal/data/ent ./internal/data/schema

.PHONY: wire
# 依赖注入
wire:
	cd internal && wire && cd -

.PHONY: clean
# 清理代码
clean:
	gofmt -l -d -w -s .

addr?=0.0.0.0:8000

.PHONY: run
# 运行项目
run:
	go run ./... -addr $(addr)
