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

type EmptyStackError struct {}

func (e *EmptyStackError) Error() string {
	return "Attempted to pop from an empty stack."
}

// Push adds in integer to the top of the stack
func (s *Stack) Push(val int) {
	if s == nil {
		panic("Attempted operation on uninitialized Stack variable")
	}
	s.head = &node{s.head, val}
}

// Pop removes and returns the integer at the top of the stack
func (s *Stack) Pop() (val int, err error) {
	if s == nil {
		panic("Attempted operation on uninitialized Stack variable")
	} else if s.head == nil {
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


// Queue implements a queue of integers using a stack. Dequeue is inefficient.
type Queue struct {
	stack *Stack
}

type EmptyQueueError struct {}

func (e *EmptyQueueError) Error() string {
	return "Attempted to Dequeue from an empty queue"
}

// Enqueue adds an int to the back of the queue
func (q *Queue) Enqueue(val int) {
	if q == nil {
		panic("Attempted operation on uninitialized Queue variable")
	} else if q.stack == nil {
		q.stack = NewStack(val)
	} else {
		q.stack.Push(val)
	}
}

// Dequeue removes and returns the int at the front of the queue
func (q *Queue) Dequeue() (val int, err error) {
	if q == nil {
		panic("Attempted operation on uninitialized Queue variable")
	} else if q.stack == nil {
		return 0, &EmptyQueueError{}
	}
	temp, err := q.stack.Pop()
	if err != nil {
		return 0, &EmptyQueueError{}
	}
	for {
		val = temp
		temp, err = q.stack.Pop()
		if err == nil {
			defer q.stack.Push(val)
		} else {
			break
		}
	}
	return val, nil
}

// NewQueue creates and returns a new Queue and optionally initializes it with a given list of integers
func NewQueue(input ...int) (q *Queue) {
	q = &Queue{}
	for _, v := range input {
		q.Enqueue(v)
	}
	return q
}


//// Enqueue adds an int to the back of the queue
//func (s *Stack) Enqueue(val int) {
//	newnode := &node{nil, val}
//	if s.Head == nil {
//		s.Head = newnode
//	} else {
//		tail := s.Head
//		for ; tail.Next != nil; tail = tail.Next {
//		}
//		tail.Next = newnode
//	}
//}
//
//// Dequeue removes and returns the int at the front of the queue
//func (s *Stack) Dequeue() (val int, err error) {
//	if s.Head == nil {
//		return 0, &EmptyStackError{}
//	}
//	val = s.Head.Val
//	s.Head = s.Head.Next
//	return val, nil
//}
//
//// NewQueue creates and returns a new Queue and optionally initializes it with a given list of integers
//func NewQueue(input ...int) *Stack {
//	q := &Stack{nil}
//	for _, v := range input {
//		q.Enqueue(v)
//	}
//	return q
//}