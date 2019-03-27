package queue

import (
	"fmt"
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	q := new(Queue)
	q.Push(1)
	q.Push("2")
	q.Push(3)
	pop := q.Pop()
	fmt.Printf("type: %T, value: %v\n", pop, pop)
	pop = q.Pop()
	fmt.Printf("type: %T, value: %v\n", pop, pop)
	fmt.Println(q.IsEmpty())
	pop = q.Pop()
	fmt.Printf("type: %T, value: %v\n", pop, pop)
	fmt.Println(q.IsEmpty())
}
