# Nexus SDK

## 引入

```shell
go get -u github.com/xuxiaowei-com-cn/go-nexus@main
```

## 使用

- 使用方式见测试文件

## 开发命令

### get

```shell
go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct
go get -u github.com/hashicorp/go-retryablehttp
go get -u github.com/stretchr/testify
```

### mod

```shell
go mod tidy
```

```shell
go mod download
```

### test

```shell
go test ./... -v
```
