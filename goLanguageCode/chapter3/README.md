## 第3章 面向对象编程
* 了解一下go的非侵入式接口
* 对比一下其他语言的侵入式的接口设计

```
type IFile interface {
    Read(buf []byte) (n int, err error)
    Write(buf []byte) (n int, err error)
    Seek(off int64, whence int) (pos int64, err error)
    Close() error
}

type IReader interface {
    Read(buf []byte) (n int, err error)
}

type IWriter interface {
    Write(buf []byte) (n int, err error)
}

type ICloser interface {
    Close() error
}
```

* 尽管File类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是File类实现了 这些接口，可以进行赋值：

```
var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWriter = new(File)
var file4 ICloser = new(File)
```

