// searchtree is a trie implementation that allows that addition of strings associated with slices of integers and
// fast prefix searching of said strings
package searchtree

import (
	"fmt"
	"unicode/utf8"
)

type node struct {
	name     string         // string to search on
	val      []int          // data associated with the string
	children map[rune]*node // map of any child nodes keyed on their next rune after the current name
}

type Tree struct {
	head *node
}

// Add adds a given string to the trie and associates a slice of ints data with it
func (t *Tree) Add(name string, data []int) {
	if t.head == nil {
		t.init()
	}
	cur := t.head

	i := -1
	for idx, val := range name {
		i = idx
		if _, ok := cur.children[val]; ok {
			cur = cur.children[val]
		} else {
			break
		}
	}

	for j, val := range name {
		rl := utf8.RuneLen(val)
		if j < i {
			continue
		} else if j+rl <= len(name) {
			if _, ok := cur.children[val]; !ok && cur.name != name {
				cur.children[val] = &node{name[:j+rl], nil, make(map[rune]*node)}
				cur = cur.children[val]
			}
		}
	}
	cur.val = data
}

// Find searches the trie for any entries with the given prefix and returns a map of strings to slices of ints
func (t *Tree) Find(name string) (results map[string][]int, err error) {
	results = make(map[string][]int)
	if t.head == nil {
		t.init()
	}
	cur := t.head
	for i, r := range name {
		if _, ok := cur.children[r]; !ok {
			return nil, &PrefixNotFoundError{name, i}
		}
		cur = cur.children[r]
	}

	datanodes := traverse(cur)
	for _, n := range datanodes {
		results[n.name] = n.val
	}

	return results, nil
}

// Remove searches the trie for any entries with the given prefix and removes them from the trie, returning the number of elements removed
func (t *Tree) Remove(name string) (number int, err error) {
	if t.head == nil {
		t.init()
	}
	cur := t.head
	for i, r := range name {
		if _, ok := cur.children[r]; !ok {
			return 0, &PrefixNotFoundError{name, i}
		}
		cur = cur.children[r]
	}

	datanodes := traverse(cur)
	number = len(datanodes)
	cur.val = nil
	cur.children = make(map[rune]*node)

	return number, nil
}

// traverse recursively searches a node and returns a slice containing pointers to all nodes with a non-nil value
func traverse(n *node) (coll []*node) {
	if n.val != nil {
		coll = []*node{n}
	}
	for _, v := range n.children {
		vals := traverse(v)
		coll = append(vals, coll...)
	}
	return coll
}

// init creates a dummy head node for the trie
func (t *Tree) init() {
	t.head = &node{"", nil, make(map[rune]*node)}
}

// PrefixNotFoundError is returned when no strings with a given prefix could be found in the try
type PrefixNotFoundError struct {
	name  string
	index int
}

func (p *PrefixNotFoundError) Error() string {
	return fmt.Sprintf("\"%v\" not found. Stopped at index %v (%v).", p.name, p.index, p.name[:p.index+1])
}
