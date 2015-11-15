// Call with "fizzbuzz <max> <fizz> <buzz>"
package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	//n, a, b := 100, 3, 5
	n, _ := strconv.Atoi(flag.Arg(0)) // max range
	a, _ := strconv.Atoi(flag.Arg(1)) // Fizz
	b, _ := strconv.Atoi(flag.Arg(2)) // Buzz

	astr := strconv.Itoa(a)
	bstr := strconv.Itoa(b)

	for i, printed := 1, false; i <= n; i++ {
		istr := strconv.Itoa(i)
		if i%a == 0 || strings.Contains(istr, astr) {
			fmt.Print("Fizz")
			printed = true
		}
		if i%b == 0 || strings.Contains(istr, bstr) {
			fmt.Print("Buzz")
			printed = true
		}
		if printed == false {
			fmt.Print(i)
		}
		fmt.Println()
		printed = false
	}
}
