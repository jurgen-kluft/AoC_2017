package day21

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"strings"
)

func onesCount(v uint) int {
	cnt := 0
	bit := uint(1)
	for bit <= 0x80000000 {
		if (v & bit) == bit {
			cnt++
		}
		bit = bit << 1
	}
	return cnt
}

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
		//printTransform(src, dst)
	}
	iterateOverLinesInTextFile(filename, reader)

	//s.printTransforms()
	return s
}

// 2x2, 4 bits pattern:
//    3 2 -> 1 3
//    1 0    0 2
func rotate2x2(grid uint) uint {
	rot := uint(0)
	pat := (grid & 0xF0) >> 4
	rot |= ((pat & (1 << 1)) << 2) // bit 3 = bit 1
	rot |= ((pat & (1 << 3)) >> 1) // bit 2 = bit 3
	rot |= ((pat & (1 << 0)) << 1) // bit 1 = bit 0
	rot |= ((pat & (1 << 2)) >> 2) // bit 0 = bit 2
	if onesCount(pat) != onesCount(rot) {
		fmt.Println("Error, rotate2x2 bug")
	}
	return (rot << 4) | 2
}

// 3x3, 9 bits pattern:
//    8 7 6    5 8 7
//    5 4 3 -> 2 4 6
//    2 1 0    1 0 3
func rotate3x3(grid uint) uint {
	pat := (grid & 0x1FF0) >> 4
	rot := uint(0)
	rot |= ((pat & (1 << 5)) << 3) // bit 8 = bit 5
	rot |= ((pat & (1 << 8)) >> 1) // bit 7 = bit 8
	rot |= ((pat & (1 << 7)) >> 1) // bit 6 = bit 7
	rot |= ((pat & (1 << 2)) << 3) // bit 5 = bit 2
	rot |= (pat & (1 << 4))        // bit 4
	rot |= ((pat & (1 << 6)) >> 3) // bit 3 = bit 6
	rot |= ((pat & (1 << 1)) << 1) // bit 2 = bit 1
	rot |= ((pat & (1 << 0)) << 1) // bit 1 = bit 0
	rot |= ((pat & (1 << 3)) >> 3) // bit 0 = bit 3
	if onesCount(pat) != onesCount(rot) {
		fmt.Println("Error, rotate3x3 bug")
	}
	return (rot << 4) | 3
}

