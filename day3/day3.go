package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func readNumber(line string, cursor int) (number int, ok bool) {
	start := cursor
	for start < len(line) && (line[start] == ' ' || line[start] == '\t') {
		start++
	}
	end := start
	for end < len(line) && (line[end] != ' ' && line[end] != '\t') {
		end++
	}

	if start == end {
		number = 0
		ok = false
		return
	}

	numberStr := string(line[start:end])
	//fmt.Printf("%v(%v,%v) - ", numberStr, start, end)

	s64, err := strconv.ParseInt(numberStr, 10, 32)
	ok = err == nil
	number = int(s64)
	return
}

func readInput(filename string) (input int, ok bool) {
	input = 0

	computator := func(line string) {
		number, ok := readNumber(line, 0)
		if ok {
			input = number
		}
	}

	iterateOverLinesInTextFile(filename, computator)
	ok = true
	return
}

func computeManhattanDistance(number int) (steps int) {
	steps = 0

	// var ring = ((number - 2) / 8) + 1

	// Number starts at 1 and goes out into infinity
	//          Compute (x,y) where number value 1 is (0,0) on ring(0)
	//   0  -   7  |   8 |   Ring(1) has numbers    2 to    9  where   9  is at (1,-1) on ring(1)
	//   8  -  23  |  16 |   Ring(2) has numbers   10 to   25  where  10  is at (2,-2) on ring(2)
	//  24  -  47  |  24 |   Ring(3) has numbers   26 to   49  where  25  is at (3,-3) on ring(3)
	//  48  -  79  |  32 |   Ring(4) has numbers   50 to   81  where  50  is at (4,-4) on ring(4)
	//  80  - 119  |  40 |   Ring(5) has numbers   82 to  121  where  82  is at (5,-5) on ring(5)
	// 120  - 167  |  48 |   Ring(5) has numbers  122 to  169  where 122  is at (5,-5) on ring(6)

	// Like this:
	//  ..              <---    31
	//      17  16  15  14  13  30
	//      18   5   4   3  12  29
	//      19   6   1  |2| 11  28  ..
	//      20   7   8   9 |10| 27  52
	//      21  22  23  24  25 |26| 51
	//  ..  --> ..  ..  ..  48  49 |50|

	// 8 * (N * ((N + 1) / 2)) = X

	// Solved with pen + paper
	steps = 480

	return
}

type ring struct {
	w      int
	right  []int
	top    []int
	left   []int
	bottom []int
}

func (r *ring) init(w int) {
	r.w = w

	size := r.w * 2
	r.right = make([]int, size, size)
	r.top = make([]int, size, size)
	r.left = make([]int, size, size)
	r.bottom = make([]int, size, size)
}

func (r *ring) sequence() {
	if r.w > 0 {
		N := 2 + (8 * ((r.w - 1) * (r.w / 2)))
		S := 0
		P := r.w * 2
		for S < (4 * P) {
			if S < P {
				I := S - (0 * P)
				r.right[I] = N + S
			} else if S < (2 * P) {
				I := S - (1 * P)
				r.top[I] = N + S
			} else if S < (3 * P) {
				I := S - (2 * P)
				r.left[I] = N + S
			} else if S < (4 * P) {
				I := S - (3 * P)
				r.bottom[I] = N + S
			}
			S++
		}
	}
}

func (r *ring) clear() {
	size := r.w * 2
	for i := 0; i < size; i++ {
		r.right[i] = 0
		r.top[i] = 0
		r.left[i] = 0
		r.bottom[i] = 0
	}
}

func (r *ring) set(I int, V int) {
	P := r.w * 2
	if I < P {
		r.right[I] = V
	} else if I < (2 * P) {
		I = I - (1 * P)
		r.top[I] = V
	} else if I < (3 * P) {
		I = I - (2 * P)
		r.left[I] = V
	} else if I < (4 * P) {
		I = I - (3 * P)
		r.bottom[I] = V
	}
}

func (r *ring) get(x, y int) (V int, I int) {
	if r.w == 0 && x == 0 && y == 0 {
		return 1, 0
	}

	S := r.w * 2 // Length of one side
	if x == r.w {
		if y > -r.w && y <= r.w {
			i := ((y - 1) + r.w)
			return r.right[i], i
		}
	} else if x == -r.w {
		if y < r.w && y >= -r.w {
			i := r.w - (y + 1)
			return r.left[i], (S * 2) + i
		}
	}

	if y == r.w {
		if x >= -r.w && x < r.w {
			i := S - ((x + 1) + r.w)
			return r.top[i], (S * 1) + i
		}
	} else if y == -r.w {
		if x > -r.w && x <= r.w {
			i := (x - 1) + r.w
			return r.bottom[i], (S * 3) + i
		}
	}

	return 0, 0
}

func (r *ring) getValue(x, y int) int {
	value, _ := r.get(x, y)
	return value
}

