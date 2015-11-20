package main

import (
	"bufio"
	"fmt"
	"github.com/driver8x/firstgo/searchtree"
	"os"
)

func main() {
	var tree1 searchtree.Trie

	//fmt.Println(os.Getwd())
	infile, err := os.Open("./dict.txt")
	if err != nil {
		panic("Failed to open file")
	}

	i := 0
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {

		line := scanner.Text()
		if e := scanner.Err(); e != nil {
			panic("Scanner error")
		}
		tree1.Add(line, []int{8})
		i++
	}
	infile.Close()
	fmt.Println("Tree built\n")

	fmt.Println(tree1.Find("alba"))
	fmt.Println()
	fmt.Println(tree1.Find("ceru"))
	fmt.Println()
	fmt.Println(tree1.Find("adglkahl"))
	fmt.Println()
	fmt.Println(tree1.Find("zi"))
	fmt.Println()
}
