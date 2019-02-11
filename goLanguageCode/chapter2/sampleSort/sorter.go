package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
	"strconv"
	"practice/goLanguageCode/chapter2/sampleSort/algorithms/bubble"
)

var infile *string = flag.String("i", "infile", "File contains value for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
	}
	values, err := readValue(*infile)
	if err != nil {
		log.Println(err)
	}
	err = bubble.BubbleSort(values)
	if err !=nil {
		log.Println(err)
	}

	fmt.Println(values)
}

//从文件中读取数据
func readValue(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		log.Println("Failed to open the input file", err)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("a too long line,seems unexpected.")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

//将处理后的数据写入到文件中
func writeValues(values []int, outfile string) (err error) {
	file, err := os.Open(outfile)
	if err != nil {
		return
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str)
	}
	return nil
}

/**
## 排序算法
* 程序用法如下：

```
USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>
```

## 注意
* main函数中`err = bubble.BubbleSort(values)`的调用排序函数后，无需返回数值，因为切片是引用类型，在函数内部排序后就直接影响了原切片数组



 */
