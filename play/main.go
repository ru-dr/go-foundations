package main

import "fmt"

type Queue []int

func (q *Queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		front := (*q)[0]
		*q = (*q)[1:]
		return front, true
	}
}

func (q *Queue) Peek() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		front := (*q)[0]
		return front, true
	}
}

func main() {
	q := Queue{}
	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Peek())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
