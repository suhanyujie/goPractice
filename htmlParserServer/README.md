# 爬虫解析服务

## 实现
* 通过http POST的接口提供服务
* 解析 `dom` 使用 `goquery`

## build
* 在仓库根目录下，`go build -o htmlParseServer main.go`
* 执行`./a.exe`即可启动服务

## benchamark
* `go test testMain/main_test.go -bench=. -benchtime=3s`

## 调试工具
* 可以借助调试工具delve进行调试 https://www.cnblogs.com/li-peng/p/8522592.html
