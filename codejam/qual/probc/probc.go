package main
import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"math/big"
	"fmt"
)

var tests []int

func main() {
	infile, _ := os.Open("C:\\app\\code\\CodeJam16\\C-test.in")
	//infile, _ := os.Open("C:\\app\\code\\CodeJam16\\C-small.in")
	//infile := os.Stdin
	defer infile.Close()

	scanner := (*Scanner)(bufio.NewScanner(infile))

	//numTests, _ := strconv.Atoi(scanner.ReadLine())
	scanner.ReadLine()
	toConvert := strings.Split(scanner.ReadLine(), " ")
	tempN, _ := strconv.Atoi(toConvert[0])
	tempJ, _ := strconv.Atoi(toConvert[1])

	N := big.NewInt(int64(tempN - 1))
	J := big.NewInt(int64(tempJ))

	answers := make([][]big.Int, tempJ)
	for i, _ := range answers {
		answers[i] = make([]big.Int, 9)
	}

//	number := big.NewInt(10).Exp(big.NewInt(10), N, big.NewInt(0))
//	number.Add(number, big.NewInt(1))

	number := ni(0).Exp(ni(2), N, ni(0))
	number.Sub(number, ni(1))

	foundcoins := 0

	// while found jamcoins < J
	for ; foundcoins < tempJ ; {
		number.Add(number, ni(2))
		str := number.Text(2)
		prime := false
		for i := 2; i < 11 && !prime ; i++ {
			num, _ := ni(0).SetString(str, i)

		}
	}

	// increment jamcoin

	// for each base

	// test primeness

	fmt.Println(number.Text(2), J)
}


type Scanner bufio.Scanner

func (s *Scanner) ReadLine() string {
	sc := (*bufio.Scanner)(s)
	sc.Scan()
	return sc.Text()
}

func ni(x int) *big.Int {
	return big.NewInt(int64(x))
}