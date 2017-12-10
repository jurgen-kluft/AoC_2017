package day9

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

func readLineToStream(line string, stream *Stream) {
	for _, c := range line {
		stream.stream = append(stream.stream, c)
	}
	return
}

func read(filename string) (stream *Stream) {
	stream = &Stream{}
	stream.stream = []rune{}
	stream.pos = 0
	stream.metrics = &Metrics{}

	reader := func(line string) {
		readLineToStream(line, stream)
	}
	iterateOverLinesInTextFile(filename, reader)

	return stream
}

// Stream is used to read characters
type Stream struct {
	stream  []rune
	metrics *Metrics
	pos     int
}

func (s *Stream) readCharacter() (rune, bool) {
	if s.pos < len(s.stream) {
		r := s.stream[s.pos]
		s.pos++
		return r, false
	}
	return 0, true
}
func (s *Stream) readCharacterIf(c rune) (rune, bool) {
	if s.pos < len(s.stream) {
		r := s.stream[s.pos]
		if r == c {
			s.pos++
		}
		return r, false
	}
	return 0, true
}

// Metrics keeps track of all the events that appear in the stream
type Metrics struct {
	score        int
	group        int
	garbageCount int
}

func (m *Metrics) init() {
	m.score = 0
	m.group = 0
	m.garbageCount = 0
}

func (m *Metrics) beginGarbage() {
	fmt.Print("Begin Garbage\n")
}
func (m *Metrics) garbage() {
	m.garbageCount++
}
func (m *Metrics) endGarbage() {
	fmt.Print("End Garbage\n")
}
func (m *Metrics) beginGroup() {
	m.score += m.group
	m.group++
	fmt.Printf("Begin Group, %d\n", m.group)
}
func (m *Metrics) endGroup() {
	fmt.Print("End Group\n")
	m.group--
}

func readIgnore(stream *Stream) {
	c, _ := stream.readCharacter()
	fmt.Printf("Ignore %v\n", c)
}

func readGarbage(stream *Stream) {
	stream.metrics.beginGarbage()
	for true {
		cc, eos := stream.readCharacter()
		if eos {
			break
		}
		if cc == '!' {
			readIgnore(stream)
		} else if cc == '>' {
			stream.metrics.endGarbage()
			break
		} else {
			stream.metrics.garbage()
		}
	}
}

func readGroup(stream *Stream) {
	stream.metrics.beginGroup()
	for true {
		cc, eos := stream.readCharacter()
		if eos {
			break
		}
		if cc == '{' {
			readGroup(stream)
		} else if cc == '!' {
			readIgnore(stream)
		} else if cc == '<' {
			readGarbage(stream)
		} else if cc == '}' {
			stream.metrics.endGroup()
			break
		}
	}
}

// Run1 is the primary solution
func Run1() {
	var stream = read("day9/input.text")
	readGroup(stream)
	fmt.Printf("Day 9.1: Score: %v \n", stream.metrics.score)
}

// Run2 is the secondary solution
func Run2() {
	var stream = read("day9/input.text")
	readGroup(stream)
	fmt.Printf("Day 9.2: Garbage count: %v \n", stream.metrics.garbageCount)
}
