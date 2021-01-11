package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Printf("Type of %v is %T\n", World, World)
	fmt.Printf("Happy %v Day and the type is %T\n", Pi, Pi)

	const Truth = true
	fmt.Printf("Go rules? and the type is %T\n", Truth, Truth)
}
