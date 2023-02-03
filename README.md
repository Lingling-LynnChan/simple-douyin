# Simple Douyin

这是一个简单的抖音项目，api源自字节官方

原本希望直接使用 kratos 框架直接进行开发，但是由于组员缺乏对 grpc 以及 微服务框架的理解，同时会增加使用负担，改用 http 的后端框架 iris。

为什么不用 gin？

gin 无疑是一个简洁实用，容易上手的框架，但是倘若直接使用引擎路由绑定函数，则很难以从结构上体现出层次性，同时也难以自定义的进行接口分类(就像controller那样)，只能依据路由进行简单的分组操作。

同时一个函数若仅使用 context 进行传值和交互，在程序员对业务和项目架构不熟悉，或者开始犯病时，则会将代码弄得一团糟，导致难以维护。

grpc 很好的限制了程序员的可能性，从数据的获取与转换的层面遏制了程序员犯病。

无法使用 grpc，因此使用了 iris 的 mvc 模式，它能够做到以上我所提及的 gin 无法做到的事情，当然由于框架内部使用反射，运行效率可能会下降。

呼吁组员在编写代码时，不要将 dto 的数据穿透到 repo 层，这种数据的层级穿透会导致代码难以维护。

## 目录简介
```
// 这里省略了一些细节上的文件表示
// 比如 wire 表示的依赖注入文件都没有展现，大家可以查阅相关文档进行理解
├── api // 下面维护了接口描述文件，他们依据字节官方给定的接口分类为 核心,互动,社交
│   └── core
│   └── interaction
│   └── socialize
├── cmd  // 整个项目启动的入口文件
│   └── main.go
├── configs  // 这里通常维护一些本地调试用的样例配置文件
│   └── config.yaml
├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑
│   ├── biz   // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
│   │   ├── repo      // 仓储层 接口定义，
│   │   └── usecase   // domain层(DDD)，也称为用例层(简洁架构)
│   ├── conf  // 内部使用的config的结构定义
│   │   └── config.go
│   ├── data  // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来
│   │   ├── ent      // 这个文件夹是自动生成的orm代码
│   │   ├── schema   // 这个文件夹中存放orm的数据定义，用于自动生成代码
│   │   ├── data.go
│   │   └── ...
│   ├── server  // http实例的创建和配置
│   │   └── http.go
│   └── service  // 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
│       └── core
│       └── interaction
│       └── socialize
├── third_party  // api 依赖的第三方proto
├── Makefile  
├── README.md
├── go.mod
└── go.sum
```

## 工具介绍

orm框架使用 ent，后端框架使用 iris 框架的 mvc 模式

```sh
# 安装 ent 命令行工具
go install entgo.io/ent/cmd/ent@latest
# 拉取 iris
go get github.com/kataras/iris/v12@master
```

## 命令

### 初始化

```sh
# 这个操作是快速拉取所有依赖
make install
# 这个初始化包含以下所有初始化操作，相当于快捷键
make init
```

### 初始化api

```sh
# 这个命令编译 api 目录中 proto 文件为 go 文件
make api
```

### 初始化orm
```sh
# ent 框架仅需要我们对元数据做出定义
# 之后通过自动代码生成工具生成相应的行为
make ent
```

### 初始化依赖注入
```sh
# 使用 go wire 作为依赖注入工具，在编译期间完成依赖注入
make wire
```

### 运行项目
```sh
# 默认跑在 0.0.0.0:8000 端口
make run
# 指定端口的方法
make run ADDR=:8001
# or
make run ADDR=0.0.0.0:8001
```