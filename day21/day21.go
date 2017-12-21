package day21

import (
	"bufio"
	"fmt"
	"math/bits"
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

func readLine(line string, s *Solution) (uint, uint) {
	src := uint(0)
	srcc := uint(1)
	dst := uint(0)
	dstc := uint(1)

	bit := uint(0x8000)
	state := 0
	for _, c := range line {
		if state == 0 {
			if c == '>' {
				state = 1
				bit = uint(0x8000)
			} else {
				if c == '/' {
					srcc++
				} else if c == '.' {
					bit = bit >> 1
				} else if c == '#' {
					src = src | bit
					bit = bit >> 1
				}
			}
		} else if state == 1 {
			if c == '/' {
				dstc++
			} else if c == '.' {
				bit = bit >> 1
			} else if c == '#' {
				dst = dst | bit
				bit = bit >> 1
			}
		}
	}
	if srcc == 2 {
		src = src >> 12
	} else if srcc == 3 {
		src = src >> 7
	}
	if dstc == 2 {
		dst = dst >> 12
	} else if dstc == 3 {
		dst = dst >> 7
	}
	src = (src << 4) | srcc
	dst = (dst << 4) | dstc

	return src, dst
}

func read(filename string) *Solution {
	s := &Solution{}
	s.transforms = map[uint]uint{}

	reader := func(line string) {
		src, dst := readLine(line, s)
		s.transforms[src] = dst
	}
	iterateOverLinesInTextFile(filename, reader)

	//s.printTransforms()
	return s
}

// 2x2, 4 bits pattern:
//    3 2
//    1 0
func rotate2x2(grid uint) uint {
	rot := uint(0)
	pat := (grid & 0xF0) >> 4
	rot = rot | ((pat & (1 << 1)) << 2) // bit 3 = bit 1
	rot = rot | ((pat & (1 << 3)) >> 1) // bit 2 = bit 3
	rot = rot | ((pat & (1 << 0)) << 1) // bit 1 = bit 0
	rot = rot | ((pat & (1 << 2)) >> 2) // bit 0 = bit 2
	if bits.OnesCount(pat) != bits.OnesCount(rot) {
		fmt.Println("Error, rotate2x2 bug")
	}
	return (rot << 4) | 2
}

// 3x3, 9 bits pattern:
//    8 7 6
//    5 4 3
//    2 1 0
func rotate3x3(grid uint) uint {
	pat := (grid & 0xFFF0) >> 4
	rot := uint(0)
	rot = rot | ((pat & (1 << 5)) << 3) // bit 8 = bit 5
	rot = rot | ((pat & (1 << 8)) >> 1) // bit 7 = bit 8
	rot = rot | ((pat & (1 << 7)) >> 1) // bit 6 = bit 7
	rot = rot | ((pat & (1 << 2)) << 3) // bit 5 = bit 2
	rot = rot | (pat & (1 << 4))        // bit 4
	rot = rot | ((pat & (1 << 6)) >> 3) // bit 3 = bit 6
	rot = rot | ((pat & (1 << 1)) << 1) // bit 2 = bit 1
	rot = rot | ((pat & (1 << 0)) << 1) // bit 1 = bit 0
	rot = rot | ((pat & (1 << 3)) >> 3) // bit 0 = bit 3
	if bits.OnesCount(pat) != bits.OnesCount(rot) {
		fmt.Println("Error, rotate3x3 bug")
	}
	return (rot << 4) | 3
}

func flipPattern(pattern uint) uint {
	size := pattern & 0xF
	if size == 2 {
		flipped := (pattern & uint(0x30)) << 2
		flipped = flipped | (pattern&uint(0xC0))>>2
		return flipped | size
	} else if size == 3 {
		flipped := (pattern & uint(0x0070)) << 6
		flipped = flipped | (pattern&uint(0x1C00))>>6
		flipped = flipped | (pattern & uint(0x0380))
		return flipped | size
	}
	return pattern
}

func printPattern(p uint) {
	size := p & 0xF
	bit := uint(0x10 << ((size * size) - 1))
	r := size
	for r > 0 {
		bits := uint(0)
		for bits < size {
			if p&bit == bit {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			bits++
			bit = bit >> 1
		}
		r--
		if r != 0 {
			fmt.Print("/")
		}
	}

}

func (s *Solution) printTransforms() {
	for key, value := range s.transforms {
		if key&0xF == 2 {
			printPattern(key)
			fmt.Print(" => ")
			printPattern(value)
			fmt.Println()
		}
	}
	for key, value := range s.transforms {
		if key&0xF == 3 {
			printPattern(key)
			fmt.Print(" => ")
			printPattern(value)
			fmt.Println()
		}
	}
}

func (s *Solution) addTransform(key uint, value uint) {
	if _, ok := s.transforms[key]; !ok {
		s.transforms[key] = value
	}
}

func (s *Solution) findTransform(key uint) uint {
	i := 0
	for i < 2 {
		// Rotate
		size := key & 0xF
		r := uint(0)
		nr := uint(size*size) - 1
		kr := key
		for r <= nr {
			if size == 2 {
				if rotated, ok := s.transforms[kr]; ok {
					return rotated
				}
				kr = rotate2x2(kr)
			} else if size == 3 {
				if rotated, ok := s.transforms[kr]; ok {
					return rotated
				}
				kr = rotate3x3(kr)
			}
			r++
		}

		// Flip
		key = flipPattern(key)
		i++
	}

	fmt.Println("Encountered a pattern without a transform")
	return key
}

func splitPattern4x4(pattern uint) (uint, uint, uint, uint) {
	pattern = pattern >> 4
	p1 := pattern & uint(0xCC00)
	p2 := pattern & uint(0x3300)
	p3 := pattern & uint(0x00CC)
	p4 := pattern & uint(0x0033)
	p1 = (p1 << 4) | 2
	p2 = (p2 << 4) | 2
	p3 = (p3 << 4) | 2
	p4 = (p4 << 4) | 2
	return p1, p2, p3, p4
}

// Solution is the primary object used to solve the puzzle
type Solution struct {
	transforms map[uint]uint // first 4 bits is size, then size 2 = (4 bits) / size 3 = (9 bits) / size 4 = (16 bits)
}

func pixelsToPattern(posx int, posy int, width int, pixels [][]byte) (pattern uint) {
	ps := width
	pbit := uint(1 << uint((ps*ps)-1))
	pattern = 0

	y := 0
	for y < ps {
		x := 0
		for x < ps {
			if pixels[posy+y][posx+x] == 1 {
				pattern = pattern | pbit
			}
			pbit = pbit >> 1
			x++
		}
		y++
	}
	return (pattern << 4) | uint(width)
}

func convertPixelsToPatterns(pixels [][]byte) (patterns [][]uint) {
	w := len(pixels[0])
	pw := 0
	if w%2 == 0 {
		pw = w / 2
	} else {
		pw = w / 3
	}

	// Convert pixel image into pw x pw pattern image
	patterns = make([][]uint, pw, pw)
	for h := range patterns {
		patterns[h] = make([]uint, pw, pw)
	}

	y := 0
	py := 0
	for y < w {
		x := 0
		px := 0
		for x < w {
			p := pixelsToPattern(x, y, w, pixels)
			patterns[py][px] = p
			x += w
			px++
		}
		y += w
		py++
	}

	return
}

func patternToPixels(pattern uint, posx int, posy int, pixels [][]byte) {
	ps := int(pattern & 0xF)
	pattern = pattern >> 4
	pbit := uint(1 << uint((ps*ps)-1))

	y := 0
	for y < ps {
		x := 0
		for x < ps {
			if pattern&pbit == 0 {
				pixels[posy+y][posx+x] = 0
			} else {
				pixels[posy+y][posx+x] = 1
			}
			x++
		}
		y++
	}
}

func convertPatternsToPixels(patterns [][]uint) (pixels [][]byte) {

	// width/height in pixels
	// allocate pixels
	ps := int(patterns[0][0] & 0xF)
	w := len(patterns) * ps
	pixels = make([][]byte, w, w)
	for h := range pixels {
		pixels[h] = make([]byte, w, w)
	}

	y := 0
	for _, lp := range patterns {
		x := 0
		for _, p := range lp {
			patternToPixels(p, x, y, pixels)
			x += ps
		}
		y += ps
	}

	return
}

func (s *Solution) solve() int {
	// Pattern we start with
	patterns := [][]uint{}

	// Start pattern (size = 3)
	// [.#.]
	// [..#]  -> 000(010)(001)(111) = 08F
	// [###]
	patterns = append(patterns, []uint{0x08F3})
	pixels := convertPatternsToPixels(patterns)

	iteration := 0
	iterations := 5
	for iteration < iterations {
		// According to the 'size' rule, convert pixels to patterns
		patterns = convertPixelsToPatterns(pixels)

		// Upscale all patterns
		upscaledPatternImage := [][]uint{}
		for _, pp := range patterns {
			upscaledPatternLine := []uint{}
			for _, p := range pp {
				up := s.findTransform(p)
				upscaledPatternLine = append(upscaledPatternLine, up)
			}
			upscaledPatternImage = append(upscaledPatternImage, upscaledPatternLine)
		}

		// Convert patterns back to pixelized image
		pixels = convertPatternsToPixels(patterns)
		iteration++
	}

	// Count the '1's
	pixelsSet := 0
	for _, pp := range pixels {
		for _, p := range pp {
			if p == 1 {
				pixelsSet++
			}
		}
	}

	return pixelsSet
}

func (s *Solution) solve2() int {
	return 0
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
