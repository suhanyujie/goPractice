package main

import (
	"fmt"
	"os"
	"strconv"
	"practice/goLanguageCode/chatper1/calc/simplemath"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[1] {
	case "add"://加法
		if len(args) != 4 {
			fmt.Println("Usage: calc add <int1> <int2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("Usage: calc add <int1> <int2>")
			return
		}
		ret := simplemath.Add(v1,v2);
		fmt.Printf("%d + %d = %d\n", v1,v2,ret)

	case "mul"://乘法
		if len(args) != 4 {
			fmt.Println("Usage: calc mul <int1> <int2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("Usage: calc mul <int1> <int2>")
			return
		}
		ret := simplemath.Multiple(v1,v2);
		fmt.Printf("%d * %d = %d\n", v1,v2,ret)
	}
}
