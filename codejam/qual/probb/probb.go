package main
import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

var tests [][]string

func main() {
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\B-test.in")
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\B-small.in")
	infile, _ := os.Open("C:\\app\\code\\CodeJam16\\B-big.in")
	defer infile.Close()
	scanner := (*Scanner)(bufio.NewScanner(infile))

	numTests, _ := strconv.Atoi(scanner.ReadLine())

	tests = make([][]string, numTests)
	for i := 0; i < numTests; i++ {
		tests[i] = strings.Split(scanner.ReadLine(), "")
	}


	answers := make([]int, numTests)
	for testNum, str := range tests {

		enc := str[0]
		total := 0
		if enc == "-" {
			total = 1
		}
		for i := 1; i < len(str); i++ {
			if (str[i] == "+") {
				enc = "+"
				continue
			}
			if (str[i] == "-") {
				if (enc == "+") {
					total += 2
					enc = "-"
				}
			}
		}
		answers[testNum] = total
	}

	for i, ans := range answers {
		fmt.Printf("Case #%d: %v\n", i + 1, ans)
	}

}


type Scanner bufio.Scanner

func (s *Scanner) ReadLine() string {
	sc := (*bufio.Scanner)(s)
	sc.Scan()
	return sc.Text()
}