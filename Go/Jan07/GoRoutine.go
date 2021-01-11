package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, time.Now(), i)
	}
}

func main() {

	go say("hello")

	say("Hi   ")
}
