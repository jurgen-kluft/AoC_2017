package day18

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

func readLine(line string, s *Program) {

	ins := &Instruction{}
	result := strings.Split(line, " ")
	result[0] = strings.TrimSpace(result[0])
	opcode := 0
	if result[0] == "set" {
		opcode |= OPCODE_SET
	} else if result[0] == "mul" {
		opcode |= OPCODE_MUL
	} else if result[0] == "jgz" {
		opcode |= OPCODE_JGZ
	} else if result[0] == "add" {
		opcode |= OPCODE_ADD
	} else if result[0] == "mod" {
		opcode |= OPCODE_MOD
	} else if result[0] == "rcv" {
		opcode |= OPCODE_RCV
	} else if result[0] == "snd" {
		opcode |= OPCODE_SND
	}

	if len(result) > 1 {
		r := strings.TrimSpace(result[1])
		if r == "a" {
			opcode |= OPCODE_LEFT_REG
			ins.left = int('a' - 'a')
		} else if r == "b" {
			opcode |= OPCODE_LEFT_REG
			ins.left = int('b' - 'a')
		} else if r == "i" {
			opcode |= OPCODE_LEFT_REG
			ins.left = int('i' - 'a')
		} else if r == "f" {
			opcode |= OPCODE_LEFT_REG
			ins.left = int('f' - 'a')
		} else if r == "p" {
			opcode |= OPCODE_LEFT_REG
			ins.left = int('p' - 'a')
		} else {
			opcode |= OPCODE_LEFT_IMM
			imm, _ := strconv.ParseInt(r, 10, 32)
			ins.left = int(imm)
		}
	}

	if len(result) > 2 {
		r := strings.TrimSpace(result[2])
		if r == "a" {
			opcode |= OPCODE_RIGHT_REG
			ins.right = int('a' - 'a')
		} else if r == "b" {
			opcode |= OPCODE_RIGHT_REG
			ins.right = int('b' - 'a')
		} else if r == "i" {
			opcode |= OPCODE_RIGHT_REG
			ins.right = int('i' - 'a')
		} else if r == "f" {
			opcode |= OPCODE_RIGHT_REG
			ins.right = int('f' - 'a')
		} else if r == "p" {
			opcode |= OPCODE_RIGHT_REG
			ins.right = int('p' - 'a')
		} else {
			opcode |= OPCODE_RIGHT_IMM
			imm, _ := strconv.ParseInt(r, 10, 32)
			ins.right = int(imm)
		}
	}

	ins.opcode = opcode

	s.program = append(s.program, ins)

	return
}

func read(filename string) *Program {
	s := &Program{}
	s.program = []*Instruction{}
	s.reg = make([]int, 32, 32)

	reader := func(line string) {
		readLine(line, s)
	}
	iterateOverLinesInTextFile(filename, reader)

	return s
}

type queue struct {
	data  []int
	count int
}

func (q *queue) init() {
	q.data = []int{}
	q.count = 0
}

func (q *queue) pop() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, true
}

func (q *queue) push(value int) {
	q.data = append(q.data, value)
	q.count++
}

// Program iiii
type Program struct {
	program []*Instruction
	reg     []int
	PC      int
	sndq    *queue
	rcvq    *queue
}

func (s *Program) copy() *Program {
	clone := &Program{}
	clone.program = s.program
	clone.program = make([]*Instruction, len(s.program), len(s.program))
	for i, ins := range s.program {
		clone.program[i] = ins
	}
	clone.reg = make([]int, 32, 32)
	clone.PC = 0
	return clone
}

func (i *Instruction) print() {
	opcode := ""
	switch i.opcode & 0xFF {
	case OPCODE_SET:
		opcode = "set"
		break
	case OPCODE_MUL:
		opcode = "mul"
		break
	case OPCODE_ADD:
		opcode = "add"
		break
	case OPCODE_MOD:
		opcode = "mod"
		break
	case OPCODE_RCV:
		opcode = "rcv"
		break
	case OPCODE_SND:
		opcode = "snd"
		break
	case OPCODE_JGZ:
		opcode = "jgz"
		break
	}
	left := ""
	if flagIs(i.opcode, OPCODE_LEFT_REG) {
		left = fmt.Sprintf("%s", string(rune('a'+i.left)))
	} else if flagIs(i.opcode, OPCODE_LEFT_IMM) {
		left = fmt.Sprintf("%d", i.left)
	}
	right := ""
	if flagIs(i.opcode, OPCODE_RIGHT_REG) {
		right = fmt.Sprintf("%s", string(rune('a'+i.right)))
	} else if flagIs(i.opcode, OPCODE_RIGHT_IMM) {
		right = fmt.Sprintf("%d", i.right)
	}

	fmt.Printf("%s %s %s\n", opcode, left, right)
}

func (s *Program) print() {
	for _, i := range s.program {
		i.print()
	}
}

func flagIs(flags int, flag int) bool {
	return (flags & flag) == flag
}

func (s *Program) set(opcode int, left int, right int) {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		if flagIs(opcode, OPCODE_RIGHT_REG) {
			s.reg[left] = s.reg[right]
		} else {
			s.reg[left] = right
		}
	}
}
func (s *Program) mul(opcode int, left int, right int) {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		if flagIs(opcode, OPCODE_RIGHT_REG) {
			s.reg[left] = s.reg[left] * s.reg[right]
		} else {
			s.reg[left] = s.reg[left] * right
		}
	}
}
func (s *Program) add(opcode int, left int, right int) {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		if flagIs(opcode, OPCODE_RIGHT_REG) {
			s.reg[left] = s.reg[left] + s.reg[right]
		} else {
			s.reg[left] = s.reg[left] + right
		}
	}
}
func (s *Program) mod(opcode int, left int, right int) {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		if flagIs(opcode, OPCODE_RIGHT_REG) {
			s.reg[left] = s.reg[left] % s.reg[right]
		} else {
			s.reg[left] = s.reg[left] % right
		}
	}
}

