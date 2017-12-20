package day20

import (
	"bufio"
	"fmt"
	"math"
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
	// Example line: p=<5528,2008,1661>, v=<-99,-78,-62>, a=<-17,-2,-2>
	p := make([]int, 3, 3)
	v := make([]int, 3, 3)
	a := make([]int, 3, 3)
	fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &p[0], &p[1], &p[2], &v[0], &v[1], &v[2], &a[0], &a[1], &a[2])

	particle := &Particle{}
	particle.pos = make([]int, 3, 3)
	particle.p = p
	particle.v = v
	particle.a = a

	s.particles = append(s.particles, particle)
	return
}

func read(filename string) *Solution {
	s := &Solution{}

	reader := func(line string) {
		readLine(line, s)
	}
	iterateOverLinesInTextFile(filename, reader)

	return s
}

// Position holds a (x,y,z) position
type Position struct {
	x, y, z int
}

// Particle holds a position, velocity and accelaration
type Particle struct {
	pos    []int
	dist   int
	active bool

	p []int
	v []int
	a []int
}

// Solution holds particles
type Solution struct {
	particles []*Particle
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (s *Solution) solve() int {
	fmt.Printf("Read %d number of particles\n", len(s.particles))

	// Initialize running position of every particle
	for _, p := range s.particles {
		for i, axis := range p.p {
			p.pos[i] = axis
			p.active = true
		}
	}

	tick := 0
	maxDist := math.MinInt32
	minDist := math.MaxInt32
	minParticle := -1
	for tick < 1000 {
		minParticle = -1
		minDist = math.MaxInt32
		for i, p := range s.particles {
			p.v[0] += p.a[0]
			p.v[1] += p.a[1]
			p.v[2] += p.a[2]
			p.p[0] += p.v[0]
			p.p[1] += p.v[1]
			p.p[2] += p.v[2]
			p.dist = abs(p.p[0]) + abs(p.p[1]) + abs(p.p[2])
			if p.dist > maxDist {
				maxDist = p.dist
			}
			if p.dist < minDist {
				minDist = p.dist
				minParticle = i
			}
		}
		tick++
	}

	fmt.Printf("Particle %d is closest to (0,0,0)", minParticle)
	return 0
}

func (s *Solution) solve2() int {
	fmt.Printf("Read %d number of particles\n", len(s.particles))

	// Initialize running position of every particle
	for _, p := range s.particles {
		for i, axis := range p.p {
			p.pos[i] = axis
			p.active = true
		}
	}

	tick := 0
	posmap := map[Position]int{}
	for tick < 1000 {
		for _, p := range s.particles {
			if p.active == true {
				p.v[0] += p.a[0]
				p.v[1] += p.a[1]
				p.v[2] += p.a[2]
				p.p[0] += p.v[0]
				p.p[1] += p.v[1]
				p.p[2] += p.v[2]
				p.dist = abs(p.p[0]) + abs(p.p[1]) + abs(p.p[2])
			}
		}
		for i, p := range s.particles {
			var pos Position
			pos.x = p.p[0]
			pos.y = p.p[1]
			pos.z = p.p[2]
			if ii, ok := posmap[pos]; ok {
				s.particles[i].active = false
				s.particles[ii].active = false
			} else {
				posmap[pos] = i
			}
		}

		tick++
	}

	stillActive := 0
	for _, p := range s.particles {
		if p.active == true {
			stillActive++
		}
	}

	fmt.Printf("%d number of particles are still active", stillActive)
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
