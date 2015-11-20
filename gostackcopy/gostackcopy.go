/*
Package gostackcopy implements a simple stack of integers implemented using a linked list and a queue of integers using the stack
*/
package gostackcopy

type node struct {
	Next *node
	Val  int
}

type Stack struct {
	head *node
}

type EmptyStackError struct{}

func (e EmptyStackError) Error() string {
	return "Attempted to pop from an empty stack."
}

// Push adds in integer to the top of the stack
func (s Stack) Push(val int) Stack {
	s.head = &node{s.head, val}
	return s
}

// Pop removes and returns the integer at the top of the stack
func (s Stack) Pop() (sr Stack, val int, err error) {
	if s.head == nil {
		return s, 0, EmptyStackError{}
	}
	val = s.head.Val
	s.head = s.head.Next
	return s, val, nil
}

// NewStack creates and returns a new Stack and optionally initializes it with a given list of integers
func NewStack(input ...int) Stack {
	var s Stack
	for _, v := range input {
		s = s.Push(v)
	}
	return s
}

// Queue implements a queue of integers using two Stacks.
type Queue struct {
	stack1 Stack
	stack2 Stack
}

type EmptyQueueError struct{}

func (e EmptyQueueError) Error() string {
	return "Attempted to Dequeue from an empty queue"
}

// Enqueue adds an int to the back of the queue
func (q Queue) Enqueue(val int) Queue {
	q.stack1 = q.stack1.Push(val)
	return q
}

// Dequeue removes and returns the int at the front of the queue
func (q Queue) Dequeue() (qr Queue, val int, err error) {
	var temp int
	if q.stack2, temp, err = q.stack2.Pop(); err == nil {
		return q, temp, nil
	} else if q.stack1, temp, err = q.stack1.Pop(); err != nil {
		return q, 0, EmptyQueueError{}
	}
	for {
		val = temp
		if q.stack1, temp, err = q.stack1.Pop(); err == nil {
			q.stack2 = q.stack2.Push(val)
		} else {
			return q, val, nil
		}
	}
}

// NewQueue creates and returns a new Queue and optionally initializes it with a given list of integers
func NewQueue(input ...int) (q Queue) {
	for _, v := range input {
		q = q.Enqueue(v)
	}
	return q
}
