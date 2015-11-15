package main

import(
	"fmt"
	"github.com/driver8x/firstgo/sortedlist"
)

func main() {
	list1 := &sortedlist.List{}
	list1.PrintAll()
	list1.Add(25)
	list1.PrintAll()
	list1.Remove(25)
	list1.PrintAll()
	list1.Add(26)
	list1.Add(35)
	list1.Add(27)
	list1.PrintAll()
	list1.Add(17)
	list1.Add(26)
	list1.Add(19)
	list1.Add(112)
	list1.PrintAll()
	list1.Add(26)
	fmt.Println(list1.Remove(26))
	fmt.Println(list1.Min())
	fmt.Println(list1.Max())
	list1.Remove(112)
	list1.Remove(17)
	fmt.Println(list1.Max())
	fmt.Println(list1.Min())
	list1.PrintAll()

	fmt.Println()
	list2 := &sortedlist.List{}
	list2.Add(3)
	list2.Add(5)
	list2.Add(4)
	list2.PrintAll()
	list2.Remove(3)
	fmt.Println(list2.Remove(7))
	fmt.Println(list2.Remove(4))
	fmt.Println(list2.Remove(5))
	fmt.Println(list2.Min())
	fmt.Println(list2.Max())
}