func (s *Program) jgz(opcode int, left int, right int) int {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		if s.reg[left] > 0 {
			if flagIs(opcode, OPCODE_RIGHT_REG) {
				return s.reg[right]
			}
			return right
		}
	} else {
		if left > 0 {
			if flagIs(opcode, OPCODE_RIGHT_REG) {
				return s.reg[right]
			}
			return right
		}
	}
	return 1
}

func (s *Program) rcv(opcode int, left int, right int) bool {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		if s.reg[left] > 0 {
			return true
		}
	}
	return false
}
func (s *Program) snd(opcode int, left int, right int) int {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		return s.reg[left]
	}
	return left
}

func (s *Program) rcv2(opcode int, left int, right int) bool {
	value, exists := s.rcvq.pop()
	if exists {
		s.reg[left] = value
	}
	return exists
}

func (s *Program) snd2(opcode int, left int, right int) {
	if flagIs(opcode, OPCODE_LEFT_REG) {
		s.sndq.push(s.reg[left])
	} else {
		s.sndq.push(left)
	}
}

func (s *Program) execute() int {
	freq := 0
	recover := 0

	s.PC = 0
	for s.PC >= 0 && s.PC < len(s.program) {
		ins := s.program[s.PC]
		//fmt.Printf("Line %d - ", s.PC)
		//ins.print()
		switch ins.opcode & 0xFF {
		case OPCODE_SET:
			s.set(ins.opcode, ins.left, ins.right)
			s.PC++
			break
		case OPCODE_MUL:
			s.mul(ins.opcode, ins.left, ins.right)
			s.PC++
			break
		case OPCODE_ADD:
			s.add(ins.opcode, ins.left, ins.right)
			s.PC++
			break
		case OPCODE_MOD:
			s.mod(ins.opcode, ins.left, ins.right)
			s.PC++
			break
		case OPCODE_RCV:
			if s.rcv(ins.opcode, ins.left, ins.right) {
				recover = freq
				s.PC = len(s.program)
			}
			s.PC++
			break
		case OPCODE_SND:
			freq = s.snd(ins.opcode, ins.left, ins.right)
			s.PC++
			break
		case OPCODE_JGZ:
			s.PC += s.jgz(ins.opcode, ins.left, ins.right)
			break
		default:
			fmt.Println("Error, unknown instruction")
			break
		}
	}
	return recover
}

func (s *Program) step() (int, bool) {
	if s.PC >= 0 && s.PC < len(s.program) {
		ins := s.program[s.PC]
		//fmt.Printf("Line %d - ", s.PC)
		//ins.print()
		PCA := 1
		switch ins.opcode & 0xFF {
		case OPCODE_SET:
			s.set(ins.opcode, ins.left, ins.right)
			break
		case OPCODE_MUL:
			s.mul(ins.opcode, ins.left, ins.right)
			break
		case OPCODE_ADD:
			s.add(ins.opcode, ins.left, ins.right)
			break
		case OPCODE_MOD:
			s.mod(ins.opcode, ins.left, ins.right)
			break
		case OPCODE_RCV:
			if s.rcv2(ins.opcode, ins.left, ins.right) {
				PCA = 1
			} else {
				PCA = 0
			}
			break
		case OPCODE_SND:
			s.snd2(ins.opcode, ins.left, ins.right)
			break
		case OPCODE_JGZ:
			PCA = s.jgz(ins.opcode, ins.left, ins.right)
			break
		default:
			fmt.Println("Error, unknown instruction")
			PCA = 0
			break
		}
		s.PC += PCA
		return PCA, true
	}
	return 0, false

}

const (
	OPCODE_SET       = 0x001
	OPCODE_MUL       = 0x002
	OPCODE_JGZ       = 0x004
	OPCODE_ADD       = 0x008
	OPCODE_MOD       = 0x010
	OPCODE_RCV       = 0x020
	OPCODE_SND       = 0x040
	OPCODE_LEFT_IMM  = 0x100
	OPCODE_LEFT_REG  = 0x200
	OPCODE_RIGHT_IMM = 0x400
	OPCODE_RIGHT_REG = 0x800
)

// Instruction contains an OPCODE and LEFT RIGHT target
type Instruction struct {
	opcode int
	left   int
	right  int
}

func (s *Program) solve() int {
	s.print()
	return s.execute()
}

func (s *Program) solve2() int {
	p0 := s.copy()
	p1 := s.copy()
	q1 := &queue{}
	q2 := &queue{}
	q1.init()
	q2.init()
	p0.sndq = q1
	p0.rcvq = q2
	p1.sndq = q2
	p1.rcvq = q1

	p0.reg[int('p'-'a')] = 1
	p1.reg[int('p'-'a')] = 0

	P0PCA := 0
	P1PCA := 0
	p0IsRunning := true
	p1IsRunning := true
	for p0IsRunning && p1IsRunning {
		P0PCA, p0IsRunning = p0.step()
		P1PCA, p1IsRunning = p1.step()
		if P0PCA == 0 && P1PCA == 0 {
			break
		}
	}
	return q1.count
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
	fmt.Printf("Day%d.2: Program 1 send count : %v \n", day, result)
}
