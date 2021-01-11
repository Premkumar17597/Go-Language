// This method should use only small functions as below. Otherwice, It will affect the readablity of the program.

package main

import "fmt"

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	z = 2
	return
}

func main() {
	fmt.Println(split(17))
}
