package main

import "fmt"

func main() {
	slice1 := []int{1,2,445}
	slice2 := []int{12,22,145,13,24}
	copy(slice2,slice1)
	fmt.Println(slice2)
}

/**
## slice切片

### 切片的复制
* 使用函数`copy`函数：

```
copy(slice2,slice1)
```

* 因为切片是引用的方式，所以无需在复制后再次赋值
* 复制切片时，注意两个变量的大小，所能复制的是两个切片中数量较小的那个，如下从slice1中复制到slice2中，只能将slice1中的3个复制到slice2中

 */