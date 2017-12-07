package day5

import (
	"bufio"
	"fmt"
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

func readInstruction(line string) (instruction int, ok bool) {
	s64, err := strconv.ParseInt(line, 10, 32)
	ok = err == nil
	instruction = int(s64)
	return
}

func readInstructions(filename string) (program []int) {
	program = []int{}

	computator := func(line string) {
		instruction, ok := readInstruction(line)
		if ok {
			program = append(program, instruction)
		}
	}

	iterateOverLinesInTextFile(filename, computator)
	return
}

func executeProgram(program []int) (steps int) {
	steps = 0

	PC := 0
	for PC < len(program) {
		OFFSET := program[PC]
		if OFFSET < 3 {
			program[PC]++
		} else {
			program[PC]--
		}
		PC = PC + OFFSET
		steps++
	}

	return
}

// Run1 is the primary solution
func Run1() {
	var program = readInstructions("day5/input.text")
	var steps = executeProgram(program)
	fmt.Printf("Day 5.1: Number of steps: %v \n", steps)
}

// Run2 is the secondary solution
func Run2() {
	var program = readInstructions("day5/input.text")
	var steps = executeProgram(program)
	fmt.Printf("Day 5.2: Number of steps: %v \n", steps)
}
