package queue

import (
	"fmt"
)

func ExampleQueue_Pop() {
	var q = Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())

	// Output:
	// 3
	// 2
	// false
	// 1
	// true
}
