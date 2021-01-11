package main

import (
	"bufio"
	"fmt"
	"os"
)

/*func remove(s []int, i int) []int {
	temp :=  s[:i]
	return s[:len(s)-1]
}*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	/*s := []int{1, 2, 3, 4, 5, 6}
	i := 3
	fmt.Println(s)
	temp := s[:i-1]
	fmt.Println(temp)
	temp1 := s[i:]
	fmt.Println(temp1)
	result := append(temp, temp1...)
	fmt.Println(result)*/
	fmt.Println("Enter the string")
	temp2, _ := reader.ReadString('"')
	fmt.Println(temp2)
	fmt.Println(temp2)
	fmt.Println("Hello")

}
