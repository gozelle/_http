# _http 网络请求库

## 特性

* 上手可用
* 并发安全
* 默认执行 JSON 请求
* 支持上下文日志
* 统计请求耗时
* 支持网络代理
* 支持多种请求类型、JSON、表单、文件上传、下载等
* 支持请求参数排序签名自定义实现
* 支持相应结果转化为 FastJson

## TODO

* 拼接请求上下文日志

## 安装

```go
go get -u github.com/koyeo/_http
```

## 示例

```go
 resp := _http.Post("https://api.exmaple.com", Param("username","root"), Param("level", 1))
 if resp.Error() != nil{
    return resp.Error()	
 }

 fmt.Println(resp.String())
```

