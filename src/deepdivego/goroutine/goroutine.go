package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d", i)
			}
		}(i)
		// 这里 i 如果在里面引用不安全，在这里传入；
	}
}
