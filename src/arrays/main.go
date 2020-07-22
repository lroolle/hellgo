package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 3 parts:
	var a = [6]int{2, 3, 5, 7, 11, 13}
	var as = a[1:4]
	var ass = a[4:]
	// var s = make([]int, 1, 8)
	fmt.Printf("as: %v cap(as): %v, ass: %v cap(ass): %v\n", as, cap(as), ass, cap(ass))
	as[1] = 4
	ass[1] = 114
	ass = append(ass, 115, 116, 117, 118, 119)
	fmt.Printf("as: %v cap(as): %v, ass: %v cap(ass): %v, a: %v\n", as, cap(as), ass, cap(ass), a)
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(ass).Kind())
}
