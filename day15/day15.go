package day15

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"strings"
)

// Day define the specific puzzle from 'Advent of Code'
const Day = "Day15"

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

// Solution is the hex-grid
type Solution struct {
}

func (s *Solution) solve() int {
	return 0
}

func (s *Solution) solve2() int {
	return 0
}

func readLine(line string, s *Solution) {
	return
}

func read(filename string) *Solution {
	s := &Solution{}

	reader := func(line string) {
		readLine(line, s)
	}
	iterateOverLinesInTextFile(filename, reader)

	return s
}

// Run1 is the primary solution
func Run1() {
	var s = read(Day + "/input.text")
	var result = s.solve()
	fmt.Printf(Day+".1: . : %v \n", result)
}

// Run2 is the secondary solution
func Run2() {
	var s = read(Day + "/input.text")
	var result = s.solve2()
	fmt.Printf(Day+".2: . : %v \n", result)
}
