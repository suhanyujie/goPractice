## 第3章 面向对象编程
### 匿名组合
* 看如下一段代码：

```
type X struct {
    Name string
}
type Y struct {
    X
    Name string
}
```

* 组合的类型和被组合的类型都包含一个Name成员
* 这样组合是合法的，所有的Y类型的Name成员的访问都只会访问到最外层的Name变量
* X.Name相当于被隐藏起来

### 可见性
* Go语言中符号的可访问性是包一级的，而不是类型一级的
* 如果要使某个符号对其他包可见，需要将该符号定义为以大写字母开头

### 接口
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

* Go语言的非侵入式接口，看似只是做了很小的文法调整，实则影响深远。
* 接口赋值并不要求两个接口必须等价。如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A。

#### 接口查询
* 在Go语言中，你可以询问接口它指向的对象是否是某个类型

```
var file1 Writer = ...
if file6,ok:=file1.(*File);ok {
    ...
}
```

#### 类型查询
* 在Go语言中，可以直接了当地查询接口指向的对象实例得类型：

```
var v1 interface{} = ...
switch v:= v1.(type) {
    case int:
    case string:
    ...
}
```


