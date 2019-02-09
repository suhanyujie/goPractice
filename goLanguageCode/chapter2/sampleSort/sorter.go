package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i","infile","File contains value for sorting")
var outfile *string = flag.String("o","outfile","File to receive sorted values")
var algorithm *string = flag.String("a","qsort","Sort algorithm")


func main() {
	flag.Parse()

	if infile != nil {
		 fmt.Println("infile = ", *infile,"outfile = ",*outfile, "algorithm = ", *algorithm)
	}


}

/**
## 排序算法
* 程序用法如下：

```
USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>
```




 */

