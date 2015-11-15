package sortedlist

import (
	"fmt"
)

type Element struct {
	Value int
	Next  *Element
}

type List struct {
	Head *Element
}

// Add adds value int to list in the correct sorted location
func (list *List) Add(value int) {
	if list.Head == nil {
		list.Head = &Element{-1, nil} // dummy head node simplifys iteration
	}
	if list.Head.Next == nil {
		list.Head.Next = &Element{value, nil}
	} else {
		for cur, next := list.Head, list.Head.Next; cur != nil; cur, next = next, next.Next {
			if next == nil || value <= next.Value {
				cur.Next = &Element{value, next}
				break
			}
		}
	}
}

// PrintAll prints all of the elements
func (list *List) PrintAll() {
	if list.Head != nil && list.Head.Next != nil {
		for cur := list.Head.Next; cur != nil; cur = cur.Next {
			fmt.Print(cur.Value, " ")
		}
	}
	fmt.Println()
}

// Remove removes all instances of value from the list and returns the number of elements that were removed
func (list *List) Remove(value int) (count int) {
	count = 0
	if list.Head != nil && list.Head.Next != nil {
		for prev, cur := list.Head, list.Head.Next; cur != nil; cur = cur.Next {
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
	if list.Head != nil && list.Head.Next != nil {
		value = list.Head.Next.Value
	} else {
		err = &EmptyListError{}
	}
	return value, err
}

// Max returns the max value. If the list is empty, it returns an error
func (list *List) Max() (value int, err error) {
	if list.Head != nil && list.Head.Next != nil {
		cur := list.Head.Next
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

func (err *EmptyListError) Error() (string) {
	return "Error attempting to read from empty list"
}
