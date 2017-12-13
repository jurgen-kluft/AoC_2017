package day11

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

const Day = "Day11"

const (
	dirUp    = 0x10
	dirDown  = 0x20
	dirRight = 0x01
	dirLeft  = 0x02

	hexDirN  = dirUp
	hexDirNE = dirUp | dirRight
	hexDirNW = dirUp | dirLeft
	hexDirSE = dirDown | dirRight
	hexDirSW = dirDown | dirLeft
	hexDirS  = dirDown
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
	trail []byte
	step  int
	x, y  int
}

func (s *Solution) move() bool {
	dir := s.trail[s.step]
	switch dir {
	case hexDirN:
		s.y++
		break
	case hexDirS:
		s.y--
		break
	case hexDirNE:
		s.x++
		break
	case hexDirNW:
		s.x--
		s.y++
		break
	case hexDirSE:
		s.x++
		s.y--
		break
	case hexDirSW:
		s.x--
		break
	}
	s.step++
	return s.step < len(s.trail)
}

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Solution) printTrail() {
	for _, dir := range s.trail {
		switch dir {
		case hexDirN:
			fmt.Print("n,")
			break
		case hexDirS:
			fmt.Print("s,")
			break
		case hexDirNE:
			fmt.Print("ne,")
			break
		case hexDirNW:
			fmt.Print("nw,")
			break
		case hexDirSE:
			fmt.Print("se,")
			break
		case hexDirSW:
			fmt.Print("sw,")
			break
		default:
			fmt.Printf("Error reading direction %d\n", dir)
		}
	}
}

func (s *Solution) solve() int {
	fmt.Printf("Trail has %d steps\n", len(s.trail))

	// Find out the shortest path to the location of the lost child process
	for s.move() {
	}
	fmt.Printf("Trail stops at (%d,%d) \n", s.x, s.y)
	z := -(s.x + s.y)

	// Find the largest axis value
	d := abs(z)
	d = max(abs(s.x), d)
	d = max(abs(s.y), d)
	return d
}

func (s *Solution) solve2() int {
	fmt.Printf("Trail has %d steps\n", len(s.trail))

	// Find out the shortest path to the location of the lost child process
	maxD := 0
	for s.move() {
		z := -(s.x + s.y)
		d := abs(z)
		d = max(abs(s.x), d)
		d = max(abs(s.y), d)
		if d > maxD {
			maxD = d
		}
	}
	fmt.Printf("Trail stops at (%d,%d) \n", s.x, s.y)

	// Find the largest axis value
	return maxD
}

func readLine(line string, s *Solution) {
	dirs := strings.Split(line, ",")
	for _, dir := range dirs {
		var step byte
		switch dir {
		case "n":
			step = hexDirN
			break
		case "s":
			step = hexDirS
			break
		case "ne":
			step = hexDirNE
			break
		case "nw":
			step = hexDirNW
			break
		case "se":
			step = hexDirSE
			break
		case "sw":
			step = hexDirSW
			break
		default:
			fmt.Printf("Error reading direction %s\n", dir)
		}
		s.trail = append(s.trail, step)
	}
	return
}

func read(filename string) *Solution {
	s := &Solution{}
	s.x = 0
	s.y = 0
	s.step = 0
	s.trail = []byte{}

	reader := func(line string) {
		readLine(line, s)
	}
	iterateOverLinesInTextFile(filename, reader)

	return s
}

// Run1 is the primary solution
func Run1() {
	var s = read(Day + "/input.text")
	//s.printTrail()
	var result = s.solve()
	fmt.Printf(Day+".1: . : %v \n", result)
}

// Run2 is the secondary solution
func Run2() {
	var s = read(Day + "/input.text")
	var result = s.solve2()
	fmt.Printf(Day+".2: . : %v \n", result)
}
