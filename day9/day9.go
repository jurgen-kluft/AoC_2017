package day9

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

func decode(line string) (i int, ok bool) {

	return 0, true
}

func read(filename string) (is []int) {

	reader := func(line string) {
		i, ok := decode(line)
		if ok {
			is = append(is, i)
		}
	}
	iterateOverLinesInTextFile(filename, reader)
	return is
}

// Run1 is the primary solution
func Run1() {
	var is = read("day9/input.text")
	fmt.Printf("Day 9.1: bleh: %v \n", is)
}

// Run2 is the secondary solution
func Run2() {
	var is = read("day9/input.text")
	fmt.Printf("Day 9.2: bleh %v \n", is)
}
