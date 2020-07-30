package main

import (
	"fmt"
	"unicode"

	"github.com/lroolle/deepdivego/testing/substring"
)

func main() {
	fmt.Println(substring.LongestSubstring("   "))
	var mixed = "\b5Ὂg̀9! ℃ᾭG"
	fmt.Printf("Mixed Type is: %T", mixed)
	for _, c := range []rune(mixed) {
		fmt.Printf("%q\n", c)
		unicode.IsLower(c)
	}
}
