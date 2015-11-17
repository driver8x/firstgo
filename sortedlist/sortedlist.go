// Package sortedlist implements a sorted linked list of integers
package sortedlist

import (
	"fmt"
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
	if list.head == nil {
		list.head = &element{-1, nil} // dummy head node simplifies iteration
	}
	if list.head.Next == nil {
		list.head.Next = &element{value, nil}
	} else {
		for cur, next := list.head, list.head.Next; cur != nil; cur, next = next, next.Next {
			if next == nil || value <= next.Value {
				cur.Next = &element{value, next}
				break
			}
		}
	}
}

// PrintAll prints all of the elements
func (list *List) PrintAll() {
	if list.head != nil && list.head.Next != nil {
		for cur := list.head.Next; cur != nil; cur = cur.Next {
			fmt.Print(cur.Value, " ")
		}
	}
	fmt.Println()
}

// Remove removes all instances of value from the list and returns the number of elements that were removed
func (list *List) Remove(value int) (count int) {
	count = 0
	if list.head != nil && list.head.Next != nil {
		for prev, cur := list.head, list.head.Next; cur != nil; cur = cur.Next {
			if value == cur.Value {
				prev.Next = cur.Next
				count++
			} else {
				prev = cur
			}
		}
	}
	return count
}

// Min returns the min value. If the list is empty, it returns an error
func (list *List) Min() (value int, err error) {
	if list.head != nil && list.head.Next != nil {
		value = list.head.Next.Value
	} else {
		err = &EmptyListError{}
	}
	return value, err
}

// Max returns the max value. If the list is empty, it returns an error
func (list *List) Max() (value int, err error) {
	if list.head != nil && list.head.Next != nil {
		cur := list.head.Next
		for ; cur.Next != nil; cur = cur.Next {
		}
		value = cur.Value
	} else {
		err = &EmptyListError{}
	}
	return value, err
}

// EmptyListError is returned when an operation can't be performed due a list being empty
type EmptyListError struct{}

func (err *EmptyListError) Error() string {
	return "Error attempting to read from empty list"
}
