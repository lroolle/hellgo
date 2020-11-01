package main

import (
	"fmt"
	"time"
)

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}() // Note the parentheses - must call the function.
}

func main() {
	Announce("Hi from goroutine", 1000)
	time.Sleep(time.Millisecond)
}
