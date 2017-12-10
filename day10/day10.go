package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (h *KnotHash) compute() int {
	pos := 0
	skip := 0
	for _, l := range h.lengths {
		if int(l) <= len(h.numbers) {
			h.reverse(pos, int(l))
			pos = (pos + int(l) + skip) % len(h.numbers)
			skip++
		}
	}
	return int(h.numbers[0]) * int(h.numbers[1])
}

func (h *KnotHash) compute2() string {
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

func (h *KnotHash) computeDenseHash() string {
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
	s := ""
	for _, b := range hash {
		c1, c2 := byteToChars(b)
		s = s + string(c1) + string(c2)
	}
	return s
}

func readLengths(line string, kh *KnotHash) {
	numbers := strings.Split(line, ",")
	for _, number := range numbers {
		l, _ := strconv.ParseInt(number, 10, 32)
		kh.lengths = append(kh.lengths, int(l))
	}
	return
}

func read(filename string) *KnotHash {
	kh := &KnotHash{}
	kh.init()

	reader := func(line string) {
		readLengths(line, kh)
	}
	iterateOverLinesInTextFile(filename, reader)

	return kh
}

func readLengths2(line string, kh *KnotHash) {
	for _, c := range line {
		kh.lengths = append(kh.lengths, int(c))
	}
	return
}

func read2(filename string) *KnotHash {
	kh := &KnotHash{}
	kh.init()

	reader := func(line string) {
		readLengths2(line, kh)
	}
	iterateOverLinesInTextFile(filename, reader)

	kh.lengths = append(kh.lengths, int(17))
	kh.lengths = append(kh.lengths, int(31))
	kh.lengths = append(kh.lengths, int(73))
	kh.lengths = append(kh.lengths, int(47))
	kh.lengths = append(kh.lengths, int(23))

	return kh
}

// Run1 is the primary solution
func Run1() {
	var kh = read("day10/input.text")
	var result = kh.compute()
	fmt.Printf("Day 10.1: First 2 numbers multiplied: %v \n", result)
}

// Run2 is the secondary solution
func Run2() {
	var kh = read2("day10/input.text")
	var result = kh.compute2()
	fmt.Printf("Day 10.2: Hash of input: %v \n", result)
}
