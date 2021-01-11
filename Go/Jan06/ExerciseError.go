package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func squareroot(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError{
			time.Now(),
			"it didn't work",
		}
	}
	return math.Sqrt(x), nil
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Time : %v, Error Message : %v ", e.When, e.What)
}

func main() {

	var x float64 = 2
	i, err := squareroot(x)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

}
