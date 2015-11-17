// For testing "github.com/driver8x/firstgo/gostack"
package main

import (
	"fmt"
	"github.com/driver8x/firstgo/gostack"
)

func main() {

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
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println("\n\n")

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
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(s, q)

	fmt.Println("\n\n")

	var s2 *gostack.Stack
	s2.Push(8)
	s2.Push(1337)
	fmt.Println(s2.Pop())
}
