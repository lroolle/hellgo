package main

import (
	"fmt"
	"sync"
)

func PrintNums(numChan, charChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(2*i + j + 1)
		}
		<-numChan
		charChan <- true
	}
}

func PrintChars(numChan, charChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		numChan <- true
		for j := 0; j < 2; j++ {
			fmt.Printf("%c", 'A'+(2*i+j))
		}
		<-charChan
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	numChan := make(chan bool)
	charChan := make(chan bool)
	go PrintNums(numChan, charChan, &wg)
	go PrintChars(numChan, charChan, &wg)

	// Wait util 2 goroutines finished print
	wg.Wait()
}
