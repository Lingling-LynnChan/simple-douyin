# Simple Douyin

这是一个简单的抖音项目，api源自字节官方

## 工具介绍

orm框架使用 ent，后端框架使用 grpc?

## 命令

### 初始化

```sh
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
# todo
# 使用 go wire 作为依赖注入工具，在编译期间完成依赖注入
make wire
```