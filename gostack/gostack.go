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
	Tail *node
}

// Push adds in integer to the top of the stack
func (s *Stack) Push(i int) {
	s.Head = &node{s.Head, i}
}

// Pop removes and returns the integer at the top of the stack
func (s *Stack) Pop() int {
	temp := s.Head.Val
	s.Head = s.Head.Next
	return temp
}

// NewStack creates and returns a new Stack and optionally initializes it with a given list of integers
func NewStack(input ...int) *Stack {
	s := &Stack{nil, nil}
	for _, v := range input {
		s.Push(v)
	}
	return s
}

type Queue Stack

// Enqueue adds an int to the back of the queue
func (q *Queue) Enqueue(i int) {
	newnode := &node{nil, i}
	if q.Tail != nil {
		q.Tail.Next = newnode
		q.Tail = q.Tail.Next
	} else {
		q.Head = newnode
		q.Tail = newnode
	}
}

// Dequeue removes and returns the int at the front of the queue
func (q *Queue) Dequeue() int {
	var temp int
	if q.Head != nil {
		temp = q.Head.Val
		q.Head = q.Head.Next
		if q.Head == nil {
			q.Tail = nil
		}
	}
	return temp
}

// NewQueue creates and returns a new Queue and optionally initializes it with a given list of integers
func NewQueue(input ...int) *Queue {
	q := &Queue{nil, nil}
	for _, v := range input {
		q.Enqueue(v)
	}
	return q
}
