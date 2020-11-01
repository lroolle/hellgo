package main

import (
	"fmt"
	"strings"
)

func main() {
	// ZERO-VALUE:
	//
	// It's ready to use from the get-go.
	// You don't need to initialize it.
	var str strings.Builder

	for i := 0; i < 10; i++ {
		str.WriteString("a")
	}

	fmt.Println(str.String())
}
