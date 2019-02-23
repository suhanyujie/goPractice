package main

import "fmt"

func main() {
	n,_ := Read([]byte("hello world..."))
	fmt.Println(n)
}

func Read(b []byte) (n int, err error) {
	n = 1
	err = nil
	return
}


/**
## 多返回值

* 现有函数原型如下：

```
func (file *File) Read(b []byte) (n int, err Error)
```

* 其中的`n`,`err`是返回值名称，他们的值在函数开始的时候被自动初始化为空
* 在函数内部时，可以分别对其赋值


 */

