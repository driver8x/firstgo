package main

import (
	"github.com/driver8x/firstgo/gostack"
	"fmt"
)

func main() {

//	s := gostack.NewStack(1)
//	gostack.Push(s, 2)
//	gostack.Push(s, 3)
//	gostack.Push(s, 4)
//	gostack.Push(s, 5)
//	fmt.Println(gostack.Pop(s))
//	gostack.Push(s, 6)
//	fmt.Println(gostack.Pop(s))
//	fmt.Println(gostack.Pop(s))
//	gostack.Push(s, 7)
//	gostack.Push(s, 8)
//	gostack.Push(s, 9)
//	fmt.Println(gostack.Pop(s))
//	fmt.Println(gostack.Pop(s))
//	fmt.Println(gostack.Pop(s))
//	fmt.Println(gostack.Pop(s))
//	fmt.Println(gostack.Pop(s))
//	fmt.Println(gostack.Pop(s))
//	fmt.Println(gostack.Pop(s))

	s := gostack.NewStack(1, 2, 3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Pop())
	s.Push(6)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(7)
	s.Push(8)
	s.Push(9)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop(), "\n\n")

	q := gostack.NewQueue(1, 2, 3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.Dequeue())
	q.Enqueue(6)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}


