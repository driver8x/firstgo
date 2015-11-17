/*
Package gostack implements a simple stack of integers implemented using a linked list
*/
package gostack

type node struct {
	Next *node
	Val  int
}

type Stack struct {
	Head *node
}

// Push adds in integer to the top of the stack
func (s *Stack) Push(i int) {
	s.Head = &node{s.Head, i}
}

// Pop removes and returns the integer at the top of the stack
func (s *Stack) Pop() int {
	if s.Head == nil {
		panic("Tried to pop from empty stack")
	}
	temp := s.Head.Val
	s.Head = s.Head.Next
	return temp
}

// NewStack creates and returns a new Stack and optionally initializes it with a given list of integers
func NewStack(input ...int) *Stack {
	s := &Stack{nil}
	for _, v := range input {
		s.Push(v)
	}
	return s
}

// Enqueue adds an int to the back of the queue
func (s *Stack) Enqueue(i int) {
	newnode := &node{nil, i}
	if s.Head == nil {
		s.Head = newnode
	} else {
		tail := s.Head;
		for ; tail.Next != nil; tail = tail.Next {}
		tail.Next = newnode
	}
}

// Dequeue removes and returns the int at the front of the queue
func (s *Stack) Dequeue() int {
	if s.Head == nil {
		panic("Tried to Dequeue from empty queue");
	}
	val := s.Head.Val
	s.Head = s.Head.Next
	return val
}

// NewQueue creates and returns a new Queue and optionally initializes it with a given list of integers
func NewQueue(input ...int) *Stack {
	q := &Stack{nil}
	for _, v := range input {
		q.Enqueue(v)
	}
	return q
}
