package main

import "fmt"

func main() {
	s := []interface{}{"one", nil}
	s[1] = s // fmt.Println(s)  will cause max recursion exceed to stack overflow

	fmt.Println(s[0])

	s2 := s[1].([]interface{})
	fmt.Println(s2[0])

	s3 := s2[1].([]interface{})
	fmt.Println(s3[0])

	traverse(s)
}

func traverse(s []interface{}) {
	s1 := s[1].([]interface{})
	fmt.Println(s1[0])
	traverse(s1)
}
