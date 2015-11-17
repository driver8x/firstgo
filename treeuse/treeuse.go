package main

import (
	"fmt"
	"github.com/driver8x/firstgo/searchtree"
)

func main() {
	tree1 := searchtree.Tree{}
	tree1.Add("Abe", []int{69, 2015})
	tree1.Add("Abel", []int{666})
	tree1.Add("Abraham", []int{1776, 2015})

	res1, err1 := tree1.Find("A")
	res2, err2 := tree1.Find("Ab")
	res3, err3 := tree1.Find("Abe")
	res4, err4 := tree1.Find("Abel")
	res5, err5 := tree1.Find("Abeline")
	res6, err6 := tree1.Find("Abrah")
	res7, err7 := tree1.Find("Bro")
	res8, err8 := tree1.Find("be")
	res9, err9 := tree1.Find("")
	fmt.Println(res1, err1)
	fmt.Println(res2, err2)
	fmt.Println(res3, err3)
	fmt.Println(res4, err4)
	fmt.Println(res5, err5)
	fmt.Println(res6, err6)
	fmt.Println(res7, err7)
	fmt.Println(res8, err8)
	fmt.Println(res9, err9)
}
