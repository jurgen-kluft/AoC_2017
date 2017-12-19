package day23

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"strings"
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
func Run1(day int) {
	var s = read(fmt.Sprintf("Day%d/input.text", day))
	var result = s.solve()
	fmt.Printf("Day%d.1: . : %v \n", day, result)
}

// Run2 is the secondary solution
func Run2(day int) {
	var s = read(fmt.Sprintf("Day%d/input.text", day))
	var result = s.solve2()
	fmt.Printf("Day%d.2: . : %v \n", day, result)
}
