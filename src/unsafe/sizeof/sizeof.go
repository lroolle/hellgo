package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s string
	var c complex128
	var a [3]uint32
	var u16 uint16
	var u32 uint32
	fmt.Printf("s=%q, SizeOf(s)=%v\n", s, unsafe.Sizeof(s))
	fmt.Printf("c=%v, SizeOf(c)=%v\n", c, unsafe.Sizeof(c))
	fmt.Printf("a=%v, SizeOf(a)=%v\n", a, unsafe.Sizeof(a))
	fmt.Printf("u16=%v, SizeOf(u16)=%v\n", u16, unsafe.Sizeof(u16))
	fmt.Printf("u32=%v, SizeOf(u32)=%v\n", u32, unsafe.Sizeof(u32))

	type S struct {
		a uint16 // size 2, alignment 2, padding to 4
		b uint32 // 4
		s string // 16
	}
	var st S
	fmt.Printf("st=%q, SizeOf(st)=%v\n", st, unsafe.Sizeof(st))

	// Empty Struct
	type ES struct{}
	es := ES{}
	// es := struct{}{}
	fmt.Printf("es=%q, SizeOf(es)=%v\n", es, unsafe.Sizeof(es))

	// Embedded Empty Struct
	type EES struct {
		es  ES
		es2 struct{}
		es3 struct{}
	}
	var ees = EES{}
	fmt.Printf("ees=%q, SizeOf(ees)=%v\n", ees, unsafe.Sizeof(ees))

	// List of empty struct
	var les [1000000000]EES
	fmt.Printf("len(les)=%d, SizeOf(les)=%v\n", len(les), unsafe.Sizeof(les))
}
