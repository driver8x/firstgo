package searchtree

import(
	"strings"
)

type Node struct {
	Val      byte
	Children []*Node
}

type Tree struct {
	Head *Node
}

func (t *Tree) Add(name string) {
	if t.Head == nil {
		t.Head = &Node{' ', make([]*Node, 26)}
	}

	chars := []byte(strings.ToLower(name))
	curnode := t.Head
	for _, v := range chars {
		idx := int(v - 'a')
		if curnode.Children[idx] == nil {
			curnode.Children[idx] = &Node{v, make([]*Node, 26)}
		}
		curnode = curnode.Children[idx]
	}
}
