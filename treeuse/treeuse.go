package main

import (
	"fmt"
	"github.com/driver8x/firstgo/searchtree"
)

func main() {
	tree1 := searchtree.Trie{}
	tree1.Add("Abe", []int{69, 2015})
	tree1.Add("Abel", []int{666})
	tree1.Add("Abraham", []int{1776, 2015})
	tree1.Add("Pat", []int{8})
	tree1.Add("Patrick", []int{8})
	tree1.Add("driver8", []int{8})

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

	fmt.Println("\n")

	r1, e1 := tree1.Find("P")
	fmt.Println(r1, e1)
	fmt.Println(tree1.Remove("Patrick"))
	r2, e2 := tree1.Find("Pat")
	fmt.Println(r2, e2)
	fmt.Println(tree1.Remove("driver"))
	tree1.Add("Patrick", []int{88})
	r3, e3 := tree1.Find("Pat")
	fmt.Println(r3, e3)
	fmt.Println(tree1.Remove("Pat"))
	tree1.Add("Patri8", []int{8})
	r4, e4 := tree1.Find("P")
	fmt.Println(r4, e4)

	fmt.Println()

	tree1.Add("Pat", []int{8})
	tree1.Add("Patrick", []int{888})
	tree1.Add("Patri8", []int{88})
	tree1.Add("Patri8", []int{888})
	tree1.Add("Patri8", []int{8888})
	tree1.Add("Patri8", []int{88888})
	tree1.Add("Patri8s", []int{8})
	tree1.Add("Patri8s", []int{8})
	r5, e5 := tree1.Find("Pat")
	fmt.Println(r5, e5)
	fmt.Println(tree1.Remove("Patri8"))
	fmt.Println(tree1.Remove("Patri8"))
	tree1.Add("Patri8", []int{88888888})
	tree1.Add("Patrick", []int{88888888})
	r6, e6 := tree1.Find("Pat")
	fmt.Println(r6, e6)
	tree1.Add("Patri8", []int{888888888888888})
	tree1.Add("Patri8s", []int{8})
	tree1.Add("Patri8s", []int{888})
	tree1.Add("Patri8", []int{888888888888888})
	tree1.Add("Patri8s", []int{8})
	fmt.Println(tree1.Remove("Patrick"))
	fmt.Println(tree1.Remove("Patri8s"))
	r7, e7 := tree1.Find("Pat")
	fmt.Println(r7, e7)
}
