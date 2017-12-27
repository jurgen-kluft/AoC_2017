package day23

import (
	"fmt"
)

func solve() int {
	cnt := 0

	var a, b, c, d, e, f, g, h int

	a = 0
	b = 99
	c = b
	d = 0
	e = 0
	f = 0
	g = 0
	h = 0

	if a != 0 {
		goto a1
	}
	goto a5
a1:
	b *= 100
	b -= -100000
	c = b
	c -= -17000
a5:
m23:
	f = 1
	d = 2
g13:
	e = 2
g8:
	g = d
	g = g * e
	cnt++
	g = g - b
	if g != 0 {
		goto g2
	}
	f = 0
g2:
	e -= -1
	g = e
	g = g - b
	if g != 0 {
		goto g8
	}
	d -= -1
	g = d
	g = g - b
	if g != 0 {
		goto g13
	}
	if f != 0 {
		goto f2
	}
	h++
f2:
	g = b
	g = g - c
	if g != 0 {
		goto g22
	}
	goto end
g22:
	b -= -17
	goto m23

end:
	return cnt
}

func solve2() int {
	var a, b, c, d, f, g, h int

	a = 1
	b = 99
	c = b
	d = 0
	f = 0
	g = 0
	h = 0

	if a != 0 {
		goto a1
	}
	goto a5
a1:
	b *= 100
	b += 100000
	c = b + 17000
a5:
m23:
	f = 1
	d = 2
g13:
	//e = b / d
	//g8:
	// Trying to solve
	if (d * (b / d)) == b {
		f = 0
	}
	//g2:
	d++
	g = d - b
	if g != 0 {
		goto g13
	}
	if f != 0 {
		goto f2
	}
	h++
f2:
	g = b - c
	if g != 0 {
		goto g22
	}
	goto end
g22:
	b += 17
	goto m23

end:
	return h
}

// Run1 is the primary solution
func Run1(day int) {
	var result = solve()
	fmt.Printf("Day%d.1: . : %v \n", day, result)
}

// Run2 is the secondary solution
func Run2(day int) {
	var result = solve2()
	fmt.Printf("Day%d.2: . : %v \n", day, result)
}
