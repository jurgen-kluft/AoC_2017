package day22

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

func readLine(line string, s *Solution) []byte {
	row := []byte{}

	for _, c := range line {
		switch {
		case c == '.':
			row = append(row, byte(0))
			break
		case c == '#':
			row = append(row, byte(1))
			break
		}
	}
	return row
}

func read(filename string) *Solution {
	s := &Solution{}
	s.grid = [][]byte{}
	reader := func(line string) {
		row := readLine(line, s)
		s.grid = append(s.grid, row)
	}
	iterateOverLinesInTextFile(filename, reader)
	return s
}

const (
	DIR_UP    = 0
	DIR_LEFT  = 1
	DIR_DOWN  = 2
	DIR_RIGHT = 3
)

func TurnLeft(dir int) int {
	dir = (dir + 1) & 3
	return dir
}
func TurnRight(dir int) int {
	dir = (dir + 4 - 1) & 3
	return dir
}
func Move(dir int, x int, y int) (nx int, ny int) {
	switch dir {
	case DIR_LEFT:
		nx = x - 1
		ny = y
		break
	case DIR_RIGHT:
		nx = x + 1
		ny = y
		break
	case DIR_UP:
		nx = x
		ny = y - 1
		break
	case DIR_DOWN:
		nx = x
		ny = y + 1
		break
	}
	return nx, ny
}

func dirStr(dir int) string {
	switch dir {
	case DIR_LEFT:
		return "Left"
	case DIR_RIGHT:
		return "Right"
	case DIR_UP:
		return "Up"
	case DIR_DOWN:
		return "Down"
	}
	return "?"
}

// Solution is the hex-grid
type Solution struct {
	grid  [][]byte
	dgrid map[uint64]byte
}

func (s *Solution) getCell(x int, y int) byte {
	x += 100000
	y += 100000
	key := uint64(uint64(x)<<32 | uint64(y))
	c, ok := s.dgrid[key]
	if ok {
		return c
	}
	s.dgrid[key] = 0
	return 0
}

func (s *Solution) setCell(x int, y int, c byte) {
	x += 100000
	y += 100000
	key := uint64(uint64(x)<<32 | uint64(y))
	s.dgrid[key] = c
}

func (s *Solution) solve() int {
	// Fill the dynamic grid
	s.dgrid = map[uint64]byte{}
	for gy, gl := range s.grid {
		for gx, gc := range gl {
			s.setCell(gx, gy, gc)
		}
	}

	fmt.Printf("Grid size (%d, %d)\n", len(s.grid[0]), len(s.grid))

	// Put the virus in the middle of the grid
	x := len(s.grid[0]) / 2
	y := len(s.grid) / 2
	dir := DIR_UP

	fmt.Printf("Virus starts at (%d, %d)\n", x, y)

	infections := 0
	iterations := 0
	maxIterations := 10000
	for iterations < maxIterations {
		//fmt.Printf("Virus at (%d, %d)\n", x, y)
		c := s.getCell(x, y)
		switch c {
		case 0: // clean
			dir = TurnLeft(dir)
			s.setCell(x, y, 1)
			infections++
			x, y = Move(dir, x, y)
			//fmt.Printf("   -> turns left and moves %s to (%d, %d)\n", dirStr(dir), x, y)
			break
		case 1:
			dir = TurnRight(dir)
			s.setCell(x, y, 0)
			x, y = Move(dir, x, y)
			//fmt.Printf("   -> turns right and moves %s to (%d, %d)\n", dirStr(dir), x, y)
			break
		}
		iterations++
	}

	return infections
}

func (s *Solution) solve2() int {
	// Fill the dynamic grid
	s.dgrid = map[uint64]byte{}
	for gy, gl := range s.grid {
		for gx, gc := range gl {
			s.setCell(gx, gy, gc)
		}
	}

	fmt.Printf("Grid size (%d, %d)\n", len(s.grid[0]), len(s.grid))

	// Put the virus in the middle of the grid
	x := len(s.grid[0]) / 2
	y := len(s.grid) / 2
	dir := DIR_UP

	fmt.Printf("Virus starts at (%d, %d)\n", x, y)

	infections := 0
	iterations := 0
	maxIterations := 10000000
	for iterations < maxIterations {
		//fmt.Printf("Virus at (%d, %d)\n", x, y)
		c := s.getCell(x, y)
		switch c {
		case 0: // clean -> weakened
			dir = TurnLeft(dir)
			s.setCell(x, y, 2)
			x, y = Move(dir, x, y)
			//fmt.Printf("   -> turns left and moves %s to (%d, %d)\n", dirStr(dir), x, y)
			break
		case 1: // infected -> flagged
			dir = TurnRight(dir)
			s.setCell(x, y, 3)
			x, y = Move(dir, x, y)
			//fmt.Printf("   -> turns right and moves %s to (%d, %d)\n", dirStr(dir), x, y)
			break
		case 2: // weakened -> infected
			s.setCell(x, y, 1)
			infections++
			x, y = Move(dir, x, y)
			break
		case 3: // flagged -> clean
			s.setCell(x, y, 0)
			dir = TurnRight(dir)
			dir = TurnRight(dir)
			x, y = Move(dir, x, y)
			break
		}
		iterations++
	}

	return infections
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
