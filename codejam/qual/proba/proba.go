package main
import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

var tests []int

func main() {
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\A-test.in")
	infile, _ := os.Open("C:\\app\\code\\CodeJam16\\A-big.in")
	//infile := os.Stdin

	defer infile.Close()
	scanner := (*Scanner)(bufio.NewScanner(infile))

	numTests, _ := strconv.Atoi(scanner.ReadLine())
	tests := make([]int, numTests)
	answers := make([]string, numTests)

	// read all input
	for i := 0; i < numTests; i++ {
		tests[i], _ = strconv.Atoi(scanner.ReadLine())
	}

	// create slice with "0" .. "9"
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// for each test case, while digits remain in slice
	for i, testNum := range tests {
		if testNum == 0 {
			answers[i] = "INSOMNIA"
			continue
		}
		digs := make([]string, 10)
		copy(digs, digits)
		original := uint64(testNum)
		var num uint64
		for ; len(digs) > 0; {
			num = num + original
			numString := strconv.FormatUint(num, 10)
			//fmt.Println("Checking ", num, "\t Need ", digs)
			// check each remaining digit
			for j := 0; j < len(digs); {
				checkNum := digs[j]
				if strings.Contains(numString, checkNum) {
					// remove digit from slice if found
					for k := j; k + 1 < len(digs); k++ {
						digs[k] = digs[k+1]
					}
					digs = digs[:len(digs) - 1]
				} else {
					j++
				}

			}
		}
		answers[i] = strconv.FormatUint(num, 10)
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