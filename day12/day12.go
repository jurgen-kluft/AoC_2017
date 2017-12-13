package day12

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

// Day indicates the specific Advent Of Code puzzle
const Day = "Day12"

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

// Program contains references to other programs
type Program struct {
	name     string
	parent   *Program
	programs map[string]*Program
}

// Solution is the hex-grid
type Solution struct {
	programs map[string]*Program
}

func (p *Program) link() {
	for _, r := range p.programs {
		if r.parent == nil {
			r.parent = p
			r.link()
		}
	}
}

func (s *Solution) solve() int {
	s.programs["0"].link()

	programs := 0
	for _, p := range s.programs {
		if p.parent != nil {
			programs++
		}
	}
	return programs
}

func (s *Solution) solve2() int {
	groups := 0
	for _, p := range s.programs {
		if p.parent == nil {
			groups++
			p.link()
		}
	}
	return groups
}

func (p *Program) addProgram(ref *Program) {
	p.programs[ref.name] = ref
}

func (s *Solution) regProgram(name string) *Program {
	program, exists := s.programs[name]
	if !exists {
		program = &Program{}
		program.name = name
		program.programs = map[string]*Program{}
		s.programs[name] = program
	}
	return program
}

func readProgram(line string, s *Solution) *Program {
	result := strings.Split(line, "<->")

	programName := strings.TrimSpace(result[0])
	program := s.regProgram(programName)

	programRefs := strings.Split(result[1], ",")
	for _, ref := range programRefs {
		ref = strings.TrimSpace(ref)
		refProgram := s.regProgram(ref)
		program.addProgram(refProgram)
	}
	return program
}

func read(filename string) *Solution {
	s := &Solution{}
	s.programs = map[string]*Program{}

	reader := func(line string) {
		readProgram(line, s)
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
