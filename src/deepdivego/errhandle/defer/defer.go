package main

import "fmt"

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error")
	defer fmt.Println(5)
	return
	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 5 {
			panic("too many")
		}
	}
}

func main() {
	tryDefer2()
	tryDefer()
}
