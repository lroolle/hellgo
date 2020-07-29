
# Table of Contents

1.  [Formatting](#org8d160ff)
2.  [Commentary](#org9c9ff02)
3.  [Names](#org5f93c94)
    1.  [Package Names](#orgb7015a9)
    2.  [Getters](#org2da569a)
    3.  [Interface names](#orgf4ceed7)
    4.  [MixedCaps](#orgcc18255)
4.  [Semicolons](#org7d5f828)
5.  [Control structures](#org0cfed3d)
    1.  [If](#org447134c)
    2.  [Redeclaration and reassignment](#org539de00)
    3.  [For](#org7e27607)
    4.  [Switch](#org1af5b79)
6.  [Functions](#org2bd594d)
    1.  [Multiple return values](#orgd6424dd)
    2.  [Named result parameters](#org542d94e)
    3.  [Defer](#orga0b0f0d)
7.  [Data](#org528cecf)
    1.  [Allocation with `new`](#org7abddf6)
    2.  [Constructors and composite literals](#orge13ab2a)
    3.  [Allocation with `make`](#orgfb471a9)
    4.  [Arrays](#orga85c840)
    5.  [Slices](#org41d6bbd)
    6.  [Two-dimensional slices](#org8263566)
        1.  [Slice internals (Go Slices: usage and internals)](#orgefea994):ATTACH:
        2.  [Growing slices (the `copy` and `append` functions)](#org03062fb)
        3.  [A possible &ldquo;gotcha&rdquo;](#org79b29aa)
    7.  [Map](#orgccf7f58)
    8.  [Printing](#org995d932)
    9.  [Append](#org545ef71)
8.  [Initialization](#org5fe8a97)
    1.  [Constants](#orgcf94293)
    2.  [Variables](#org4393214)
    3.  [The init Function](#org8a071fe)
9.  [Methods](#org7bfbc72)
    1.  [Pointers vs. Values](#orgb6650c3)
10. [Interfaces and other types](#orga5fe412)
    1.  [Interfaces](#orgba96913)
    2.  [Conversions](#orgfc5a6ce)
    3.  [Interface conversions and type assertions](#orgd5a9920)
    4.  [Generality](#org93d2316)
    5.  [Interfaces and methods](#orgf7b30b8)
11. [The blank identifier](#org5735754)
    1.  [Unused imports and variables](#orga7a5e4a)
    2.  [Import for side effect](#org02bc32e)
    3.  [Interface checks](#org20f76a0)
12. [Embedding](#org475d19f)
13. [Concurrency](#org0ad0522)
    1.  [Share by communicating](#org969c2d3)
14. [Errors](#orga1e8a30)
    1.  [By convention](#orgd023c62)
    2.  [Panic](#org99e21fc)
    3.  [Recover](#orgf3c3c3c)
15. [A Web Server](#orgdefac4e)
16. [Go Docs](#org575871f)
17. [References](#orgb0fa0a9):W:



<a id="org8d160ff"></a>

# Formatting

-   Indentation
    
    We use tabs for indentation and gofmt emits them by default. Use spaces only if
    you must.

-   Line length
    
    Go has no line length limit. Don&rsquo;t worry about overflowing a punched card. If a
    line feels too long, wrap it and indent with an extra tab.

-   Parentheses
    
    Go needs fewer parentheses than C and Java: control structures (if, for, switch)
    do not have parentheses in their syntax. Also, the operator precedence hierarchy
    is shorter and clearer, so
    
        x<<8 + y<<16

means what the spacing implies, unlike in the other languages.


<a id="org9c9ff02"></a>

# Commentary

Line comments(`//`) are the norm; block comments(`/**/`)appear mostly as package
comments, but are useful within an expression or to disable large swaths of
code.

-   Package Comment
    Every package should have a package comment, a block comment preceding the package clause.

    /*
    Package regexp implements a simple library for regular expressions.
    
    The syntax of the regular expressions accepted is:
    
        regexp:
            concatenation { '|' concatenation }
        concatenation:
            { closure }
        closure:
            term [ '*' | '+' | '?' ]
        term:
            '^'
            '$'
            '.'
            character
            '[' [ '^' ] character-ranges ']'
            '(' regexp ')'
    */
    package regexp

If the package is simple, the package comment can be brief.

    // Package path implements utility routines for
    // manipulating slash-separated filename paths.


<a id="org5f93c94"></a>

# Names


<a id="orgb7015a9"></a>

## Package Names

-   the package name should be good: **short**, **concise**, **evocative**.
-   packages are given lower case, single-word names; there should be no need for
    underscores or mixedCaps
-   package name **is the base name of its source directory**; the package in
    src/encoding/base64 is imported as &ldquo;encoding/base64&rdquo; but has name base64, not
    encoding<sub>base64</sub> and not encodingBase64.
-   the buffered reader type in the bufio package is called Reader, not BufReader,
    because users see it as bufio.Reader, Reader does not conflict with io.Reader
-   once.Do; once.Do(setup) reads well and would not be improved by writing once.DoOrWaitUntilDone(setup).
    ring.New.(not NewRing)

**Use the package structure to help you choose good names.**


<a id="org2da569a"></a>

## Getters

If you have a field called owner (lower case, unexported), the getter method
should be called Owner (upper case, exported), not GetOwner

    // The use of upper-case names for export provides the hook to discriminate the field from the method.
    owner := obj.Owner()
    if owner != user {
        obj.SetOwner(user)
    }


<a id="orgf4ceed7"></a>

## Interface names

-   By convention, one-method interfaces are named by the method name plus an `-er`
    suffix or similar modification to construct an agent noun: `Reader`, `Writer`,
    `Formatter`, `CloseNotifier` etc.
-   call your string-converter method String not ToString


<a id="orgcc18255"></a>

## MixedCaps

Finally, the convention in Go is to use `MixedCaps` or `mixedCaps` rather than
<del>under<sub>scores</sub></del> to write multiword names.


<a id="org7d5f828"></a>

# Semicolons

“if the newline comes after a token that could end a statement, insert a semicolon”.

    if i < f()  // wrong!
    {           // wrong!
        g()
    }


<a id="org0cfed3d"></a>

# Control structures


<a id="org447134c"></a>

## If

-   statement initialization

    if err := file.Chmod(0664); err != nil {
        log.Print(err)
        return err
    }

-   no `else` needed

    f, err := os.Open(name)
    if err != nil {
        return err
    }
    d, err := f.Stat()
    if err != nil {
        f.Close()
        return err
    }
    codeUsing(f, d)


<a id="org539de00"></a>

## Redeclaration and reassignment

    f, err := os.Open(name)
    // This duplication is legal: err is declared by the first statement,
    // but only re-assigned in the second.
    d, err := f.Stat()


<a id="org7e27607"></a>

## For

-   Three forms

    // Like a C for
    for init; condition; post { }
    
    // Like a C while
    for condition { }
    
    // Like a C for(;;)
    for { }

-   For range

    for key := range m {
        if key.expired() {
            delete(m, key)
        }
    }
    // for _, value := range m {}

-   For `String`: the range does more work for you

    import "fmt"
    
    func main() {
    	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
    	}
    }

    character U+65E5 '日' starts at byte position 0
    character U+672C '本' starts at byte position 3
    character U+FFFD '�' starts at byte position 6
    character U+8A9E '語' starts at byte position 7

-   if you want to run multiple variables in a for you should use parallel
    assignment (although that precludes ++ and &#x2013;)

    // Reverse a
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }


<a id="org1af5b79"></a>

## Switch

-   comma-separated lists

    func shouldEscape(c byte) bool {
        switch c {
        case ' ', '?', '&', '=', '#', '+', '%':
            return true
        }
        return false
    }

-   Break `Switch` in **loop**

    Loop:  // the loop label
    	for n := 0; n < len(src); n += size {
    		switch {
    		case src[n] < sizeOne:
    			if validateOnly {
    				break       // break switch
    			}
    			size = 1
    			update(src[n])
    
    		case src[n] < sizeTwo:
    			if n+1 >= len(src) {
    				err = errShortInput
    				break Loop  // break for
    			}
    			if validateOnly {
    				break
    			}
    			size = 2
    			update(src[n] + src[n+1]<<shift)
    		}
    	}

-   Type `Switch`

    var t interface{}
    t = functionOfSomeType()
    switch t := t.(type) {
    default:
        fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
    case bool:
        fmt.Printf("boolean %t\n", t)             // t has type bool
    case int:
        fmt.Printf("integer %d\n", t)             // t has type int
    case *bool:
        fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
    case *int:
        fmt.Printf("pointer to integer %d\n", *t) // t has type *int
    }


<a id="org2bd594d"></a>

# Functions


<a id="orgd6424dd"></a>

## Multiple return values

-   Return err

    // it returns the number of bytes written and a non-nil error when n != len(b).
    func (file *File) Write(b []byte) (n int, err error)

-   Simple-minded

    func nextInt(b []byte, i int) (int, int) {
    	//...
        return x, i
    }
    
    x, i = nextInt(b, i)


<a id="org542d94e"></a>

## Named result parameters

Result named, they are initialized to the zero values for their types when the
function begins; if the function executes a return statement with no arguments,
the current values of the result parameters are used as the returned values.

    func ReadFull(r Reader, buf []byte) (n int, err error) {
        for len(buf) > 0 && err == nil {
            var nr int
            nr, err = r.Read(buf)
            n += nr
            buf = buf[nr:]
        }
        return  // this return named n, err
    }


<a id="orga0b0f0d"></a>

## Defer

The deferred call&rsquo;s arguments are evaluated immediately, but the function call
is not executed until the surrounding function returns.

    import "fmt"
    
    func main() {
    	defer fmt.Println("world")
    	fmt.Println("hello")
    	for i := 0; i < 5; i++ {
    		defer fmt.Printf("%d ", i)
    	}
    }

    hello
    4 3 2 1 0 world

Deferred functions are executed in **LIFO** order

    import "fmt"
    
    func trace(s string) string {
    	fmt.Println("entering:", s)
    	return s
    }
    
    func un(s string) { fmt.Println("leaving:", s) }
    
    func a() {
    	defer un(trace("a"))
    	fmt.Println("in a")
    }
    
    func b() {
    	defer un(trace("b"))
    	fmt.Println("in b")
    	a()
    }
    
    func main() {
    	b()
    }

    entering: b
    in b
    entering: a
    in a
    leaving: a
    leaving: b


<a id="org528cecf"></a>

# Data


<a id="org7abddf6"></a>

## Allocation with `new`

-> @[golang-spec: make & new](golang-spec.md)

built-in function that allocates memory, but unlike its namesakes in some other
languages it does **not initialize the memory**, it only **zeros it**.


<a id="orge13ab2a"></a>

## Constructors and composite literals

    func NewFile(fd int, name string) *File {
    	if fd < 0 {
    		return nil
    	}
    	// f := File{fd, name, nil, 0}
    	// return &f
    	return &File{fd, name, nil, 0}
    }

    import "fmt"
    
    const (
    	Enone int = 1
    	Eio int = 2
    	Einval int = 3
    )
    
    func main() {
        // the key for array or slice as index
    	a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    	s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    
    	for x, i := range a { fmt.Println(x, i) }
    	for x, i := range s { fmt.Println(x, i) }
    	for x, i := range m { fmt.Println(x, i) }
    }

    0 
    1 no error
    2 Eio
    3 invalid argument
    0 
    1 no error
    2 Eio
    3 invalid argument
    2 Eio
    3 invalid argument
    1 no error


<a id="orgfb471a9"></a>

## Allocation with `make`

It creates *slices*, *maps*, and *channels* **only**, and it returns an
initialized (not zeroed) value of type T (not \*T).

    var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
    // Idiomatic
    var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints


<a id="orga85c840"></a>

## Arrays

There are major differences between the ways arrays work in Go and C. In Go,

-   Arrays are values. Assigning one array to another copies all the elements.
-   In particular, if you pass an array to a function, it will receive a copy of
    the array, not a pointer to it.
-   The size of an array is part of its type. The types [10]int and [20]int are distinct.

    // The value property can be useful but also expensive; if you want C-like behavior
    // and efficiency, you can pass a pointer to the array.
    // But even this style isn't idiomatic Go. Use slices instead.
    
    func Sum(a *[3]float64) (sum float64) {
        for _, v := range *a {
            sum += v
        }
        return
    }
    
    array := [...]float64{7.0, 8.5, 9.1}
    x := Sum(&array)  // Note the explicit address-of operator

-   Arrays do not need to be initialized explicitly; the zero value of an array is a
    ready-to-use array whose elements are themselves zeroed:

    import "fmt"
    
    func main() {
    	var a [3]int
    	a2 := a[2]
    	fmt.Println(a2)
    
    	var s []int
    	// s0 := s[0] // index out of range
    	s = []int{a[2]}
    	s0 := s[0]
    	fmt.Println(s0)
    }

    0
    0


<a id="org41d6bbd"></a>

## Slices

-> @[Go Slices: usage and internals](https://blog.golang.org/slices-intro)

Slices hold references to an underlying array, and if you assign one slice to
another, both refer to the same array.

    // To create a slice given array
    x := [3]string{"Лайка", "Белка", "Стрелка"}
    s := x[:] // a slice referencing the storage of x

    func (f *File) Read(buf []byte) (n int, err error)

    func Append(slice, data []byte) []byte {
        l := len(slice)
        if l + len(data) > cap(slice) {  // reallocate
            // Allocate double what's needed, for future growth.
            newSlice := make([]byte, (l+len(data))*2)
            // The copy function is predeclared and works for any slice type.
            copy(newSlice, slice)
            slice = newSlice
        }
        slice = slice[0:l+len(data)]
        copy(slice[l:], data)
        return slice
    }


<a id="org8263566"></a>

## Two-dimensional slices

    import "fmt"
    
    func main() {
    	type LinesOfText [][]byte
    	text := LinesOfText{
    		[]byte("Now is the time"),
    		[]byte("for all good gophers"),
    		[]byte("to bring some fun to the party."),
    	}
    	fmt.Println(text)
    }

    [[78 111 119 32 105 115 32 116 104 101 32 116 105 109 101] [102 111 114 32 97 108 108 32 103 111 111 100 32 103 111 112 104 101 114 115] [116 111 32 98 114 105 110 103 32 115 111 109 101 32 102 117 110 32 116 111 32 116 104 101 32 112 97 114 116 121 46]]


<a id="orgefea994"></a>

### Slice internals ([Go Slices: usage and internals](https://blog.golang.org/slices-intro))     :ATTACH:

A slice is a descriptor of an array segment. It consists of

-   a pointer to the array
-   the length of the segment
-   and its capacity (the maximum length of the segment).

![img](/Users/eric/G/w/golangtour/img/_20200725_011025slice-struct.png)

Our variable s, created earlier by make([]byte, 5), is structured like this:

![img](/Users/eric/G/w/golangtour/img/_20200725_011002slice-1.png)

    s = s[2:4]

Slicing does not copy the slice&rsquo;s data. It creates a new slice value that points
to the original array. This makes slice operations as efficient as manipulating
array indices. Therefore, modifying the elements (not the slice itself) of a
re-slice modifies the elements of the original slice:

    d := []byte{'r', 'o', 'a', 'd'}
    e := d[2:]
    // e == []byte{'a', 'd'}
    e[1] = 'm'
    // e == []byte{'a', 'm'}
    // d == []byte{'r', 'o', 'a', 'm'}

    s = s[:cap(s)]


<a id="org03062fb"></a>

### Growing slices (the `copy` and `append` functions)

built-in copy function.
As the name suggests, copy copies data from a source slice to a destination
slice. It returns the number of elements copied.

    func copy(dst, src []T) int

    func AppendByte(slice []byte, data ...byte) []byte {
        m := len(slice)
        n := m + len(data)
        if n > cap(slice) { // if necessary, reallocate
            // allocate double what's needed, for future growth.
            newSlice := make([]byte, (n+1)*2)
            copy(newSlice, slice)
            slice = newSlice
        }
        slice = slice[0:n]
        copy(slice[m:n], data)
        return slice
    }
    
    p := []byte{2, 3, 5}
    p = AppendByte(p, 7, 11, 13)
    // p == []byte{2, 3, 5, 7, 11, 13}

Since the zero value of a slice (nil) acts like a zero-length slice, you can
declare a slice variable and then append to it in a loop:

    // Filter returns a new slice holding only
    // the elements of s that satisfy fn()
    func Filter(s []int, fn func(int) bool) []int {
        var p []int // == nil
        for _, v := range s {
            if fn(v) {
                p = append(p, v)
            }
        }
        return p
    }


<a id="org79b29aa"></a>

### A possible &ldquo;gotcha&rdquo;

This code behaves as advertised, but the returned []byte points into an array
containing the entire file. Since the slice references the original array, as
long as the slice is kept around the garbage collector can&rsquo;t release the array;
the few useful bytes of the file keep the entire contents in memory.

    var digitRegexp = regexp.MustCompile("[0-9]+")
    
    func FindDigits(filename string) []byte {
        b, _ := ioutil.ReadFile(filename)
        return digitRegexp.Find(b)
    }

To fix this problem one can copy the interesting data to a new slice before returning it:

    func CopyDigits(filename string) []byte {
        b, _ := ioutil.ReadFile(filename)
        b = digitRegexp.Find(b)
        c := make([]byte, len(b))
        copy(c, b)
        return c
    }


<a id="orgccf7f58"></a>

## Map

    import "fmt"
    
    func main() {
    	var m = map[string]int{
    		"UTC":  0*60*60,
    		"EST": -5*60*60,
    		"CST": -6*60*60,
    		"MST": -7*60*60,
    		"PST": -8*60*60,
    	}
    	// Non-exist Key will return the zero value for the type of the entries in the map
    	non, ok := m["0"]
    	fmt.Println(non, ok)
    
    	offset := func(tz string) int {
    		if seconds, ok := m[tz]; ok {
    			return seconds
    		}
    		fmt.Println("unknown time zone:", tz)
    		return 0
    	}
    	fmt.Println(offset("EST"), offset("0"))
    	// to delete
    	delete(m, "PDT")
    }

    0 false
    unknown time zone: 0
    -18000 0


<a id="org995d932"></a>

## Printing

-   %v %+v %#v %q %x

    import "fmt"
    
    func main() {
    	type T struct {
    		a int
    		b float64
    		c string
    	}
    	t := &T{ 7, -2.35, "abc\tdef" }
    	fmt.Printf(" %%v: %v \n", t)
    	fmt.Printf("%%+v: %+v //+v annotates the fields of the structure with their names.\n", t)
    	fmt.Printf("%%#v: %#v //#v prints the value in full Go syntax.\n", t)
    	fmt.Printf(" %%q: %q  //q applies to int and runes producing a single-quoted rune constant.\n", t)
    	fmt.Printf(" %%x: %x  //x applies to int and runes\n", t)
    	fmt.Printf("%%#q: %#q //#q prints the value in full Go syntax.\n", t)
    }

    %v: &{7 -2.35 abc	def} 
    %+v: &{a:7 b:-2.35 c:abc	def} //+v annotates the fields of the structure with their names.
    %#v: &main.T{a:7, b:-2.35, c:"abc\tdef"} //#v prints the value in full Go syntax.
     %q: &{'\a' %!q(float64=-2.35) "abc\tdef"}  //q applies to int and runes producing a single-quoted rune constant.
     %x: &{7 -0x1.2cccccccccccdp+01 61626309646566}  //x applies to int and runes
    %#q: &{'\a' %!q(float64=-2.35000) `abc	def`} //#q prints the value in full Go syntax.

-   Custom method to control the print

    import "fmt"
    
    type T struct {
    	a int
    	b float64
    	c string
    }
    
    // a method with the signature String() string on the type.
    // this example used a pointer because that's more efficient and idiomatic for
    // struct types.
    func (t *T) String() string {
    	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
    }
    
    func main() {
    	t := &T{ 7, -2.35, "abc\tdef" }
    	fmt.Printf("%v\n", t)
    }

    7/-2.35/"abc\tdef"

If you need to print values of type T as well as pointers to T, the receiver for
String must be of value type;

    type MyString string
    
    func (m MyString) String() string {
        return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
    }
    
    // to fix
    func (m MyString) String() string {
        return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
    }

-   The signature of `Printf` uses the type `...interface{}` for its final argument to
    specify that an arbitrary number of parameters (of arbitrary type) can appear
    after the format.

    // v acts like a variable of type []interface{}
     func Printf(format string, v ...interface{}) (n int, err error) {}
    
    // Fprintln formats using the default formats for its operands and writes to w.
    // Spaces are always added between operands and a newline is appended.
    // It returns the number of bytes written and any write error encountered.
    func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
    	p := newPrinter()
    	p.doPrintln(a)
    	n, err = w.Write(p.buf)
    	p.free()
    	return
    }
    
    // Println formats using the default formats for its operands and writes to standard output.
    // Spaces are always added between operands and a newline is appended.
    // It returns the number of bytes written and any write error encountered.
    func Println(a ...interface{}) (n int, err error) {
    	return Fprintln(os.Stdout, a...)
    }
    // to tell the compiler to treat v as a list of arguments; otherwise it would just pass v as a single slice argument.

-   a `...` parameter can be of a specific type, for instance `...int` for a min
    function that chooses the least of a list of integers:

    import "fmt"
    
    func Min(a ...int) int {
        min := int(^uint(0) >> 1)  // largest int
        for _, i := range a {
            if i < min {
                min = i
            }
        }
        return min
    }
    
    func main() {
    	fmt.Println(^uint(0)) // bitwise NOT
    	fmt.Println(^uint(0) >> 1)
    	fmt.Println(Min(2, 3, 4, 5))
    }

    18446744073709551615
    9223372036854775807
    2


<a id="org545ef71"></a>

## Append

    // where T is a placeholder for any given type.
    // You can't actually write a function in Go where the type T is determined by
    // the caller. That's why append is built in: it needs support from the compiler.
    func append(slice []T, elements ...T) []T

    import "fmt"
    
    func main() {
    	x := []int{1,2,3}
    	x = append(x, 4, 5, 6)
    	fmt.Println(x)
    
    	// Append a slice to a slice
    	y := []int{7,8,9}
    	x = append(x, y...)
    	fmt.Println(x)
    }

    [1 2 3 4 5 6]
    [1 2 3 4 5 6 7 8 9]


<a id="org5fe8a97"></a>

# Initialization


<a id="orgcf94293"></a>

## Constants

    import "fmt"
    type ByteSize float64
    
    const (
    	_           = iota // ignore first 0 value by assigning to blank identifier
    	KB ByteSize = 1 << (10 * iota)
    	MB
    	GB
    	TB
    	PB
    	EB
    	ZB
    	YB
    )
    
    func (b ByteSize) String() string {
        switch {
        case b >= YB:
            return fmt.Sprintf("%.2fYB", b/YB)
        case b >= ZB:
            return fmt.Sprintf("%.2fZB", b/ZB)
        case b >= EB:
            return fmt.Sprintf("%.2fEB", b/EB)
        case b >= PB:
            return fmt.Sprintf("%.2fPB", b/PB)
        case b >= TB:
            return fmt.Sprintf("%.2fTB", b/TB)
        case b >= GB:
            return fmt.Sprintf("%.2fGB", b/GB)
        case b >= MB:
            return fmt.Sprintf("%.2fMB", b/MB)
        case b >= KB:
            return fmt.Sprintf("%.2fKB", b/KB)
        }
        return fmt.Sprintf("%.2fB", b)
    }
    
    func main() {
    	fmt.Println(KB, MB, GB, YB)
    }

    1.00KB 1.00MB 1.00GB 1.00YB


<a id="org4393214"></a>

## Variables

    import (
    	"fmt"
    	"os"
    )
    
    var (
    	home   = os.Getenv("HOME")
    	user   = os.Getenv("USER")
    	gopath = os.Getenv("GOPATH")
    )
    
    func main() {
    	fmt.Println(home, user, gopath)
    }

    /Users/eric eric /Users/eric/go


<a id="org8a071fe"></a>

## The init Function

-   each source file can define its own \*niladic\*(no parameters) `init` function
    to set up whatever state is required.
-   and Actually each file can have **multiple init functions**.
-   `init` is called after all the variable declarations in the package have
    evaluated their initializers, and those are evaluated only after all the
    imported packages have been initialized.
-   a common use of `init` functions is to verify or repair correctness of the
    program state before real execution begins.

    func init() {
        if user == "" {
            log.Fatal("$USER not set")
        }
        if home == "" {
            home = "/home/" + user
        }
        if gopath == "" {
            gopath = home + "/go"
        }
        // gopath may be overridden by --gopath flag on command line.
        flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
    }


<a id="org7bfbc72"></a>

# Methods


<a id="orgb6650c3"></a>

## TODO Pointers vs. Values

    import "fmt"
    
    type ByteSlice []byte
    
    func (p *ByteSlice) Append(data []byte) {
    	slice := *p
    	// Body as above, without the return.
    	*p = slice
    }
    
    func (p *ByteSlice) Write(data []byte) (n int, err error) {
    	slice := *p
    	// Again as above.
    	*p = slice
    	return len(data), nil
    }
    
    func main() {
    	var b ByteSlice
    	fmt.Fprintf(&b, "This hour has %d days\n", 7)
    	fmt.Println(b)
    	b.Write([]byte{7})
    	fmt.Println(b)
    }

    []
    []


<a id="orga5fe412"></a>

# Interfaces and other types


<a id="orgba96913"></a>

## Interfaces

    import (
    	"fmt"
    	"sort"
    )
    type Sequence []int
    
    // Methods required by sort.Interface.
    func (s Sequence) Len() int {
        return len(s)
    }
    func (s Sequence) Less(i, j int) bool {
        return s[i] < s[j]
    }
    func (s Sequence) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
    }
    
    // Copy returns a copy of the Sequence.
    func (s Sequence) Copy() Sequence {
        copy := make(Sequence, 0, len(s))
        return append(copy, s...)
    
    }
    
    // Method for printing - sorts the elements before printing.
    func (s Sequence) String() string {
        s = s.Copy() // Make a copy; don't overwrite argument.
        sort.Sort(s)
        str := "[{"
        for i, elem := range s { // Loop is O(N²); will fix that in next example.
            if i > 0 {
                str += " "
            }
            str += fmt.Sprint(elem)
        }
        return str + "}]"
    }
    
    func main() {
    	var s Sequence = []int{4, 3, 2, 1, 0, -1}
    	var ss = []int(s)
    	fmt.Println(ss, s)
    }

    [4 3 2 1 0 -1] [{-1 0 1 2 3 4}]


<a id="orgfc5a6ce"></a>

## Conversions

We can share the effort (and also speed it up) if we convert the Sequence to a
plain []int before calling Sprint.

    import (
    	"fmt"
    	"sort"
    )
    
    type Sequence []int
    
    func (s Sequence) Copy() Sequence {
        copy := make(Sequence, 0, len(s))
        return append(copy, s...)
    }
    
    func (s Sequence) String() string {
        s = s.Copy()
        sort.IntSlice(s).Sort()
        return fmt.Sprint([]int(s))
    }
    
    func main() {
    	var s Sequence = []int{3, 1, 2 }
    	fmt.Println(s)
    }

    [1 2 3]


<a id="orgd5a9920"></a>

## Interface conversions and type assertions

    type Stringer interface {
        String() string
    }
    
    var value interface{} // Value provided by caller.
    switch str := value.(type) {
    case string:
        return str
    case Stringer:
        return str.String()
    }

-   To extract the string we know is in the value, we could write:

    str := value.(string)
    
    str, ok := value.(string)
    if ok {
        fmt.Printf("string value is: %q\n", str)
    } else {
        fmt.Printf("value is not a string\n")
    }
    
    if str, ok := value.(string); ok {
        return str
    } else if str, ok := value.(Stringer); ok {
        return str.String()
    }


<a id="org93d2316"></a>

## Generality

-   The `crypto/cipher` interfaces look like this:

    type Block interface {
        BlockSize() int
        Encrypt(dst, src []byte)
        Decrypt(dst, src []byte)
    }
    
    type Stream interface {
        XORKeyStream(dst, src []byte)
    }

-   turns a block cipher into a streaming cipher; notice that the block cipher&rsquo;s details are abstracted away:

    // NewCTR returns a Stream that encrypts/decrypts using the given Block in
    // counter mode. The length of iv must be the same as the Block's block size.
    func NewCTR(block Block, iv []byte) Stream

NewCTR applies not just to one specific encryption algorithm and data source but
to any implementation of the Block interface and any Stream.


<a id="orgf7b30b8"></a>

## TODO Interfaces and methods

-   Any object that implements Handler can serve HTTP requests.

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }

-   Here&rsquo;s a trivial but complete implementation of a handler to count the number
    of times the page is visited.

    // Simple counter server.
    import (
    	"fmt"
    	"net/http"
    )
    type Counter struct {
        n int
    }
    
    func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
        ctr.n++
        fmt.Fprintf(w, "counter = %d\n", ctr.n)
    }
    
    func main() {
        ctr := new(Counter)
        http.Handle("/counter", ctr)
    }

-   But why make Counter a struct? An integer is all that&rsquo;s needed.
    (The receiver needs to be a pointer so the increment is visible to the caller.)


<a id="org5735754"></a>

# The blank identifier

-   The blank identifier in multiple assignment

    if _, err := os.Stat(path); os.IsNotExist(err) {
    	fmt.Printf("%s does not exist\n", path)
    }
    
    // Ignore the err
    // Bad! This code will crash if path does not exist.
    fi, _ := os.Stat(path)
    if fi.IsDir() {
        fmt.Printf("%s is a directory\n", path)
    }


<a id="orga7a5e4a"></a>

## Unused imports and variables

-   It is an error to import a package or to declare a variable without using it.

-   Assigning the unused variable fd to the blank identifier will silence the unused variable error

    package main
    
    import (
        "fmt"
        "io"
        "log"
        "os"
    )
    
    var _ = fmt.Printf // For debugging; delete when done.
    var _ io.Reader    // For debugging; delete when done.
    
    func main() {
        fd, err := os.Open("test.go")
        if err != nil {
            log.Fatal(err)
        }
        // TEDO: use fd.
        _ = fd
    }


<a id="org02bc32e"></a>

## Import for side effect

it is useful to import a package only for its side effects, without any explicit use.
For example, during its init function, the net/http/pprof package registers
HTTP handlers that provide debugging information.

    import _ "net/http/pprof"


<a id="org20f76a0"></a>

## Interface checks

-   The encoder checks this property at run time with a type assertion like

    m, ok := val.(json.Marshaler)

-   use the blank identifier to ignore the type-asserted value

    if _, ok := val.(json.Marshaler); ok {
        fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
    }

-   To guarantee that the implementation, a global declaration using the blank
    identifier can be used in the package

    // Should the json.Marshaler interface change, this package will no longer
    // compile and we will be on notice that it needs to be updated.
    var _ json.Marshaler = (*RawMessage)(nil)


<a id="org475d19f"></a>

# TODO Embedding

-   Embedded interface

    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    
    type Writer interface {
        Write(p []byte) (n int, err error)
    }
    
    // ReadWriter is the interface that combines the Reader and Writer interfaces.
    type ReadWriter interface {
        Reader
        Writer
    }

*Only interfaces can be embedded within interfaces.*

    // ReadWriter stores pointers to a Reader and a Writer.
    // It implements io.ReadWriter.
    type ReadWriter struct {
        *Reader  // *bufio.Reader
        *Writer  // *bufio.Writer
    }


<a id="org0ad0522"></a>

# TODO Concurrency


<a id="org969c2d3"></a>

## Share by communicating

-   **Do not communicate by sharing memory; instead, share memory by Dcommunicatingd.**


<a id="orga1e8a30"></a>

# Errors


<a id="orgd023c62"></a>

## By convention

-   Build-in Error interfce

    type error interface {
        Error() string
    }

-   A library writer is free to implement this interface

    // PathError records an error and the operation and
    // file path that caused it.
    type PathError struct {
        Op string    // "open", "unlink", etc.
        Path string  // The associated file.
        Err error    // Returned by the system call.
    }
    
    func (e *PathError) Error() string {
        return e.Op + " " + e.Path + ": " + e.Err.Error()
    }

-   Type assertion

    for try := 0; try < 2; try++ {
        file, err = os.Create(filename)
        if err == nil {
            return
        }
        if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
            deleteTempFiles()  // Recover some space.
            continue
        }
        return
    }


<a id="org99e21fc"></a>

## Panic

-   If the error is unrecoverable?(Something impossible has happened)

    // A toy implementation of cube root using Newton's method.
    func CubeRoot(x float64) float64 {
        z := x/3   // Arbitrary initial value
        for i := 0; i < 1e6; i++ {
            prevz := z
            z -= (z*z*z-x) / (3*z*z)
            if veryClose(z, prevz) {
                return z
            }
        }
        // A million iterations has not converged; something is wrong.
        panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
    }

-   During initialization: if the library truly cannot set itself up, it might be
    reasonable to panic, so to speak.

    var user = os.Getenv("USER")
    
    func init() {
        if user == "" {
            panic("no value for $USER")
        }
    }


<a id="orgf3c3c3c"></a>

## Recover

-   Shut down a failing goroutine inside a server without killing the other
    executing goroutines.

    func server(workChan <-chan *Work) {
        for work := range workChan {
            go safelyDo(work)
        }
    }
    
    func safelyDo(work *Work) {
        defer func() {
            if err := recover(); err != nil {
                log.Println("work failed:", err)
            }
        }()
        do(work)
    }

> In this example, if do(work) panics, the result will be logged and the goroutine
> will exit cleanly without disturbing the others. There&rsquo;s no need to do anything
> else in the deferred closure; calling recover handles the condition completely.

-   Reports parsing errors by calling panic with a local error type.

> The definition of Error, an error method, and the Compile function.

    // Error is the type of a parse error; it satisfies the error interface.
    type Error string
    func (e Error) Error() string {
        return string(e)
    }
    
    // error is a method of *Regexp that reports parsing errors by panicking with an Error.
    func (regexp *Regexp) error(err string) {
        panic(Error(err))
    }
    
    // Compile returns a parsed representation of the regular expression.
    func Compile(str string) (regexp *Regexp, err error) {
        regexp = new(Regexp)
        // doParse will panic if there is a parse error.
        defer func() {
            if e := recover(); e != nil {
                regexp = nil    // Clear return value.
                err = e.(Error) // Will re-panic if not a parse error.
            }
        }()
        return regexp.doParse(str), nil
    }

With error handling in place, the error method (because it&rsquo;s a method bound to a
type, it&rsquo;s fine, even natural, for it to have the same name as the builtin error
type) makes it easy to report parse errors without worrying about unwinding the
parse stack by hand:

    if pos == 0 {
        re.error("'*' illegal at start of expression")
    }

Useful though this pattern is, it should be used only within a package. Parse
turns its internal panic calls into error values; it does not expose panics to
its client. That is a good rule to follow.

By the way, this re-panic idiom changes the panic value if an actual error
occurs. However, both the original and new failures will be presented in the
crash report, so the root cause of the problem will still be visible. Thus this
simple re-panic approach is usually sufficient—it&rsquo;s a crash after all—but if you
want to display only the original value, you can write a little more code to
filter unexpected problems and re-panic with the original error.


<a id="orgdefac4e"></a>

# A Web Server

-> src @[awebserver/server.go](src/awebserver/server.go)


<a id="org575871f"></a>

# Go Docs


<a id="orgb0fa0a9"></a>

# References     :W:

-   Golang Doc: [Effective Go](https://golang.org/doc/effective_go.html#introduction)
-   [Golang Src](https://golang.org/src/)
-   [Go Slices: usage and internals](https://blog.golang.org/slices-intro)
-   [Code boilerplate: Is it always bad? | by Andrey Petrov | Medium](https://medium.com/@shazow/code-boilerplate-is-it-always-bad-934827efcfc7)
-   [godoc: Tips & Tricks. Go has a great emphasis on simple… | by Elliot Chance |&#x2026;](https://medium.com/@elliotchance/godoc-tips-tricks-cda6571549b)
