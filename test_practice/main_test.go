package test_practice

import (
	"fmt"
	"testing"
)

// 用两个stack，作一个队列
type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

func NewQueue(col int) *Queue {
	return &Queue{
		stack1: NewStack(col),
		stack2: NewStack(col),
	}
}
func (q Queue) Push(elem int) {
	for q.stack2.Len() >0 {
		q.stack1.Push(q.stack2.Pop())
	}
	q.stack1.Push(elem)
}

func (q Queue) Pop() int {
	for q.stack1.Len() >0 {
		q.stack2.Push(q.stack1.Pop())
	}
	rs:=q.stack2.Pop()
	return rs
}

type Stack struct {
	element []int
}

func NewStack(col int) *Stack {
	return &Stack{
		element: make([]int, 0, col),
	}
}

func (s Stack) Len() int{
	return len(s.element)
}

func (s *Stack) Push(elem int) {
	s.element = append(s.element, elem)
}

func (s *Stack) Pop() int {
	if len(s.element) == 0 {
		return 0
	}
	tmp := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return tmp
}

func TestStack(t *testing.T) {
	stack := NewStack(10)

	stack.Push(5)
	stack.Push(6)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}


func TestQueue(t *testing.T) {
	queue := NewQueue(10)

	queue.Push(5)
	queue.Push(6)
	queue.Push(7)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	queue.Push(8)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
