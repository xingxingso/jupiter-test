[TOC]

# README

## start

> 安装 go

```bash
go version
# go version go1.14 darwin/amd64
```

> 开启 GO111MODULE

```bash
export GO111MODULE=on
```

> 设置 gopath 等

```bash
vi ~/.bash_profile
```

```text
# go
export GOPATH=/opt/www/go
export PATH=$PATH:${GOPATH//://bin:}/bin
export GOBIN=
```

> 安装 protobuf

```bash
brew install protobuf
```

> 安装 jupiter, 初始化项目

```bash
go get -u github.com/douyu/jupiter/tools/jupiter
cd /path/to/workspace
jupiter new jupiter-test
```