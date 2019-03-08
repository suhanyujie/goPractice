package main

import "fmt"

type Human struct {
	Head interface{}
	Hand interface{}
}

type User struct {
	Human
	Name string
}

func main() {
	var1 := &User{Human{"head", "hand"}, "samuel"}
	fmt.Printf("%V", var1)
}
