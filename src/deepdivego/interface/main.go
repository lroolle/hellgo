package main

import (
	"bufio"
	"fmt"
	"interface/mock"
	"interface/queue"
	"interface/real"
	"io"
	"os"
	"strings"
	"time"
)

const url = "http://golang.org"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://lroolle.com")
}

// func post(poster Poster) {
// 	poster.Post("https://golang.org", map[string]string{"contents": "golangtour"})
// }

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "golangtour"})
	return s.Get(url)
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Inspect Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("Inspect UserAgent: ", v.UserAgent)
	}
}

func printFile(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var r Retriever
	fmt.Println("--- Mock Inspect ---")
	mockr := &mock.Retriever{"This is mock golang.org"}
	fmt.Printf("Type mockr: %T %v\n", mockr, mockr)
	inspect(mockr)

	fmt.Println("--- Real Inspect ---")
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("Type real: %T %v\n", r, r)
	inspect(r)

	if _, ok := r.(*mock.Retriever); ok {
		fmt.Println("Type assertion: this is mock")
	} else {
		fmt.Println("Type assertion: this is real")
	}

	download(r)

	fmt.Println("--- Queue interface{} ---")
	q := queue.Queue{}
	fmt.Println("q is Empty: ", q.Empty())
	q.Push(1)
	q.Push(2)
	fmt.Println("q is Empty: ", q.Empty(), q)
	q.Push("a")
	q.Push("bc")
	fmt.Println("q is Empty: ", q.Empty(), q)
	fmt.Println("Popped: ", q.Pop())
	fmt.Println("Popped: ", q.Pop())
	fmt.Println("Popped: ", q.Pop())
	fmt.Println("q is Empty: ", q.Empty(), q)

	fmt.Println("--- Session Use pointer to modify mock.Retriever contents ---")
	retriever := &mock.Retriever{}
	fmt.Println(session(retriever))

	// Standard lib in io
	// var w io.WriteCloser
	fmt.Println("--- Reader & Writer  ---")
	s := `"abc"
"def"`
	printFile(strings.NewReader(s))
	file, err := os.Open("go.mod")
	if err != nil {
		panic(err)
	}
	printFile(bufio.NewReader(file))
}
