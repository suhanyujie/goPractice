>* fast-cgi 规范翻译
>* 原文链接 https://fast-cgi.github.io/spec
>* 译者：[suhanyujie](https://www.github.com/suhanyujie)

## 章节目录
- [x] 1. Introduction
- [x] 2. Initial process state
    - [x] 2.1. Argument list
    - [ ] 2.2. File descriptors
    - [ ] 2.3. Environment variables
    - [ ] 2.4. Other state
- [ ] 3. Protocol basics
    - [ ] 3.1. Notation
    - [ ] 3.2. Accepting transport connections
    - [ ] 3.3. Records
    - [ ] 3.4. Name-Value pairs
    - [ ] 3.5. Closing transport connections
- [ ] 4. Management record types
    - [ ] 4.1. FCGI_GET_VALUES, FCGI_GET_VALUES_RESULT
    - [ ] 4.2. FCGI_UNKNOWN_TYPE
- [ ] 5. Application record types
    - [ ] 5.1. FCGI_BEGIN_REQUEST
    - [ ] 5.2. Name-Value pair streams: FCGI_PARAMS
    - [ ] 5.3. Byte streams: FCGI_STDIN, FCGI_DATA, FCGI_STDOUT, FCGI_STDERR
    - [ ] 5.4. FCGI_ABORT_REQUEST
    - [ ] 5.5. FCGI_END_REQUEST
- [ ] 6. Roles
    - [ ] 6.1. Role protocols
    - [ ] 6.2. Responder
    - [ ] 6.3. Authorizer
    - [ ] 6.4. Filter
- [ ] 7. Errors
- [ ] 8. Types and constants
- [ ] 9. References
- [ ] A. Table: properties of the record types
- [ ] B. Typical protocol message flow


>* 译文内容

## 1.介绍
`FastCGI` 是一个开放的 `CGI` 扩展，它为整个互联网应用提供了高性能而不受 Web API 的影响。

这个规范有很明确的目标：从应用程序的角度看，为了指导做一个支持 FastCGI 的 应用和支持 FastCGI 的 Web 服务器。很多 Web 服务器和 FastCGI 有关联，比如应用管理工具，跟 Web 服务接口无关，这里就不做描述了。

这个规范是针对 Unix （更准确的说是针对支持 `Berkeley Sockets` 的 POSIX 系统）。这个规范的大部分是描述一个简单的通信协议，它不依赖于字节序，可以扩展到其他系统。

我们将通过对比传统的 Unix 的 `CGI/1.1` 的实现来介绍 FastCGI 。 FastCGI 被设计成支持长连接，例如应用服务器。这是和传统 Unix 的 `CGI/1.1` 的实现有很大区别， 传统的 `CGI/1.1` 创建一个进程，使其响应一个请求，然后退出。

FastCGI 的初始化状态比一个 `CGI/1.1` 进程的初始化状态更加简单，因为 FastCGI 进程在声明周期的开始时，没有连接任何东西。它没有按照惯例那样打开 `stdin` ， `stdout` ， `stderr` 这3个文件描述符，并且它无需处理来自环境变量的信息。 FastCGI 进程在初始化状态时的关键一点是监听一个 `socket` ，通过它接收来自 Web 服务的连接。

FastCGI 进程在其监听的套接字上接收到连接后，执行一个简单的协议操作来接收和发送数据。该协议有两个目的。首先，协议在几个独立的 FastCGI 请求之间能多路复用一个传输连接。它支持使用事件驱动或多线程编程技术处理并发请求。其次，在每个请求中，协议在每个方向上提供数个独立的数据流。例如，通过这种方式，stdout 和 stderr 数据都通过一个传输连接从客户端传递到 web 服务器，而不像 CGI/1.1 那样需要单独的通道。

FastCGI 应用扮演了多个角色。最熟悉的是响应者的角色，在这种场景中，应用接收 HTTP 请求相关的所有信息并生成 HTTP 响应；这就是 CGI/1.1 程序要做的。第二种角色是授权者，这种情况下应用接收所有 HTTP 相关的信息并生成一个授权/未授权的结果；第三种角色是 Filter，它是将请求数据流加上 Web 服务器上文件的数据流进行“过滤”并将结果作为响应。该框架是可扩展的，因此后续可以定义更多基于 FastCGI 的应用。

在这个规范的其他部分中，只要不会引起混淆，术语“FastCGI application”，“application process”，或者“application server”都会被简称为“应用”。

## 2. Initial process state
## 2.初始化进程状态
### 2.1. Argument list
### 2.1 参数列表
By default the Web server creates an argument list containing a single element, the name of the application, taken to be the last component of the executable’s path name. The Web server may provide a way to specify a different application name, or a more elaborate argument list.
* 通过默认的Web服务器创建一个包含单一元素的参数列表，应用的名称，成为可执行的路径名称最后一部分。
* 通过 Web 服务提供一个指向不同的应用名称的方式，或者一个更详细的参数列表

Note that the file executed by the Web server might be an interpreter file (a text file that starts with the characters #!), in which case the application’s argument list is constructed as described in the execve manpage.
* 通过 Web 服务记录执行文件可能是一个解释器文件（一个以标识符 `#!` 开始的文本文件），这种情况下，应用的参数列表就像之前描述的那样在可执行手册中被创建。

2.2. File descriptors
The Web server leaves a single file descriptor, `FCGI_LISTENSOCK_FILENO`, open when the application begins execution. This descriptor refers to a listening socket created by the Web server.

`FCGI_LISTENSOCK_FILENO` equals `STDIN_FILENO`. The standard descriptors `STDOUT_FILENO` and `STDERR_FILENO` are closed when the application begins execution. A reliable method for an application to determine whether it was invoked using CGI or FastCGI is to call `getpeername(FCGI_LISTENSOCK_FILENO)`, which returns -1 with `errno` set to `ENOTCONN` for a FastCGI application.

The Web server’s choice of reliable transport, Unix stream pipes (`AF_UNIX`) or TCP/IP (`AF_INET`), is implicit in the internal state of the `FCGI_LISTENSOCK_FILENO` socket.

### 2.3. Environment variables
The Web server may use environment variables to pass parameters to the application. This specification defines one such variable, `FCGI_WEB_SERVER_ADDRS`; we expect more to be defined as the specification evolves. The Web server may provide a way to bind other environment variables, such as the `PATH` variable.

### 2.4. Other state
The Web server may provide a way to specify other components of an application’s initial process state, such as the priority, user ID, group ID, root directory, and working directory of the process.

## 3. Protocol basics
### 3.1. Notation
We use C language notation to define protocol message formats. All structure elements are defined in terms of the `unsigned char` type, and are arranged so that an ISO C compiler lays them out in the obvious manner, with no padding. The first byte defined in the structure is transmitted first, the second byte second, etc.

We use two conventions to abbreviate our definitions.

First, when two adjacent structure components are named identically except for the suffixes “`B1`” and “`B0`”, it means that the two components may be viewed as a single number, computed as `B1<<8 + B0`. The name of this single number is the name of the components, minus the suffixes. This convention generalizes in an obvious way to handle numbers represented in more than two bytes.

Second, we extend C `structs` to allow the form

```c
struct {
    unsigned char mumbleLengthB1;
    unsigned char mumbleLengthB0;
    ... /* other stuff */
    unsigned char mumbleData[mumbleLength];
};
```

meaning a structure of varying length, where the length of a component is determined by the values of the indicated earlier component or components.

### 3.2. Accepting transport connections
A FastCGI application calls `accept()` on the socket referred to by file descriptor `FCGI_LISTENSOCK_FILENO` to accept a new transport connection. If the `accept()` succeeds, and the `FCGI_WEB_SERVER_ADDRS` environment variable is bound, the application application immediately performs the following special processing:

- `FCGI_WEB_SERVER_ADDRS`: The value is a list of valid IP addresses for the Web server.

If `FCGI_WEB_SERVER_ADDRS` was bound, the application checks the peer IP address of the new connection for membership in the list. If the check fails (including the possibility that the connection didn’t use TCP/IP transport), the application responds by closing the connection.

`FCGI_WEB_SERVER_ADDRS` is expressed as a comma-separated list of IP addresses. Each IP address is written as four decimal numbers in the range `[0..255]` separated by decimal points. So one legal binding for this variable is `FCGI_WEB_SERVER_ADDRS=199.170.183.28,199.170.183.71`.

An application may accept several concurrent transport connections, but it need not do so.

### 3.3. Records
Applications execute requests from a Web server using a simple protocol. Details of the protocol depend upon the application’s role, but roughly speaking the Web server first sends parameters and other data to the application, then the application sends result data to the Web server, and finally the application sends the Web server an indication that the request is complete.

All data that flows over the transport connection is carried in _FastCGI_ records. FastCGI records accomplish two things. First, records multiplex the transport connection between several independent FastCGI requests. This multiplexing supports applications that are able to process concurrent requests using event-driven or multi-threaded programming techniques. Second, records provide several independent data streams in each direction within a single request. This way, for instance, both `stdout` and `stderr` data can pass over a single transport connection from the application to the Web server, rather than requiring separate connections.

```c
typedef struct {
    unsigned char version;
    unsigned char type;
    unsigned char requestIdB1;
    unsigned char requestIdB0;
    unsigned char contentLengthB1;
    unsigned char contentLengthB0;
    unsigned char paddingLength;
    unsigned char reserved;
    unsigned char contentData[contentLength];
    unsigned char paddingData[paddingLength];
} FCGI_Record;
```

A FastCGI record consists of a fixed-length prefix followed by a variable number of content and padding bytes. A record contains seven components:

- `version`: Identifies the FastCGI protocol version. This specification documents FCGI_VERSION_1.

- `type`: Identifies the FastCGI record type, i.e. the general function that the record performs. Specific record types and their functions are detailed in later sections.

- `requestId`: Identifies the FastCGI request to which the record belongs.

- `contentLength`: The number of bytes in the contentData component of the record.

- `paddingLength`: The number of bytes in the paddingData component of the record.

- `contentData`: Between 0 and 65535 bytes of data, interpreted according to the record type.

- `paddingData`: Between 0 and 255 bytes of data, which are ignored.

We use a relaxed C `struct` initializer syntax to specify constant FastCGI records. We omit the `version` component, ignore padding, and treat `requestId` as a number. Thus `{FCGI_END_REQUEST, 1, {FCGI_REQUEST_COMPLETE,0}}` is a record with `type == FCGI_END_REQUEST, requestId == 1, and contentData == {FCGI_REQUEST_COMPLETE,0}`.

#### Padding
The protocol allows senders to pad the records they send, and requires receivers to interpret the `paddingLength` and skip the `paddingData`. Padding allows senders to keep data aligned for more efficient processing. Experience with the X window system protocols shows the performance benefit of such alignment.

We recommend that records be placed on boundaries that are multiples of eight bytes. The fixed-length portion of a `FCGI_Record` is eight bytes.

#### Managing request IDs
The Web server re-uses FastCGI request IDs; the application keeps track of the current state of each request ID on a given transport connection. A request ID R becomes active when the application receives a record `{FCGI_BEGIN_REQUEST, R, ...}` and becomes inactive when the application sends a record `{FCGI_END_REQUEST, R, ...}` to the Web server.

While a request ID `R` is inactive, the application ignores records with `requestId == R`, except for `FCGI_BEGIN_REQUEST` records as just described.

The Web server attempts to keep FastCGI request IDs small. That way the application can keep track of request ID states using a short array rather than a long array or a hash table. An application also has the option of accepting only one request at a time. In this case the application simply checks incoming `requestId` values against the current request ID.

Types of record types
There are two useful ways of classifying FastCGI record types.

The first distinction is between _management_ records and _application_ records. A management record contains information that is not specific to any Web server request, such as information about the protocol capabilities of the application. An application record contains information about a particular request, identified by the `requestId` component.

Management records have a `requestId` value of zero, also called the null request ID. Application records have a nonzero `requestId`.

The second distinction is between _discrete_ and _stream_ records. A discrete record contains a meaningful unit of data all by itself. A stream record is part of a stream, i.e. a series of zero or more non-empty records (`length != 0`) of the stream type, followed by an empty record (`length == 0`) of the stream type. The `contentData` components of a stream’s records, when concatenated, form a byte sequence; this byte sequence is the value of the stream. Therefore the value of a stream is independent of how many records it contains or how its bytes are divided among the non-empty records.

These two classifications are independent. Among the record types defined in this version of the FastCGI protocol, all management record types are also discrete record types, and nearly all application record types are stream record types. But three application record types are discrete, and nothing prevents defining a management record type that’s a stream in some later version of the protocol.

### 3.4. Name-Value pairs