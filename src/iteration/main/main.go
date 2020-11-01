package main

import (
	"fmt"
)

func main() {
	s := []int{0}
	for i, v := range s {
		s = append(s, i+1)
		fmt.Println(v)
		if i > 10 {
			break
		}
	}
	fmt.Println(s)
}
