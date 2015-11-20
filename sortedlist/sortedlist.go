// Package sortedlist implements a sorted linked list of integers
package sortedlist

import (
	"bytes"
	"fmt"
	"strconv"
)

type element struct {
	Value int
	Next  *element
}

type List struct {
	head *element
}

// Add adds value int to list in the correct sorted location
func (list *List) Add(value int) {
	node := list.findLocation(value)
	node.Next = &element{value, node.Next}
}

// PrintAll prints all of the elements
func (list *List) PrintAll() {
	var buff bytes.Buffer
	if list.head != nil && list.head.Next != nil {
		for cur := list.head.Next; cur != nil; cur = cur.Next {
			temp := strconv.Itoa(cur.Value)
			buff.WriteString(temp)
			if cur.Next != nil {
				buff.WriteString(", ")
			}
		}
	}
	fmt.Println(buff.String())
}

// Remove removes all instances of value from the list and returns the number of elements that were removed
func (list *List) Remove(value int) (count int) {
	count = 0
	node := list.findLocation(value)
	for node.Next != nil && node.Next.Value == value {
		node.Next = node.Next.Next
		count++
	}
	return count
}

// Min returns the min value. If the list is empty, it returns an error
func (list *List) Min() (int, error) {
	if list.head != nil && list.head.Next != nil {
		return list.head.Next.Value, nil
	} else {
		return 0, &EmptyListError{}
	}
}

// Max returns the max value. If the list is empty, it returns an error
func (list *List) Max() (int, error) {
	if cur := list.head; cur != nil && cur.Next != nil {
		for ; cur.Next != nil; cur = cur.Next {
		}
		return cur.Value, nil
	} else {
		return 0, &EmptyListError{}
	}
}

func (list *List) findLocation(value int) *element {
	if list.head == nil {
		list.head = &element{-1, nil} // dummy head node simplifies iteration
	}
	for cur := list.head; cur != nil; cur = cur.Next {
		if cur.Next == nil || value <= cur.Next.Value {
			return cur
		}
	}
	return list.head
}

// EmptyListError is returned when an operation can't be performed due a list being empty
type EmptyListError struct{}

func (err *EmptyListError) Error() string {
	return "Error attempting to read from empty list"
}
