package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	s string
}

type S struct {
	n int64
}

func (t T) M() {
	fmt.Println(t.s)
}

func (t S) M() {
	fmt.Println(t.n)
}

func main() {
	var i I = T{"Hello World"}
	var a I
	var m I = S{10}
	a = i
	a.M()
	a = m
	a.M()

}
