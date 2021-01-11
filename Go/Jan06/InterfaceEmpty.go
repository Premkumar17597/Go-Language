package main

import "fmt"

type I interface{}

func main() {
	//var i interface{}
	var i I
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
