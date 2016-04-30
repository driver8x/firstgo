package main
import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"bytes"
)

var tests []string

func main() {
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\1A\\A-test.in")
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\1A\\A-small.in")
	infile, _ := os.Open("C:\\app\\code\\CodeJam16\\1A\\A-big.in")
	defer infile.Close()
	//outfile, err := os.Create("C:\\app\\code\\CodeJam16\\1A\\A-small.out")
	outfile, err := os.Create("C:\\app\\code\\CodeJam16\\1A\\A-big.out")

	if err != nil {
		fmt.Println(err)
	} else {
		defer outfile.Close()
	}

	scanner := &Scanner{bufio.NewScanner(infile)}
	scanner.Split(bufio.ScanLines)
	numTests, _ := strconv.Atoi(scanner.ReadLine())
	//	scanner.Split(bufio.ScanWords)
	//	numTests := scanner.NextInt()

	tests = make([]string, numTests)
	answers := make([]string, numTests)
	for i := 0; i < numTests; i++ {
		val := scanner.ReadLine()
		length := len(val)
		templist := &LL{}
		for j := 0; j < length; j++ {
			templist.Add(int(val[j]))
		}
		answers[i] = templist.String()
	}

	w := bufio.NewWriter(outfile)
	for i := 0; i < numTests; i++ {
		fmt.Printf("Case #%d: %v\n", i + 1, answers[i])
		fmt.Fprintf(w, "Case #%d: %v\n", i + 1, answers[i])
	}
	w.Flush()
}


type Scanner struct {
	*bufio.Scanner
}

func (s *Scanner) ReadLine() string {
	s.Scan()
	return s.Text()
}

func (s *Scanner) NextInt() int {
	s.Scan()
	x, err := strconv.Atoi(s.Text())
	for ; err != nil ; {
		s.Scan()
		x, err = strconv.Atoi(s.Text())
	}
	return x
}

type element struct {
	Value int
	Next  *element
}

type LL struct {
	head *element
	tail *element
}

func (list *LL) Add(x int) {
	if list.head == nil {
		list.head = &element{x, nil}
		list.tail = list.head
	} else {
		if x >= list.head.Value {
			temp := &element{x, list.head}
			list.head = temp
		} else {
			temp := &element{x, nil}
			list.tail.Next = temp
			list.tail = temp
		}
	}

}

func (list *LL) String() string {
	var buff bytes.Buffer
	if list.head != nil {
		for cur := list.head; cur != nil; cur = cur.Next {
			temp := string(cur.Value)
			buff.WriteString(temp)
		}
	}
	return buff.String()
}