package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	var i1, i2 int
	fmt.Scanln(&s1, &s2)
	fmt.Println(s1, s2)

	n, err := fmt.Scanf("%d %d", &i1, &i2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, i1, i2)
}
