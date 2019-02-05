package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{"a1":1,"a2":2,"a3":3}
	delete(map1, "a2")
	if val,ok := map1["a2"];ok {
		fmt.Println("值存在")
		fmt.Println(val)
	} else {
		fmt.Println("值不存在")
	}

	//fmt.Println(map1)
}


/**
## map

###  map的声明
* 声明方式1：

```
var map1 map[string]int
```

* 声明map时，可以使用`make`，并指定map的容量：

```
map1 := make(map[string]int, 100)
```

### map删除
* 使用delete函数：

```
delete(map1, "22")
```

* 注意：因为map类型是引用类型，所以使用delete后，无需再次进行赋值
* map中常用的判断值是否存在，可以使用如下的示例：

```
if val,ok := map1["a2"];ok {
	fmt.Println("值存在")
	fmt.Println(val)
} else {
	fmt.Println("值不存在")
}
```

*/

