package day24

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
func readLine(line string) (in int, out int) {
	fmt.Sscanf(line, "%d/%d", &in, &out)
	return
}
func read(filename string) *Solution {
	s := &Solution{}

	begin := []*Node{}
	nodes := make([]*Node, 64, 64)
	reader := func(line string) {
		in, out := readLine(line)
		//fmt.Printf("%d/%d\n", in, out)
		p1 := nodes[in]
		if p1 == nil {
			p1 = &Node{}
			p1.out = []*Vector{}
			nodes[in] = p1
		}
		p2 := nodes[out]
		if p2 == nil {
			p2 = &Node{}
			p2.out = []*Vector{}
			nodes[out] = p2
		}
		if in == 0 {
			begin = append(begin, p1)
		}
		if out == 0 {
			begin = append(begin, p2)
		}
		v := &Vector{src: p1, dst: p2, in: in, out: out}
		p1.connectTo(v)
		p2.connectTo(v)
	}
	iterateOverLinesInTextFile(filename, reader)

	s.graph = nodes
	s.begin = begin
	return s
}

// Solution is the hex-grid
type Solution struct {
	graph []*Node
	begin []*Node
}

// Vector is a path/connection between 2 nodes
type Vector struct {
	src     *Node
	dst     *Node
	in, out int
	visited bool
}

// PathFinder will remember a specific path
type PathFinder interface {
	remember(p *Path)
}

func (v *Vector) traverse(n *Node, p *Path, pf PathFinder) {
	if n == v.src {
		v.dst.traverse(p, pf)
	} else if n == v.dst {
		v.src.traverse(p, pf)
	}
}

func (v *Vector) visit() bool {
	if v.visited == false {
		v.visited = true
		return true
	}
	return false
}
func (v *Vector) leave() {
	v.visited = false
}

// Node is a graph node
type Node struct {
	out []*Vector // out-going vectors
}

// Path is a graph traversal
type Path struct {
	path []*Vector
	len  int
}

func (p *Path) push(v *Vector) {
	p.path = append(p.path[:p.len], v)
	p.len++
}

func (p *Path) pop() *Vector {
	p.len--
	v := p.path[p.len]
	p.path = p.path[:p.len]
	return v
}

// StrongestPathFinder will look for the strongest path no matter the length
type StrongestPathFinder struct {
	score int
}

func (pf *StrongestPathFinder) remember(p *Path) {
	score := 0
	for _, v := range p.path {
		score += v.in + v.out
	}
	if score > pf.score {
		fmt.Printf("Path with score(%d), length(%d)\n", score, p.len)
		pf.score = score
	}
}

// StrongestLongestPathFinder will look for the strongest but longest path
type StrongestLongestPathFinder struct {
	score  int
	length int
}

func (pf *StrongestLongestPathFinder) remember(p *Path) {
	score := 0
	for _, v := range p.path {
		score += v.in + v.out
	}
	length := p.len
	if length > pf.length {
		fmt.Printf("Path with larger length(%d) and score(%d)\n", length, score)
		pf.score = score
		pf.length = length
	} else if score > pf.score && length == pf.length {
		fmt.Printf("Path with same length(%d) and bigger score(%d)\n", length, score)
		pf.score = score
		pf.length = length
	}
}

func (n *Node) connectTo(v *Vector) {
	n.out = append(n.out, v)
}

func (n *Node) traverse(p *Path, pf PathFinder) {
	pf.remember(p)
	for _, v := range n.out {
		if v.visit() {
			p.push(v)
			v.traverse(n, p, pf)
			v.leave()
			p.pop()
		}
	}
}

func (s *Solution) solve() int {
	p := &Path{}
	pf := &StrongestPathFinder{}
	fmt.Printf("Graph has %d start nodes - searching strongest \n", len(s.begin))
	for _, n := range s.begin {
		n.traverse(p, pf)
	}
	return pf.score
}

func (s *Solution) solve2() int {
	p := &Path{}
	pf := &StrongestLongestPathFinder{}
	fmt.Printf("Graph has %d start nodes - searching longest and strongest \n", len(s.begin))
	for _, n := range s.begin {
		n.traverse(p, pf)
	}
	return pf.score
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
