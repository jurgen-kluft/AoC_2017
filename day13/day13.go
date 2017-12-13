package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Day = "Day13"

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
	layers   map[int]int
	scanner  map[int]int
	maxLayer int
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func (s *Solution) getScanner(layer int) (depthExists bool, depthNow int) {
	depth, exists := s.scanner[layer]
	return exists, abs(depth)
}

func (s *Solution) stepScanner(layer int) {
	depth, _ := s.scanner[layer]
	deptNow := depth

	if deptNow >= 1 {
		depthMax, _ := s.layers[layer]
		if deptNow == depthMax {
			deptNow = -deptNow
		}
	}
	s.scanner[layer] = (deptNow + 1)
}

func (s *Solution) getSeverity(layer int) int {
	maxDepth, exists := s.layers[layer]
	if exists {
		return layer * (maxDepth + 1)
	}
	return 0
}

func (s *Solution) solve() int {
	// All scanners start at depth 0
	for firewallLayer := range s.layers {
		s.scanner[firewallLayer] = 0
	}

	severity := 0
	packetLayer := -1
	for packetLayer < s.maxLayer {
		// Packet move layer
		packetLayer++

		// Is a scanner at the packet layer ?
		scannerExists, scannerDepth := s.getScanner(packetLayer)
		if scannerExists && scannerDepth == 0 {
			// We are caught
			severity += s.getSeverity(packetLayer)
		}

		// Step scanners at every layer
		for firewallLayer := range s.layers {
			s.stepScanner(firewallLayer)
		}
	}
	return severity
}

func (s *Solution) solve1() int {
	// All scanners start at depth 0, but we shift them in time so
	// that we can read them out all at once.
	for firewallLayer := range s.layers {
		s.scanner[firewallLayer] = 0
		t := 0
		for t < firewallLayer {
			s.stepScanner(firewallLayer)
			t++
		}
	}

	severity := 0
	// Step scanners at every layer
	for firewallLayer := range s.layers {
		scannerExists, scannerDepth := s.getScanner(firewallLayer)
		if scannerExists && scannerDepth == 0 {
			// We are caught
			severity += s.getSeverity(firewallLayer)
		}
	}
	return severity
}

func (s *Solution) solve2() int {
	// All scanners start at depth 0, but we shift them in time so
	// that we can read them out all at once.
	for firewallLayer := range s.layers {
		s.scanner[firewallLayer] = 0
		t := 0
		for t < firewallLayer {
			s.stepScanner(firewallLayer)
			t++
		}
	}

	delay := 0
	for true {
		caught := false
		for firewallLayer := range s.layers {
			depth, _ := s.scanner[firewallLayer]
			if depth == 0 {
				caught = true
				break
			}
		}
		if !caught {
			break
		}

		// Step scanners at every layer once
		for firewallLayer := range s.layers {
			s.stepScanner(firewallLayer)
		}

		delay++
	}
	return delay
}

func readLine(line string) (ok bool, layer int, depth int) {
	result := strings.Split(line, ":")
	result[0] = strings.TrimSpace(result[0])
	result[1] = strings.TrimSpace(result[1])

	layer64, err := strconv.ParseInt(result[0], 10, 32)
	if err != nil {
		return false, 0, 0
	}
	depth64, err := strconv.ParseInt(result[1], 10, 32)
	if err != nil {
		return false, 0, 0
	}

	layer = int(layer64)
	depth = int(depth64)

	// fmt.Printf("l: %v, d: %v\n", layer, depth)

	return true, layer, depth
}

func read(filename string) *Solution {
	s := &Solution{}
	s.layers = map[int]int{}
	s.scanner = map[int]int{}
	s.maxLayer = 0

	reader := func(line string) {
		ok, layer, depth := readLine(line)
		if ok {
			s.layers[layer] = depth - 1
			if layer > s.maxLayer {
				s.maxLayer = layer
			}
		} else {
			fmt.Printf("Read Error, %s", line)
		}
	}
	iterateOverLinesInTextFile(filename, reader)

	return s
}

// Run1 is the primary solution
func Run1() {
	var s = read(Day + "/input.text")
	var result = s.solve1()
	fmt.Printf(Day+".1: . : %v \n", result)
}

// Run2 is the secondary solution
func Run2() {
	var s = read(Day + "/input.text")
	var result = s.solve2()
	fmt.Printf(Day+".2: . : %v \n", result)
}
