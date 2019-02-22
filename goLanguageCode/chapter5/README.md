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


