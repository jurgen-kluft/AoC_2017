package day14

import (
	"fmt"
	//"strconv"
	//"strings"
)

// Day define the specific puzzle from 'Advent of Code'
const Day = "Day14"

// KnotHash invented by the elfs
type KnotHash struct {
	lengths []int
	numbers []int
}

func (h *KnotHash) init() {
	h.lengths = []int{}
	h.numbers = []int{}

	i := 0
	for i < 256 {
		h.numbers = append(h.numbers, i)
		i++
	}
}

func (h *KnotHash) reverse(start, length int) {
	if length <= 1 {
		return
	}

	size := len(h.numbers)
	end := (start + length - 1) % size

	for true {
		n1 := h.numbers[start]
		n2 := h.numbers[end]
		h.numbers[start] = n2
		h.numbers[end] = n1

		start = (start + 1) % size
		if start == end {
			break
		}
		end = (end + size - 1) % size
		if start == end {
			break
		}
	}
}
func (h *KnotHash) compute() []byte {
	pos := 0
	skip := 0
	loop := 0
	for loop < 64 {
		for _, l := range h.lengths {
			if int(l) <= len(h.numbers) {
				h.reverse(pos, int(l))
				pos = (pos + int(l) + skip) % len(h.numbers)
				skip++
			}
		}
		loop++
	}
	return h.computeDenseHash()
}

func (h *KnotHash) computeDenseHash() []byte {
	hash := []byte{}
	g := 0
	b := 0
	for _, v := range h.numbers {
		b = int(byte(b) ^ byte(v))
		g++
		if g == 16 {
			hash = append(hash, byte(b))
			b = 0
			g = 0
		}
	}
	return hash
}

// Solution is the hex-grid
type Solution struct {
	key string
}

func numberOfBitsSetInByte(b byte) int {
	cnt := 0
	for b != 0 {
		if (b & 1) != 0 {
			cnt++
		}
		b = b >> 1
	}
	return cnt
}

func computeNumberOfBitsSet(hash []byte) int {
	cnt := 0
	for _, b := range hash {
		cnt += numberOfBitsSetInByte(b)
	}
	return cnt
}

func byteToChars(b byte) (c1 rune, c2 rune) {
	b1 := ((b >> 4) & 0xF)
	if b1 >= 0 && b1 <= 9 {
		c1 = '0' + rune(b1)
	} else if b1 >= 10 && b1 <= 15 {
		c1 = 'a' + rune(b1-10)
	}
	b2 := ((b >> 0) & 0xF)
	if b2 >= 0 && b2 <= 9 {
		c2 = '0' + rune(b2)
	} else if b2 >= 10 && b2 <= 15 {
		c2 = 'a' + rune(b2-10)
	}
	return c1, c2
}
func (s *Solution) solve() int {
	row := 0
	cnt := 0
	for row < 128 {
		hashstr := fmt.Sprintf("%s-%d", s.key, row)
		// compute knot-hash of this string
		knothash := KnotHash{}
		knothash.init()
		for _, c := range hashstr {
			knothash.lengths = append(knothash.lengths, int(c))
		}
		knothash.lengths = append(knothash.lengths, int(17))
		knothash.lengths = append(knothash.lengths, int(31))
		knothash.lengths = append(knothash.lengths, int(73))
		knothash.lengths = append(knothash.lengths, int(47))
		knothash.lengths = append(knothash.lengths, int(23))

		hash := knothash.compute()
		//fmt.Printf("KnotHash, input: %s, answer(%d): ", hashstr, len(hash))
		//for _, b := range hash {
		//	c1, c2 := byteToChars(b)
		//	fmt.Printf("%s%s", string(c1), string(c2))
		//}
		//fmt.Print("\n")

		cnt += computeNumberOfBitsSet(hash)

		row++
	}
	return cnt
}

func floodfill(grid [][]int, row int, col int, region int) {
	if row >= 0 && row < 128 && col >= 0 && col < 128 && grid[row][col] == 1 {
		grid[row][col] = region
		floodfill(grid, row+1, col, region)
		floodfill(grid, row-1, col, region)
		floodfill(grid, row, col-1, region)
		floodfill(grid, row, col+1, region)
	}
}

func (s *Solution) solve2() int {
	row := 0
	grid := [][]int{}
	for row < 128 {
		hashstr := fmt.Sprintf("%s-%d", s.key, row)
		// compute knot-hash of this string
		knothash := KnotHash{}
		knothash.init()
		for _, c := range hashstr {
			knothash.lengths = append(knothash.lengths, int(c))
		}
		knothash.lengths = append(knothash.lengths, int(17))
		knothash.lengths = append(knothash.lengths, int(31))
		knothash.lengths = append(knothash.lengths, int(73))
		knothash.lengths = append(knothash.lengths, int(47))
		knothash.lengths = append(knothash.lengths, int(23))

		hash := knothash.compute()

		gridline := []int{}
		for _, h := range hash {
			bit := byte(0x80)
			for bit != 0 {
				if (h & bit) == bit {
					gridline = append(gridline, 1)
				} else {
					gridline = append(gridline, 0)
				}
				bit = bit >> 1
			}
		}
		grid = append(grid, gridline)

		row++
	}

	// Find adjecent regions
	// Mark regions with their own number
	region := 2
	row = 0
	for row < 128 {
		col := 0
		for col < 128 {
			if grid[row][col] == 1 {
				fmt.Printf("Flood fill at (%d,%d) region %d\n", row, col, region)
				floodfill(grid, row, col, region)
				region++
			}
			col++
		}
		row++
	}

	return region - 2
}

func read(filename string) *Solution {
	s := &Solution{}
	s.key = "stpzcrnm"
	//s.key = "flqrgnkx"
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
