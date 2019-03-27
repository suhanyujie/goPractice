package main

/*
#include <stdlib.h>
#include <stdio.h>

int func1123() {
	printf("this is a test function...\n");
	return 123;
}
*/
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func main() {
	fmt.Println(Random())
	fmt.Println(C.func1123())
}
