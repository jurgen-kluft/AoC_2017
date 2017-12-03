package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func computeGotchaCode(filename string) (sum int, ok bool) {
	sum = 0
	computator := func(line string) {
		line += string(line[0])
		// fmt.Printf("Line: %s \n", line)
		var previousNumber = math.MaxInt32
		var currentNumber = math.MaxInt32
		for _, c := range line {
			if c >= '0' && c <= '9' {
				previousNumber = currentNumber
				currentNumber = (int(c) - '0')
				if previousNumber == currentNumber {
					sum += currentNumber
				}
			}
		}
	}

	iterateOverLinesInTextFile(filename, computator)
	ok = true
	return
}

func computeGotchaCode2(filename string) (sum int, ok bool) {
	sum = 0
	computator := func(line string) {
		line += string(line[0])
		// fmt.Printf("Line: %s \n", line)
		var previousNumber = math.MaxInt32
		var currentNumber = math.MaxInt32
		for _, c := range line {
			if c >= '0' && c <= '9' {
				previousNumber = currentNumber
				currentNumber = (int(c) - '0')
				if previousNumber == currentNumber {
					sum += currentNumber
				}
			}
		}
	}

	iterateOverLinesInTextFile(filename, computator)
	ok = true
	return
}

// Run1 is the primary solution of
func Run1() {
	var sum, ok = computeGotchaCode("day1/input.text")
	if ok {
		fmt.Printf("Day 1.1: Sum is %v \n", sum)
	} else {
		fmt.Printf("Could not process the input")
	}
}

// Run2 is the secondary solution
func Run2() {
	var sum, ok = computeGotchaCode2("day1/input.text")
	if ok {
		fmt.Printf("Day 1.2: Sum is %v \n", sum)
	} else {
		fmt.Printf("Could not process the input")
	}
}
