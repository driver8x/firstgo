// Call with "fizzbuzz <max> <fizz> <buzz>"
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	//n, a, b := 100, 3, 5
	n, err1 := strconv.Atoi(flag.Arg(0)) // max range
	a, err2 := strconv.Atoi(flag.Arg(1)) // Fizz
	b, err3 := strconv.Atoi(flag.Arg(2)) // Buzz

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Please supply numbers for N, fizz, and buzz.")
		fmt.Println("Ex: fizzbuzz 200 5 7")
		os.Exit(1)
	}
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
