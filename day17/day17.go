package day17

import (
	"fmt"
	//"strconv"
	//"strings"
)

// Solution is the hex-grid
type Solution struct {
	spin  int
	trail []int
}

func (s *Solution) solve() int {

	max := 2017
	num := 0
	pos := 0
	len := 1

	for num < max {
		num++

		// spin
		pos = (pos + s.spin) % len
		val := insertAfter(s.trail, len, pos, num)
		pos++
		len++
		if num == 2017 {
			return val
		}
	}

	return 0
}

func (s *Solution) solve2() int {
	max := 50000000
	num := 0
	pos := 0
	len := 1

	for num < max {
		num++
		pos = (pos + s.spin) % len
		insertAfter2(s.trail, len, pos, num)
		pos++
		len++
	}

	return s.trail[1]
}

func insertAfter(trail []int, len int, pos int, value int) int {
	i := len - 1
	for i > pos {
		trail[i+1] = trail[i]
		i--
	}
	trail[pos+1] = value
	return trail[pos+2]
}

func insertAfter2(trail []int, len int, pos int, value int) {
	//i := len - 1
	//for i > pos {
	//	trail[i+1] = trail[i]
	//	i--
	//}
	if pos == 0 {
		trail[pos+1] = value
	}
}

func read(filename string) *Solution {
	s := &Solution{}
	s.trail = make([]int, 50000000, 50000000)
	s.spin = 301

	return s
}

// Run1 is the primary solution
func Run1(day int) {
	var s = read(fmt.Sprintf("Day%d/input.text", day))
	var result = s.solve()
	fmt.Printf("Day%d.1: Number after 2017 : %v \n", day, result)
}

// Run2 is the secondary solution
func Run2(day int) {
	var s = read(fmt.Sprintf("Day%d/input.text", day))
	var result = s.solve2()
	fmt.Printf("Day%d.2: . : %v \n", day, result)
}
