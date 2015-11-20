package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var travelTimes [][]int
var numAttractions int
var times []int

func main() {
	infile, _ := os.Open("C:/app/code/gocode/src/github.com/driver8x/firstgo/devdraft/3/t1a.txt")
	//infile, _ := os.Open("C:/app/code/gocode/src/github.com/driver8x/firstgo/devdraft/3/t1b.txt")
	//infile := os.Stdin
	scanner := (*Scanner)(bufio.NewScanner(infile))

	numAttractions, _ = strconv.Atoi(scanner.ReadLine())
	line := scanner.ReadLine()
	times = make([]int, numAttractions)
	for i, t := range strings.Split(line, " ") {
		times[i], _ = strconv.Atoi(t)
	}

	numQueries, _ := strconv.Atoi(scanner.ReadLine())
	queries := make([][]int, numQueries)
	for i := 0; i < numQueries; i++ {
		numDests, _ := strconv.Atoi(scanner.ReadLine())
		line := scanner.ReadLine()
		temp := make([]int, numDests)
		for j, d := range strings.Split(line, " ") {
			temp[j], _ = strconv.Atoi(d)
		}
		queries[i] = temp
	}

	travelTimes = make([][]int, numAttractions)
	totalTimes := make([]int, numQueries)
	for i, query := range queries {
		total := 0
		prev := -1
		for _, d := range query {
			if prev == -1 {
				prev = d
				continue
			}
			total += calcTime(prev, d)
			prev = d
		}
		totalTimes[i] = total
	}

	for _, total := range totalTimes {
		fmt.Println(total)
	}
}

func calcTime(a, b int) int {
	if a > b {
		a, b = b, a
	}
	if travelTimes[a] != nil {
		if travelTimes[a][b-a] > 0 {
			return travelTimes[a][b-a]
		}
	} else {
		travelTimes[a] = make([]int, numAttractions-a)
	}
	var left, right int
	for i := a; i != b; i++ {
		right += times[i]
	}
	for i := b; i != a; i = (i + 1) % numAttractions {
		left += times[i]
	}
	shortest := left
	if right < left {
		shortest = right
	}

	travelTimes[a][b-a] = shortest
	return shortest
}

type Scanner bufio.Scanner

func (s *Scanner) ReadLine() string {
	sc := (*bufio.Scanner)(s)
	sc.Scan()
	return sc.Text()
}
