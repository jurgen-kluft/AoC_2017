package day16

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

func readLine(line string, s *Solution) {
	moves := strings.Split(line, ",")

	for _, move := range moves {
		runes := []rune(move)

		instr := &Instruction{}
		instr.opcode = '?'
		switch runes[0] {
		case 's':
			instr.opcode = 's'
			break
		case 'x':
			instr.opcode = 'x'
			break
		case 'p':
			instr.opcode = 'p'
			break
		}
		if instr.opcode == 'p' {
			instr.left = int(runes[1]) - 'a'
			instr.right = int(runes[3]) - 'a'
		}
		if instr.opcode == 'x' {
			slash := 2
			for runes[slash] != '/' {
				slash++
			}
			l, errl := strconv.ParseInt(string(runes[1:slash]), 10, 32)
			r, errr := strconv.ParseInt(string(runes[slash+1:]), 10, 32)
			if errl == nil && errr == nil {
				instr.left = int(l)
				instr.right = int(r)
			} else {
				fmt.Println("Error reading dance instruction")
			}
		}
		if instr.opcode == 's' {
			s, err := strconv.ParseInt(string(runes[1:]), 10, 32)
			if err == nil {
				instr.left = int(s)
			}
		}
		if instr.opcode == '?' {
			fmt.Println("Error reading dance instruction")
		}
		s.instructions = append(s.instructions, instr)
	}
	return
}

// Instruction is a dance move
type Instruction struct {
	opcode int
	left   int
	right  int
}

// Solution is program dance
type Solution struct {
	instructions []*Instruction
	place        int
	places       []byte
	programs     []byte
	work         []byte
}

func printOrder(order []byte) {
	i := 0
	for i < 16 {
		program := order[i]
		char := rune('a' + program)
		fmt.Printf("%s", string(char))
		i++
	}
	fmt.Println()
}

func orderString(order []byte) string {
	runes := make([]rune, 16, 16)
	for i, b := range order {
		runes[i] = rune('a' + b)
	}
	return string(runes)
}

func (s *Solution) exchange(x int, y int) {
	p1 := s.places[(s.place+x)%16]
	p2 := s.places[(s.place+y)%16]
	s.places[(s.place+x)%16] = p2
	s.places[(s.place+y)%16] = p1
	s.programs[p1] = byte((s.place + y) % 16)
	s.programs[p2] = byte((s.place + x) % 16)
}

func (s *Solution) partner(p1 int, p2 int) {
	place1 := s.programs[p1]
	place2 := s.programs[p2]
	s.programs[p1] = place2
	s.programs[p2] = place1
	s.places[place1] = byte(p2)
	s.places[place2] = byte(p1)
}

func (s *Solution) spin(spin int) {
	pos := 0
	for pos < 16 {
		s.work[(pos+spin)%16] = s.places[pos]
		pos++
	}
	pos = 0
	for pos < 16 {
		s.places[pos] = s.work[pos]
		pos++
	}
	pos = 0
	for pos < 16 {
		s.programs[pos] = byte((int(s.programs[pos]) + spin) % 16)
		pos++
	}
}

func (s *Solution) solve() int {
	s.place = 0
	s.places = []byte{}
	s.programs = []byte{}
	s.work = []byte{}

	program := byte(0)
	for program < 16 {
		s.places = append(s.places, program)
		s.programs = append(s.programs, program)
		s.work = append(s.work, 0)
		program++
	}

	for _, i := range s.instructions {
		switch i.opcode {
		case 's':
			// spin X
			//s.place = (s.place + 16 - i.left) % 16
			s.spin(i.left)
			break
		case 'x':
			s.exchange(i.left, i.right)
			break
		case 'p':
			s.partner(i.left, i.right)
			break
		}
		orderString(s.places)
	}
	printOrder(s.places)
	return 0
}

func (s *Solution) solve2() int {
	s.places = []byte{}
	s.programs = []byte{}
	s.work = []byte{}

	program := byte(0)
	for program < 16 {
		s.places = append(s.places, program)
		s.programs = append(s.programs, program)
		s.work = append(s.work, 0)
		program++
	}

	orders := map[int64]string{}
	history := map[string]int64{}

	iter := int64(0)
	iterations := int64(0)
	for iterations < int64(1000*1000*1000) {
		for _, i := range s.instructions {
			switch i.opcode {
			case 's':
				s.spin(i.left)
				break
			case 'x':
				s.exchange(i.left, i.right)
				break
			case 'p':
				s.partner(i.left, i.right)
				break
			}
		}
		order := orderString(s.places)
		existingIter, exists := history[order]
		if exists {
			iter = existingIter
			fmt.Printf("Had a previous order at iteration %d, %d\n", iter, iterations)
			break
		} else {
			orders[iterations] = order
			history[order] = iterations
		}
		iterations++
	}

	rest := int64(1000*1000*1000) % 60

	fmt.Println(orders[rest-1])

	return 0
}

func (s *Solution) solve3() int {
	s.places = []byte{}

	program := byte(0)
	for program < 16 {
		s.places = append(s.places, program)
		program++
	}

	// Keep applying the transform
	workOrPlaces := 0
	iterations := int64(0)
	for iterations < int64(1000*1000*1000) {
		if iterations&1 == 0 {
			// places -> work
			for i, move := range s.programs {
				s.work[move] = s.places[i]
			}
			workOrPlaces = 1
		} else {
			// work -> places
			for i, move := range s.programs {
				s.places[move] = s.work[i]
			}
			workOrPlaces = 0
		}
		if iterations&0xFFFFF == 0 {
			fmt.Printf("Iteration %v\n", iterations)
		}
		iterations++
	}

	if workOrPlaces == 0 {
		printOrder(s.places)
	} else if workOrPlaces == 1 {
		printOrder(s.work)
	}

	// wrong: fbcdegahionkmlpj

	return 0
}

func read(filename string) *Solution {
	s := &Solution{}
	s.instructions = []*Instruction{}

	reader := func(line string) {
		readLine(line, s)
	}
	iterateOverLinesInTextFile(filename, reader)

	return s
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
