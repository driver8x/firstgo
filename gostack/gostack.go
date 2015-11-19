/*
Package gostack implements a simple stack of integers implemented using a linked list and a queue of integers using the stack
*/
package gostack

type node struct {
	Next *node
	Val  int
}

type Stack struct {
	head *node
}

type EmptyStackError struct{}

func (e *EmptyStackError) Error() string {
	return "Attempted to pop from an empty stack."
}

// Push adds in integer to the top of the stack
func (s *Stack) Push(val int) {
	s.head = &node{s.head, val}
}

// Pop removes and returns the integer at the top of the stack
func (s *Stack) Pop() (val int, err error) {
	if s.head == nil {
		return 0, &EmptyStackError{}
	}
	val = s.head.Val
	s.head = s.head.Next
	return val, nil
}

// NewStack creates and returns a new Stack and optionally initializes it with a given list of integers
func NewStack(input ...int) *Stack {
	s := &Stack{nil}
	for _, v := range input {
		s.Push(v)
	}
	return s
}

// Queue implements a queue of integers using two Stacks.
type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

type EmptyQueueError struct{}

func (e *EmptyQueueError) Error() string {
	return "Attempted to Dequeue from an empty queue"
}

// Enqueue adds an int to the back of the queue
func (q *Queue) Enqueue(val int) {
	if q.stack1 == nil {
		q.stack1 = NewStack(val)
	} else {
		q.stack1.Push(val)
	}
}

// Dequeue removes and returns the int at the front of the queue
func (q *Queue) Dequeue() (val int, err error) {

	if q.stack1 == nil {
		return 0, &EmptyQueueError{}
	}
	if q.stack2 != nil {
		if temp, err := q.stack2.Pop(); err == nil {
			return temp, nil
		}
	}

	temp, err := q.stack1.Pop()
	if err != nil {
		return 0, &EmptyQueueError{}
	}
	q.stack2 = new(Stack)
	for {
		val = temp
		temp, err = q.stack1.Pop()
		if err == nil {
			q.stack2.Push(val)
		} else {
			return val, nil
		}
	}
}

// NewQueue creates and returns a new Queue and optionally initializes it with a given list of integers
func NewQueue(input ...int) (q *Queue) {
	q = &Queue{}
	for _, v := range input {
		q.Enqueue(v)
	}
	return q
}
