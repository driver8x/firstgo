package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Guest struct {
	start      uint16
	totalDests uint8
	dests      []uint16
}

type tuple struct {
	idx  uint16
	wait uint16
}

var guests []Guest
var times [][]uint16

func main() {
	infile, _ := os.Open("C:/app/code/gocode/src/github.com/driver8x/firstgo/devdraft/3/t2a.txt")
	//infile, _ := os.Open("C:/app/code/gocode/src/github.com/driver8x/firstgo/devdraft/3/t2b.txt")
	//infile := os.Stdin
	scanner := (*Scanner)(bufio.NewScanner(infile))

	line := scanner.ReadLine()
	lineSplit := strings.Split(line, " ")
	Rides, _ := strconv.Atoi(lineSplit[0])
	Hours, _ := strconv.Atoi(lineSplit[1])

	times = make([][]uint16, Rides)
	for i := 0; i < Rides; i++ {
		line = scanner.ReadLine()
		times[i] = make([]uint16, Hours)
		for j, t := range strings.Split(line, " ") {
			temp, _ := strconv.Atoi(t)
			times[i][j] = uint16(temp)
		}
	}

	Queries, _ := strconv.Atoi(scanner.ReadLine())
	guests = make([]Guest, Queries)
	for i := 0; i < Queries; i++ {
		time, _ := strconv.Atoi(scanner.ReadLine())
		totalDests, _ := strconv.Atoi(scanner.ReadLine())
		guests[i] = Guest{uint16(time), uint8(totalDests), make([]uint16, totalDests)}
		for j, d := range strings.Split(scanner.ReadLine(), " ") {
			temp, _ := strconv.Atoi(d)
			guests[i].dests[j] = uint16(temp)
		}
	}

	fmt.Println(times)
	fmt.Println(guests)

	//	for _, g := range guests {
	//
	//	}

}

type Scanner bufio.Scanner

func (s *Scanner) ReadLine() string {
	sc := (*bufio.Scanner)(s)
	sc.Scan()
	return sc.Text()
}