func flipPattern(pattern uint) uint {
	size := pattern & 0xF
	if size == 2 {
		// 3 2     1 0
		// 1 0  -> 3 2
		flipped := (pattern & uint(0x30)) << 2
		flipped |= ((pattern & uint(0x30<<2)) >> 2)
		return flipped | size
	} else if size == 3 {
		// 8 7 6    2 1 0
		// 5 4 3 -> 5 4 3
		// 2 1 0    8 7 6
		flipped := (pattern & uint(0x0070)) << 6
		flipped |= (pattern & uint(0x0070<<3))
		flipped |= ((pattern & uint(0x0070<<6)) >> 6)
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

func printTransform(key uint, value uint) {
	if key&0xF == 2 {
		printPattern(key)
		fmt.Print(" => ")
		printPattern(value)
		fmt.Println()
	} else if key&0xF == 3 {
		printPattern(key)
		fmt.Print(" => ")
		printPattern(value)
		fmt.Println()
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

func (s *Solution) findTransform(key uint) uint {
	// Rotate
	size := key & 0xF
	if size == 2 {
		r := uint(0)
		nr := uint(4)
		kr := key
		for r < nr {
			if kt, ok := s.transforms[kr]; ok {
				return kt
			}
			if kt, ok := s.transforms[flipPattern(kr)]; ok {
				return kt
			}
			kr = rotate2x2(kr)
			r++
		}
	} else if size == 3 {
		r := uint(0)
		nr := uint(4)
		kr := key
		for r < nr {
			if kt, ok := s.transforms[kr]; ok {
				return kt
			}
			if kt, ok := s.transforms[flipPattern(kr)]; ok {
				return kt
			}
			kr = rotate3x3(kr)
			kr = rotate3x3(kr)
			r++
		}
	}

	fmt.Println("Encountered a pattern without a transform")
	return key
}

// Solution is the primary object used to solve the puzzle
type Solution struct {
	transforms map[uint]uint // first 4 bits is size, then size 2 = (4 bits) / size 3 = (9 bits) / size 4 = (16 bits)
}

func (pixels *Pixels) toPattern(posx int, posy int, ps int) (pattern uint) {
	pbit := uint(1 << uint((ps*ps)-1))
	pattern = 0

	y := 0
	for y < ps {
		x := 0
		for x < ps {
			if pixels.pixels[posy+y][posx+x] != 0 {
				pattern = pattern | pbit
			}
			pbit = pbit >> 1
			x++
		}
		y++
	}
	return (pattern << 4) | uint(ps)
}

func (pixels *Pixels) convertToPatterns(patterns *Patterns) {
	w := len(pixels.pixels[0])
	dw := 3
	if w%2 == 0 {
		dw = 2
	}
	pw := w / dw

	// Convert pixel image into pw x pw pattern image
	patterns.patterns = make([][]uint, pw, pw)
	for h := range patterns.patterns {
		patterns.patterns[h] = make([]uint, pw, pw)
	}

	y := 0
	py := 0
	for py < pw {
		x := 0
		px := 0
		for px < pw {
			p := pixels.toPattern(x, y, dw)
			patterns.patterns[py][px] = p
			x += dw
			px++
		}
		y += dw
		py++
	}

	return
}

// Pixels is a byte based image, one pixel is one byte either '0' or '1'
type Pixels struct {
	pixels [][]byte
}

func (pixels *Pixels) countSetPixels() int {
	pixelsSet := 0
	for _, pp := range pixels.pixels {
		for _, p := range pp {
			if p != 0 {
				pixelsSet++
			}
		}
	}
	return pixelsSet
}

func (pixels *Pixels) patternToPixels(pt uint, posx int, posy int) {
	ps := int(pt & 0xF)
	pt = pt >> 4
	pb := uint(1 << uint((ps*ps)-1))

	y := 0
	for y < ps {
		x := 0
		for x < ps {
			if (pt & pb) == pb {
				pixels.pixels[posy+y][posx+x] = 0xFF
			} else {
				//pixels.pixels[posy+y][posx+x] = 0x00
			}
			pb = pb >> 1
			x++
		}
		y++
	}
}

// Patterns is the image split into patterns
type Patterns struct {
	patterns [][]uint
}

func (patterns *Patterns) convertToPixels(pixels *Pixels) {

	// width/height in pixels
	// allocate pixels
	ps := int(patterns.patterns[0][0] & 0xF)
	w := len(patterns.patterns) * ps
	pixels.pixels = make([][]byte, w, w)
	for h := range pixels.pixels {
		pixels.pixels[h] = make([]byte, w, w)
	}

	y := 0
	for _, lp := range patterns.patterns {
		x := 0
		for _, p := range lp {
			pixels.patternToPixels(p, x, y)
			x += ps
		}
		y += ps
	}

	return
}

func (s *Solution) solve() int {
	// Pattern we start with
	patterns := &Patterns{}
	patterns.patterns = [][]uint{}

	// Start pattern (size = 3)
	// [.#.]
	// [..#]  -> 000(010)(001)(111) = 08F
	// [###]
	patterns.patterns = append(patterns.patterns, []uint{0x08F3})

	pixels := &Pixels{}
	patterns.convertToPixels(pixels)
	fmt.Printf("First image (%d,%d) has %d pixels set\n", len(pixels.pixels), len(pixels.pixels), pixels.countSetPixels())

	iteration := 0
	iterations := 5
	for iteration < iterations {
		// According to the 'size' rule, convert pixels to patterns
		pixels.convertToPatterns(patterns)

		// Upscale all patterns
		upscaledPattern := &Patterns{}
		upscaledPattern.patterns = [][]uint{}
		for _, pp := range patterns.patterns {
			upscaledPatternLine := []uint{}
			for _, p := range pp {
				up := s.findTransform(p)
				upscaledPatternLine = append(upscaledPatternLine, up)
			}
			upscaledPattern.patterns = append(upscaledPattern.patterns, upscaledPatternLine)
		}

		// Convert patterns back to pixelized image
		upscaledPattern.convertToPixels(pixels)
		fmt.Printf("Iteration %d, image is (%d,%d), pixels set %d \n", iteration, len(pixels.pixels), len(pixels.pixels), pixels.countSetPixels())
		iteration++
	}

	// Final image dimensions
	fmt.Printf("Final image is (%d,%d)\n", len(pixels.pixels), len(pixels.pixels))

	// Count the '1's
	pixelsSet := pixels.countSetPixels()

	return pixelsSet
}

func (s *Solution) solve2() int {
	// Pattern we start with
	patterns := &Patterns{}
	patterns.patterns = [][]uint{}

	// Start pattern (size = 3)
	// [.#.]
	// [..#]  -> 000(010)(001)(111) = 08F
	// [###]
	patterns.patterns = append(patterns.patterns, []uint{0x08F3})

	pixels := &Pixels{}
	patterns.convertToPixels(pixels)
	fmt.Printf("First image (%d,%d) has %d pixels set\n", len(pixels.pixels), len(pixels.pixels), pixels.countSetPixels())

	iteration := 0
	iterations := 18
	for iteration < iterations {
		// According to the 'size' rule, convert pixels to patterns
		pixels.convertToPatterns(patterns)

		// Upscale all patterns
		upscaledPattern := &Patterns{}
		upscaledPattern.patterns = [][]uint{}
		for _, pp := range patterns.patterns {
			upscaledPatternLine := []uint{}
			for _, p := range pp {
				up := s.findTransform(p)
				upscaledPatternLine = append(upscaledPatternLine, up)
			}
			upscaledPattern.patterns = append(upscaledPattern.patterns, upscaledPatternLine)
		}

		// Convert patterns back to pixelized image
		upscaledPattern.convertToPixels(pixels)
		fmt.Printf("Iteration %d, image is (%d,%d), pixels set %d \n", iteration, len(pixels.pixels), len(pixels.pixels), pixels.countSetPixels())
		iteration++
	}

	// Final image dimensions
	fmt.Printf("Final image is (%d,%d)\n", len(pixels.pixels), len(pixels.pixels))

	// Count the '1's
	pixelsSet := pixels.countSetPixels()

	return pixelsSet
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
