package main

import (
	"fmt"
	"math"
)

type Absr interface {
	//Abs() float64
	Area() float64
	Area1() float64
}

type Vertex struct {
	x, y float64
}

type Vertex1 struct {
	x, y float64
}
type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Area() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (g *Vertex) Area1() float64 {
	return float64(g.x * g.y)
}

func main() {
	//f := Myfloat(-2)
	v := Vertex{3, 4}
	//b := Vertex1{5, 4}
	var a Absr
	//a = f
	//fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Area())
	//a = &b
	fmt.Println(a.Area1())
	//fmt.Println(a.Abs())

}
