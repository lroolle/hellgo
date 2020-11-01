package queue

// An FIFO Queue
type Queue []interface{}

// Pushes the element into the queue
//   e.g: q.Push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pops element from the queue tail
func (q *Queue) Pop() interface{} {
	head := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return head
}

// If queue is empty
func (q *Queue) Empty() bool {
	return len(*q) == 0
}
