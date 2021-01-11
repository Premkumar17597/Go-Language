package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Time : %v, Error Message : %v", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "Failed to Connect"}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}

}
