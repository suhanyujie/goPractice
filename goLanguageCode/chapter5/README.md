## 第5章 网络编程
* Go语言标准库中提供net，支持IP层、TCP/UDP层以及更高层面的网络操作
* 其中用于IP层的称谓RawSocket

### socket编程
* 主要有一下5个步骤：
    * 建立socket
    * 绑定socket
    * 监听
    * 接收连接
    * 处理连接，接收、发送数据
* 无论用什么协议建立连接，只需要调用`net.Dial()`

### Dial()函数
* 原型如下：`func Dial(net,addr string) (Conn,error)`


### 数据的序列化
* Go语言的大多数数据类型都可以转化为有效的JSON文本，但channel、complex和函数这几种
类型除外。
* 如果转化前的数据结构中出现指针，那么将会转化指针所指向的值，如果指针指向的是零值，
那么null将作为转化后的结果输出

#### json序列化
* 当JSON数据里边的结构和Go里边的目标类型的结构对不上时，会发生什么呢？示例代码如下：

```go
b := []byte(`{"Title": "Go语言编程", "Sales": 1000000}`)
var gobook Book
err := json.Unmarshal(b, &gobook)
```

* 在上面的示例代码中，由于Sales字段并没有在Book类型中定义，所以会被忽略，只有Title这个字段的值才会被填充到gobook.Title中


