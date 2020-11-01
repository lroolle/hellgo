// https://stackoverflow.com/questions/11117382/priority-in-go-select-statement-workaround

package main

import "fmt"

func sender(out chan int, exit chan bool) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	exit <- true
}

func main() {
	out := make(chan int, 10)
	exit := make(chan bool)

	go sender(out, exit)

L:
	for {
		select {
		case i := <-out:
			fmt.Printf("Value: %d\n", i)
		case <-exit:
			fmt.Println("Exiting")
			break L
		}
	}
	fmt.Println("Did we get all 10? Most likely not")
}