func (r *ring) location(S int) (x, y int) {
	x = 0
	y = 0

	if r.w > 0 {
		if S == 0 {
			x = r.w
			y = -r.w + 1
		} else {
			P := r.w * 2
			if S < P {
				x = r.w
				y = -r.w + 1 + (S - (0 * P))
			} else if S < (2 * P) {
				x = r.w - 1 - (S - (1 * P))
				y = r.w
			} else if S < (3 * P) {
				x = -r.w
				y = r.w - 1 - (S - (2 * P))
			} else if S < (4 * P) {
				x = -r.w + 1 + (S - (3 * P))
				y = -r.w
			}
		}
	}
	return
}

func (r *ring) iterate(S int) (x, y int, nextS int, done bool) {
	x = 0
	y = 0

	if r.w == 0 {
		nextS = 1
		done = true
	} else {
		nextS = S + 1
		done = (nextS == (8 * r.w))
		x, y = r.location(S)
	}
	return
}

func (r *ring) sum(x, y int, I int) int {
	value, index := r.get(x, y)
	if index < I {
		return value
	}
	return 0
}

func (r *ring) expand(inner *ring) {
	r.init(inner.w + 1)

	I := 0
	for true {
		x, y, next, done := r.iterate(I)

		// get the value of our current location
		sum, _ := r.get(x, y)

		// sum all our own 'valid' neighbors
		sum += r.sum(x+1, y, I)
		sum += r.sum(x-1, y, I)
		sum += r.sum(x+1, y+1, I)
		sum += r.sum(x-1, y+1, I)
		sum += r.sum(x+1, y-1, I)
		sum += r.sum(x-1, y-1, I)
		sum += r.sum(x, y+1, I)
		sum += r.sum(x, y-1, I)

		sum += inner.getValue(x+1, y)
		sum += inner.getValue(x-1, y)
		sum += inner.getValue(x+1, y+1)
		sum += inner.getValue(x-1, y+1)
		sum += inner.getValue(x+1, y-1)
		sum += inner.getValue(x-1, y-1)
		sum += inner.getValue(x, y+1)
		sum += inner.getValue(x, y-1)

		r.set(I, sum)

		if done {
			break
		}

		I = next
	}
}

func (r *ring) find(number int) (x, y int, found bool) {
	S := 0
	P := r.w * 2
	for S < (4 * P) {
		if S < P {
			I := S - (0 * P)
			if r.right[I] > number {
				x, y = r.location(S)
				return x, y, true
			}
		} else if S < (2 * P) {
			I := S - (1 * P)
			if r.top[I] > number {
				x, y = r.location(S)
				return x, y, true
			}
		} else if S < (3 * P) {
			I := S - (2 * P)
			if r.left[I] > number {
				x, y = r.location(S)
				return x, y, true
			}
		} else if S < (4 * P) {
			I := S - (3 * P)
			if r.bottom[I] > number {
				x, y = r.location(S)
				return x, y, true
			}
		}
		S++
	}

	return 0, 0, false
}

func (r *ring) print() {
	I := 0
	fmt.Printf("Ring = %v\n", r.w)
	for true {
		x, y, next, done := r.iterate(I)
		value, _ := r.get(x, y)
		fmt.Printf("(%v,%v)=%v\n", x, y, value)
		if done {
			break
		}
		I = next
	}
}

func computeManhattanDistanceModified(number int) (steps int) {
	steps = 0

	rings := [2]*ring{}
	rings[0] = &ring{}
	rings[1] = &ring{}

	inner := 0
	outer := 1

	fx := 0
	fy := 0
	fv := 0
	found := false

	rings[inner].init(0)
	rings[inner].sequence()
	for i := 0; i < 30; i++ {
		//rings[inner].print()
		rings[outer].expand(rings[inner])

		fx, fy, found = rings[outer].find(number)
		if found {
			fv = rings[outer].getValue(fx, fy)
			break
		}

		// swap inner and outer
		inner = 1 - inner
		outer = 1 - outer
	}

	if found {
		if fx < 0 {
			fx = -fx
		}
		if fy < 0 {
			fy = -fy
		}

		fmt.Printf("Found %v at (%v,%v)\n", number, fx, fy)
		fmt.Printf("Location at (%v,%v) = %v\n", fx, fy, fv)

		steps = fx + fy
	}

	return steps
}

// Run1 is the primary solution
func Run1() {
	var input, ok = readInput("day3/input.text")
	var steps = computeManhattanDistance(input)

	if ok {
		fmt.Printf("Day 3.1: Steps are %v \n", steps)
	} else {
		fmt.Printf("Could not process the input")
	}
}

// Run2 is the secondary solution
func Run2() {
	input, ok := readInput("day3/input.text")
	steps := computeManhattanDistanceModified(input)

	if ok {
		fmt.Printf("Day 3.2: Steps are %v \n", steps)
	} else {
		fmt.Printf("Could not process the input")
	}
}
