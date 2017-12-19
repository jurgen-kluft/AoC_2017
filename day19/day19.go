package day19

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

func readLine(line string, s *Solution) {
	row := &Row{}

	for _, c := range line {
		switch {
		case c == '-':
			row.data = append(row.data, CELL_EW)
			break
		case c == '+':
			row.data = append(row.data, CELL_NSEW)
			break
		case c == '|':
			row.data = append(row.data, CELL_NS)
			break
		case c == ' ':
			row.data = append(row.data, CELL_0)
			break
		case c >= 'A' && c <= 'Z':
			cell := CELL_CH | int(c-'A')
			row.data = append(row.data, cell)
			break
		default:
			fmt.Printf("Error in reading (%s)\n", string(c))
		}
	}

	s.rows = append(s.rows, row)
	return
}

func read(filename string) *Solution {
	s := &Solution{}
	s.rows = []*Row{}

	reader := func(line string) {
		readLine(line, s)
	}
	iterateOverLinesInTextFile(filename, reader)

	return s
}

const (
	CELL_0       = 0x00000
	CELL_E       = 0x10000
	CELL_W       = 0x20000
	CELL_N       = 0x40000
	CELL_S       = 0x80000
	CELL_NSEW    = CELL_N | CELL_S | CELL_E | CELL_W
	CELL_NS      = CELL_N | CELL_S
	CELL_EW      = CELL_E | CELL_W
	CELL_SE      = CELL_S | CELL_E
	CELL_SW      = CELL_S | CELL_W
	CELL_NE      = CELL_N | CELL_E
	CELL_NW      = CELL_N | CELL_W
	CELL_CH      = 0x100000
	CELL_CH_MASK = 0xFF

	DIR_S = 1
	DIR_N = 2
	DIR_E = 3
	DIR_W = 4
)

// Solution is the hex-grid
type Row struct {
	data []int
}

type Solution struct {
	rows []*Row
}

func (s *Solution) getDirection(d int, x int, y int) int {
	if d == DIR_W || d == DIR_E {
		cellN := s.rows[y-1].data[x]
		cellS := s.rows[y+1].data[x]
		if cellS != CELL_0 {
			return DIR_S
		} else if cellN != CELL_0 {
			return DIR_N
		}
	} else {
		cellW := s.rows[y].data[x-1]
		cellE := s.rows[y].data[x+1]
		if cellW != CELL_0 {
			return DIR_W
		} else if cellE != CELL_0 {
			return DIR_E
		}
	}
	return -1
}

func (s *Solution) shouldChangeDirection(d int, x int, y int) bool {
	cv := s.rows[y].data[x]
	return cv == CELL_NSEW
}
func (s *Solution) isEmptyCell(x int, y int) bool {
	cv := s.rows[y].data[x]
	return cv == CELL_0
}
func (s *Solution) hasCharacter(x int, y int) bool {
	cv := s.rows[y].data[x]
	return (cv & CELL_CH) == CELL_CH
}
func (s *Solution) getCharacter(x int, y int) rune {
	cv := s.rows[y].data[x]
	return rune(int(cv&CELL_CH_MASK) + 'A')
}

func (s *Solution) solve() int {
	// Find our starting point
	// Search first row for CELL_NS
	x := 0
	y := 0
	d := DIR_S
	for cx, cv := range s.rows[0].data {
		if cv == CELL_NS {
			x = cx
			break
		}
	}

	// Traverse
	chars := []rune{}
	w := len(s.rows[0].data)
	h := len(s.rows)
	fmt.Printf("Grid with dimensions (%d, %d)", w, h)
	for true {
		if s.isEmptyCell(x, y) {
			fmt.Println("Hit empty cell")
			break
		}
		if s.hasCharacter(x, y) {
			ch := s.getCharacter(x, y)
			chars = append(chars, ch)
		}
		if s.shouldChangeDirection(d, x, y) {
			d = s.getDirection(d, x, y)
		}
		switch d {
		case DIR_N:
			y--
			break
		case DIR_S:
			y++
			break
		case DIR_W:
			x--
			break
		case DIR_E:
			x++
			break
		default:
			fmt.Println("Error in direction")
			break
		}
		if x < 0 || x >= w {
			fmt.Println("Out of width bounds")
			break
		}
		if y < 0 || y >= h {
			fmt.Println("Out of height bounds")
			break
		}
		//fmt.Printf("Cursor at (%d,%d)\n", x, y)
	}

	fmt.Println(string(chars))
	return 0
}

func (s *Solution) solve2() int {
	// Find our starting point
	// Search first row for CELL_NS
	x := 0
	y := 0
	d := DIR_S
	for cx, cv := range s.rows[0].data {
		if cv == CELL_NS {
			x = cx
			break
		}
	}

	// Traverse
	steps := 0
	chars := []rune{}
	w := len(s.rows[0].data)
	h := len(s.rows)
	fmt.Printf("Grid with dimensions (%d, %d)", w, h)
	for true {
		if s.isEmptyCell(x, y) {
			fmt.Println("Hit empty cell")
			break
		}
		steps++
		if s.hasCharacter(x, y) {
			ch := s.getCharacter(x, y)
			chars = append(chars, ch)
		}
		if s.shouldChangeDirection(d, x, y) {
			d = s.getDirection(d, x, y)
		}
		switch d {
		case DIR_N:
			y--
			break
		case DIR_S:
			y++
			break
		case DIR_W:
			x--
			break
		case DIR_E:
			x++
			break
		default:
			fmt.Println("Error in direction")
			break
		}
		if x < 0 || x >= w {
			fmt.Println("Out of width bounds")
			break
		}
		if y < 0 || y >= h {
			fmt.Println("Out of height bounds")
			break
		}
		//fmt.Printf("Cursor at (%d,%d)\n", x, y)
	}

	fmt.Println(string(chars))
	return steps
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
