package day15

import (
	"fmt"
	//"strconv"
	//"strings"
)

// Day define the specific puzzle from 'Advent of Code'
const Day = "Day15"

// Solution is the hex-grid
type Solution struct {
	generatorA uint64
	generatorB uint64

	generatorAMultiplier uint64
	generatorBMultiplier uint64

	generatorADivider uint64
	generatorBDivider uint64
}

func (s *Solution) solve() int {

	var iterations = 40 * 1000 * 1000

	a := s.generatorA
	b := s.generatorB

	c := 0

	i := 0
	for i < iterations {
		resultA := (a * s.generatorAMultiplier)
		resultB := (b * s.generatorBMultiplier)

		resultA = resultA % s.generatorADivider
		resultB = resultB % s.generatorBDivider

		if (resultA & 0xFFFF) == (resultB & 0xFFFF) {
			c++
		}

		a = resultA
		b = resultB

		i++
	}
	return c
}

func (s *Solution) solve2() int {
	var iterations = 5 * 1000 * 1000

	a := s.generatorA
	b := s.generatorB

	c := 0

	i := 0
	for i < iterations {
		resultA := (a * s.generatorAMultiplier)
		resultA = resultA % s.generatorADivider
		for resultA&0x3 != 0 {
			resultA = (resultA * s.generatorAMultiplier)
			resultA = resultA % s.generatorADivider
		}
		resultB := (b * s.generatorBMultiplier)
		resultB = resultB % s.generatorBDivider
		for resultB&0x7 != 0 {
			resultB = (resultB * s.generatorBMultiplier)
			resultB = resultB % s.generatorBDivider
		}

		if (resultA & 0xFFFF) == (resultB & 0xFFFF) {
			c++
		}

		a = resultA
		b = resultB

		i++
	}
	return c
}

func readLine(line string, s *Solution) {
	return
}

func read(filename string) *Solution {
	s := &Solution{}

	s.generatorA = 783
	s.generatorB = 325

	s.generatorAMultiplier = 16807
	s.generatorBMultiplier = 48271

	s.generatorADivider = 2147483647
	s.generatorBDivider = 2147483647

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
