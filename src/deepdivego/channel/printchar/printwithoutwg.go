package main

import (
	"fmt"
)

func PrintNums(numChan, charChan, done chan bool) {
	for i := 1; i < 8; i += 2 {
		fmt.Print(i)
		fmt.Print(i + 1)
		<-numChan
		charChan <- true
	}
	defer func(done chan bool) {
		done <- true
	}(done)
	// defer close(done)
}

func PrintChars(numChan, charChan, done chan bool) {
	for i := 'A'; i < 'H'; i += 2 {
		numChan <- true
		fmt.Print(string(i))
		fmt.Print(string(i + 1))
		<-charChan
	}
	defer close(done)
}

func main() {
	var numChan = make(chan bool)
	var charChan = make(chan bool)
	var done = make(chan bool)
	go PrintNums(numChan, charChan, done)
	go PrintChars(numChan, charChan, done)

	// The unbuffered done blocks until received a value
	// Or until the done channel closed because:
	//   a receive operation on a closed channel can always proceed immediately,
	<-done
}
