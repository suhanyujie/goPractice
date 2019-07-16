## fast-cgi的翻译
* 原文为 https://fast-cgi.github.io/spec

## 章节目录
- [x] 1. Introduction
- [ ] 2. Initial process state
- [ ] 3. Protocol basics
- [ ] 4. Management record types
- [ ] 5. Application record types
- [ ] 6. Roles
- [ ] 7. Errors
- [ ] 8. Types and constants


## 译文内容

### 1.介绍
`FastCGI` 是一个开放的 `CGI` 扩展，它为整个互联网应用提供了高性能而不受 Web API 的影响。

这个规范有很明确的目标：从应用程序的角度看，为了指导做一个支持 FastCGI 的 应用和支持 FastCGI 的 Web 服务器。很多 Web 服务器和 FastCGI 有关联，比如应用管理工具，跟 Web 服务接口无关，这里就不做描述了。

这个规范是针对 Unix （更准确的说是针对支持 `Berkeley Sockets` 的 POSIX 系统）。这个规范的大部分是描述一个简单的通信协议，它不依赖于字节序，可以扩展到其他系统。

我们将通过对比传统的 Unix 的 `CGI/1.1` 的实现来介绍 FastCGI 。 FastCGI 被设计成支持长连接，例如应用服务器。这是和传统 Unix 的 `CGI/1.1` 的实现有很大区别， 传统的 `CGI/1.1` 创建一个进程，使其响应一个请求，然后退出。

FastCGI 的初始化状态比一个 `CGI/1.1` 进程的初始化状态更加简单，因为 FastCGI 进程在声明周期的开始时，没有连接任何东西。它没有按照惯例那样打开 `stdin` ， `stdout` ， `stderr` 这3个文件描述符，并且它无需处理来自环境变量的信息。 FastCGI 进程在初始化状态时的关键一点是监听一个 `socket` ，通过它接收来自 Web 服务的连接。

FastCGI 进程在其监听的套接字上接收到连接后，执行一个简单的协议操作来接收和发送数据。该协议有两个目的。首先，协议在几个独立的 FastCGI 请求之间能多路复用一个传输连接。它支持使用事件驱动或多线程编程技术处理并发请求。其次，在每个请求中，协议在每个方向上提供数个独立的数据流。例如，通过这种方式，stdout 和 stderr 数据都通过一个传输连接从客户端传递到 web 服务器，而不像 CGI/1.1 那样需要单独的通道。

FastCGI 应用扮演了多个角色。最熟悉的是响应者的角色，在这种场景中，应用接收 HTTP 请求相关的所有信息并生成 HTTP 响应；这就是 CGI/1.1 程序要做的。第二种角色是授权者，这种情况下应用接收所有 HTTP 相关的信息并生成一个授权/未授权的结果；第三种角色是 Filter，它是将请求数据流加上 Web 服务器上文件的数据流进行“过滤”并将结果作为响应。该框架是可扩展的，因此后续可以定义更多基于 FastCGI 的应用。

在这个规范的其他部分中，只要不会引起混淆，术语“FastCGI application”，“application process”，或者“application server”都会被简称为“应用”。

### 2. Initial process state
#### 2.1. Argument list

```html
By default the Web server creates an argument list containing a single element, the name of the application, taken to be the last component of the executable’s path name. The Web server may provide a way to specify a different application name, or a more elaborate argument list.

Note that the file executed by the Web server might be an interpreter file (a text file that starts with the characters #!), in which case the application’s argument list is constructed as described in the execve manpage.
```

### 2.初始化进程状态
#### 2.1 参数列表
* 通过默认的Web服务器创建一个包含单一元素的参数列表，应用的名称，成为可执行的路径名称最后一部分。
* 通过 Web 服务提供一个指向不同的应用名称的方式，或者一个更详细的参数列表
* 通过 Web 服务记录执行文件可能是一个解释器文件（一个以标识符 `#!` 开始的文本文件），这种情况下，应用的参数列表就像之前描述的那样在可执行手册中被创建。
