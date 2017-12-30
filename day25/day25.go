package day25

import (
	"fmt"
	//"strconv"
	//"strings"
)

// Tape holds the array of '0's and '1's
type Tape struct {
	tape   map[int]int
	cursor int
	ones   int
}

func (t *Tape) get() int {
	value, exists := t.tape[t.cursor]
	if !exists {
		value = 0
		t.tape[t.cursor] = value
	}
	return value
}

func (t *Tape) set(newValue int) {
	oldValue := t.get()
	if oldValue == 0 && newValue == 1 {
		t.ones++
	} else if oldValue == 1 && newValue == 0 {
		t.ones--
	}
	t.tape[t.cursor] = newValue
}

func (t *Tape) move(dir int) {
	t.cursor += dir
}

// States holds state 'A' to 'F'
type States struct {
	A State
	B State
	C State
	D State
	E State
	F State
}

// Solution is the hex-grid
type Solution struct {
	tape   *Tape
	states *States
	state  State
}

// stateA executes state A logic
type stateA struct {
	states *States
}

func (s *stateA) execute(tape *Tape) State {
	var next State
	c := tape.get()
	if c == 0 {
		tape.set(1)
		tape.move(1)
		next = s.states.B
	} else {
		tape.set(0)
		tape.move(-1)
		next = s.states.B
	}
	return next
}

type stateB struct {
	states *States
}

func (s *stateB) execute(tape *Tape) State {
	var next State
	c := tape.get()
	if c == 0 {
		tape.set(1)
		tape.move(-1)
		next = s.states.C
	} else {
		tape.set(0)
		tape.move(1)
		next = s.states.E
	}
	return next
}

type stateC struct {
	states *States
}

func (s *stateC) execute(tape *Tape) State {
	var next State
	c := tape.get()
	if c == 0 {
		tape.set(1)
		tape.move(1)
		next = s.states.E
	} else {
		tape.set(0)
		tape.move(-1)
		next = s.states.D
	}
	return next
}

type stateD struct {
	states *States
}

func (s *stateD) execute(tape *Tape) State {
	var next State
	c := tape.get()
	if c == 0 {
		tape.set(1)
		tape.move(-1)
		next = s.states.A
	} else {
		tape.set(1)
		tape.move(-1)
		next = s.states.A
	}
	return next
}

type stateE struct {
	states *States
}

func (s *stateE) execute(tape *Tape) State {
	var next State
	c := tape.get()
	if c == 0 {
		tape.set(0)
		tape.move(1)
		next = s.states.A
	} else {
		tape.set(0)
		tape.move(1)
		next = s.states.F
	}
	return next
}

type stateF struct {
	states *States
}

func (s *stateF) execute(tape *Tape) State {
	var next State
	c := tape.get()
	if c == 0 {
		tape.set(1)
		tape.move(1)
		next = s.states.E
	} else {
		tape.set(1)
		tape.move(1)
		next = s.states.A
	}
	return next
}

// State is an execution scope with it's own logic manipulating a tape
type State interface {
	execute(tape *Tape) State
}

func (s *Solution) solve() int {
	s.tape = &Tape{tape: map[int]int{}, cursor: 0}
	s.states = &States{}
	s.states.A = &stateA{states: s.states}
	s.states.B = &stateB{states: s.states}
	s.states.C = &stateC{states: s.states}
	s.states.D = &stateD{states: s.states}
	s.states.E = &stateE{states: s.states}
	s.states.F = &stateF{states: s.states}
	s.state = s.states.A
	steps := 12861455
	step := 0
	for step < steps {
		s.state = s.state.execute(s.tape)
		step++
	}

	return s.tape.ones
}

func (s *Solution) solve2() int {
	return 0
}

// Run1 is the primary solution
func Run1(day int) {
	s := &Solution{}
	var result = s.solve()
	fmt.Printf("Day%d.1: . : %v \n", day, result)
}

// Run2 is the secondary solution
func Run2(day int) {
	s := &Solution{}
	var result = s.solve2()
	fmt.Printf("Day%d.2: . : %v \n", day, result)
}
