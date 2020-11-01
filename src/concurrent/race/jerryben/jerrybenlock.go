package main

import (
	"fmt"
	"sync"
)

type IceCreamMaker interface {
	// Hello greets a customer
	Lock()
	Unlock()
	Hello()
}
type Ben struct {
	name string
	mu   sync.Mutex
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
	if b.name != "Ben" {
		panic("Data Race: I'm not ben!")
	}
}
func (b *Ben) Lock() {
	b.mu.Lock()
}
func (b *Ben) Lock() {
	b.mu.Lock()
}

type Jerry struct {
	name string
	mu   sync.Mutex
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
	if j.name != "Jerry" {
		panic("Data Race: I'm not jerry!")
	}
}

func main() {
	var ben = &Ben{name: "Ben"}
	var jerry = &Jerry{name: "Jerry"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		sync.Mutex.Lock()
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}
}
