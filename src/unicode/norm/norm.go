package main

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func main() {
	var (
		n   int
		r   rune
		it  norm.Iter
		out []byte
	)
	in := []byte(`test`)
	fmt.Printf("%s\n", in)
	fmt.Println(in)
	it.Init(norm.NFD, in)
	for !it.Done() {
		ruf := it.Next()
		r, n = utf8.DecodeRune(ruf)
		fmt.Printf("bytes read: %d. val: %q\n", n, r)
		buf := make([]byte, utf8.RuneLen(r))
		utf8.EncodeRune(buf, r)
		out = norm.NFC.Append(out, buf...)
	}
	fmt.Printf("%s\n", out)
	fmt.Println(out)
}
