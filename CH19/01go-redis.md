# 01 GO-redis

#### Markdown Notes 创建于 2024-08-19T07:09:44.542Z

Redis 教程见以前的笔记

2024 网上的 golang 操作 redis 教程使用的包基本过时。
目前最流行的是这个 redis 官方仓库中的： https://github.com/redis/go-redis
`go get github.com/redis/go-redis/v9`
[go-redis 指南](https://redis.uptrace.dev/guide/)
[package 文档](https://pkg.go.dev/github.com/redis/go-redis/v9)

连接方法：

```go
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
    Protocol: 3, // specify 2 for RESP 2 or 3 for RESP 3
})
```
