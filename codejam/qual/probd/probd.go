package main
import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"bytes"
	"fmt"
)

var tests [][3]int

func main() {
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\D-test.in")
	infile, _ := os.Open("C:\\app\\code\\CodeJam16\\D-small.in")
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\D-big.in")
	defer infile.Close()
	outfile, err := os.Create("C:\\app\\code\\CodeJam16\\D-small.out")

	if err != nil {
		fmt.Println(err)
	}
	defer outfile.Close()

	scanner := (*Scanner)(bufio.NewScanner(infile))

	numTests, _ := strconv.Atoi(scanner.ReadLine())

	tests = make([][3]int, numTests)
	for i := 0; i < numTests; i++ {
		for j, val := range strings.Split(scanner.ReadLine(), " ") {
			tests[i][j], _ = strconv.Atoi(val)
		}
	}

	answers := make([]string, numTests)
	for i:= 0; i < numTests; i++ {
		var buf bytes.Buffer
		for j := 1; j <= tests[i][2]; j++ {
			buf.WriteString(strconv.Itoa(j))
			buf.WriteString(" ")
		}
		answers[i] = strings.TrimSpace(buf.String())
	}

	// check each prime factor of C and K?

	w := bufio.NewWriter(outfile)
	for i := 0; i < numTests; i++ {
		//fmt.Printf("Case #%d: %v\n", i + 1, answers[i])
		fmt.Fprintf(w, "Case #%d: %v\n", i + 1, answers[i])
	}
	w.Flush()
}


type Scanner bufio.Scanner

func (s *Scanner) ReadLine() string {
	sc := (*bufio.Scanner)(s)
	sc.Scan()
	return sc.Text()
}