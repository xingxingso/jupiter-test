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

> 设置 GOPATH 等

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

> 本地数据库初始化, 配置数据库连接等 `config/config.toml`

```bash
cd cmd
go run main.go --config=../config/config.toml
```

>> [http://127.0.0.1:20105/jupiter](http://127.0.0.1:20105/jupiter)

>> [http://127.0.0.1:20105/api/user/1](http://127.0.0.1:20105/api/user/1)

## some exercise

> 数据库连接 - 参见 `readme-jupiter.md`

