package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(name, s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Printf("I'm %s(NO.%d) waving: %s\n", name, i, s)
	}
}

func main() {
	// By default, Go programs run with GOMAXPROCS set to the number of cores
	// available; in prior releases it defaulted to 1.
	runtime.GOMAXPROCS(2)
	go say("Sayer01", "world")
	say("Sayer02", "hello")
	time.Sleep(time.Second)
}
