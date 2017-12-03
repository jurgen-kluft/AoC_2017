package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func iterateOverLinesInTextFile(filename string, action func(string)) {
	// Open the file.
	f, _ := os.Open(filename)
	defer f.Close()

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)

	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		action(line)
	}
}

func readNumber(line string, cursor int) (newcursor int, number int, ok bool) {
	start := cursor
	for start < len(line) && (line[start] == ' ' || line[start] == '\t') {
		start++
	}
	end := start
	for end < len(line) && (line[end] != ' ' && line[end] != '\t') {
		end++
	}

	if start == end {
		number = 0
		newcursor = len(line)
		ok = false
		return
	}

	numberStr := string(line[start:end])
	//fmt.Printf("%v(%v,%v) - ", numberStr, start, end)

	s64, err := strconv.ParseInt(numberStr, 10, 32)
	ok = err == nil
	number = int(s64)
	newcursor = end + 1
	return
}

func computeChecksum(filename string) (sum int, ok bool) {
	sum = 0

	computator := func(line string) {
		cursor := 0
		min := math.MaxInt32
		max := 0
		//fmt.Printf(" - ")
		for cursor < len(line) {
			nextCursor, number, ok := readNumber(line, cursor)
			if ok {
				if number < min {
					min = number
				}
				if number > max {
					max = number
				}
			}
			cursor = nextCursor
		}
		//fmt.Printf("\n")
		sum += max - min
	}

	iterateOverLinesInTextFile(filename, computator)
	ok = true
	return
}

func computeEvenDivisorSum(filename string) (sum int, ok bool) {
	sum = 0

	computator := func(line string) {
		cursor := 0
		numbers := []int{}
		//fmt.Printf(" - ")
		for cursor < len(line) {
			nextCursor, number, ok := readNumber(line, cursor)
			if ok {
				numbers = append(numbers, number)
			}
			cursor = nextCursor
		}
		n := 0
		for i1, n1 := range numbers {
			if (i1 + 1) < len(numbers) {
				for _, n2 := range numbers[i1+1:] {
					if n1 < n2 {
						if (n2/n1)*n1 == n2 {
							n = n2 / n1
						}
					} else if n1 > n2 {
						if (n1/n2)*n2 == n1 {
							n = n1 / n2
						}
					}
				}
			}
		}
		sum += n
	}

	iterateOverLinesInTextFile(filename, computator)
	ok = true
	return
}

// Run1 is the primary solution of
func Run1() {
	var sum, ok = computeChecksum("day2/input.text")
	if ok {
		fmt.Printf("Day 2.1: Sum is %v \n", sum)
	} else {
		fmt.Printf("Could not process the input")
	}
}

// Run2 is the secondary solution
func Run2() {
	var sum, ok = computeEvenDivisorSum("day2/input.text")
	if ok {
		fmt.Printf("Day 2.2: Checksum is %v \n", sum)
	} else {
		fmt.Printf("Could not process the input")
	}
}
