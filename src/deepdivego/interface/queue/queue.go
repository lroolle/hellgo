package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return head
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}
