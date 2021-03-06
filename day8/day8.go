package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	iCMPxL       = 0x0001
	iCMPxLE      = 0x0002
	iCMPxG       = 0x0003
	iCMPxGE      = 0x0004
	iCMPxE       = 0x0005
	iCMPxNE      = 0x0006
	iOPCODExCMPI = 0x10000
	iOPCODExINCI = 0x20000
	iOPCODExDECI = 0x40000
)

// ConditionalInstruction holds the details of
//    [register 1] [INC|DEC OPCODE] [immediate value] if [register 2] [CMP OPCODE] [immediate value]
//
type ConditionalInstruction struct {
	primReg    string
	primOpcode int
	primIVal   int
	condReg    string
	condCmp    int
	condIVal   int
}

func (inst *ConditionalInstruction) print() {
	cmpstr := ""
	switch inst.condCmp {
	case iOPCODExCMPI | iCMPxL:
		cmpstr = "<"
		break
	case iOPCODExCMPI | iCMPxLE:
		cmpstr = "<="
		break
	case iOPCODExCMPI | iCMPxG:
		cmpstr = ">"
		break
	case iOPCODExCMPI | iCMPxGE:
		cmpstr = ">="
		break
	case iOPCODExCMPI | iCMPxE:
		cmpstr = "=="
		break
	case iOPCODExCMPI | iCMPxNE:
		cmpstr = "!="
		break
	}

	primOpcode := ""
	switch inst.primOpcode {
	case iOPCODExDECI:
		primOpcode = "dec"
		break
	case iOPCODExINCI:
		primOpcode = "inc"
		break
	}
	fmt.Printf("%s %s %d ", inst.primReg, primOpcode, inst.primIVal)
	fmt.Printf("if %s %s %d \n", inst.condReg, cmpstr, inst.condIVal)
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

func decodeConditionalInstruction(line string) (inst *ConditionalInstruction, ok bool) {
	inst = &ConditionalInstruction{}

	// example:
	//      fw dec -971 if fz < 1922
	result := strings.Split(line, "if")
	if len(result) != 2 {
		return nil, false
	}

	// Trim leading and trailing white space
	result[0] = strings.TrimSpace(result[0])
	result[1] = strings.TrimSpace(result[1])

	args := strings.Split(result[0], " ")
	if len(args) != 3 {
		return nil, false
	}
	if args[1] == "dec" {
		inst.primOpcode = iOPCODExDECI
	} else if args[1] == "inc" {
		inst.primOpcode = iOPCODExINCI
	} else {
		return nil, false
	}
	s64, err := strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		return nil, false
	}
	inst.primReg = args[0]
	inst.primIVal = int(s64)

	args = strings.Split(result[1], " ")
	if len(args) != 3 {
		return nil, false
	}

	if args[1] == "<" {
		inst.condCmp = iOPCODExCMPI | iCMPxL
	} else if args[1] == ">" {
		inst.condCmp = iOPCODExCMPI | iCMPxG
	} else if args[1] == "<=" {
		inst.condCmp = iOPCODExCMPI | iCMPxLE
	} else if args[1] == ">=" {
		inst.condCmp = iOPCODExCMPI | iCMPxGE
	} else if args[1] == "==" {
		inst.condCmp = iOPCODExCMPI | iCMPxE
	} else if args[1] == "!=" {
		inst.condCmp = iOPCODExCMPI | iCMPxNE
	} else {
		return nil, false
	}

	s64, err = strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		return nil, false
	}
	inst.condReg = args[0]
	inst.condIVal = int(s64)

	return inst, true
}

func readInstructions(filename string) (instructions []*ConditionalInstruction) {
	instructions = []*ConditionalInstruction{}

	reader := func(line string) {
		inst, ok := decodeConditionalInstruction(line)
		if ok {
			instructions = append(instructions, inst)
		}
	}
	iterateOverLinesInTextFile(filename, reader)
	return instructions
}

// VCPU functions as a virtual CPU that can execute instructions and keep track
// of registers and their content
type VCPU struct {
	registers    map[string]int
	highestValue int
}

func (cpu *VCPU) boot() {
	cpu.registers = map[string]int{}
	cpu.registers["a"] = 0
	cpu.highestValue = 0
}

func (cpu *VCPU) getRegister(reg string) int {
	value, exists := cpu.registers[reg]
	if !exists {
		cpu.registers[reg] = 0
		value = 0
	} else {
		if value > cpu.highestValue {
			cpu.highestValue = value
		}
	}
	return value
}

func (cpu *VCPU) setRegister(reg string, value int) {
	cpu.registers[reg] = value
}

func (cpu *VCPU) execute(inst *ConditionalInstruction) {
	condReg := cpu.getRegister(inst.condReg)
	condFlag := false
	switch inst.condCmp {
	case iOPCODExCMPI | iCMPxL:
		condFlag = condReg < inst.condIVal
		break
	case iOPCODExCMPI | iCMPxLE:
		condFlag = condReg <= inst.condIVal
		break
	case iOPCODExCMPI | iCMPxG:
		condFlag = condReg > inst.condIVal
		break
	case iOPCODExCMPI | iCMPxGE:
		condFlag = condReg >= inst.condIVal
		break
	case iOPCODExCMPI | iCMPxE:
		condFlag = condReg == inst.condIVal
		break
	case iOPCODExCMPI | iCMPxNE:
		condFlag = condReg != inst.condIVal
		break
	}
	if condFlag {
		primRegValue := cpu.getRegister(inst.primReg)
		switch inst.primOpcode {
		case iOPCODExDECI:
			primRegValue -= inst.primIVal
			break
		case iOPCODExINCI:
			primRegValue += inst.primIVal
			break
		}
		cpu.setRegister(inst.primReg, primRegValue)
	}
	return
}

func executeInstructions(instructions []*ConditionalInstruction) (largest int, highest int) {
	// Keep track of the content of all our registers
	cpu := &VCPU{}
	cpu.boot()

	for _, instruction := range instructions {
		cpu.execute(instruction)
	}

	maxValue := 0
	for _, value := range cpu.registers {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue, cpu.highestValue
}

func printInstructions(instructions []*ConditionalInstruction) {
	for _, instruction := range instructions {
		instruction.print()
	}
}

// Run1 is the primary solution
func Run1() {
	var instructions = readInstructions("day8/input.text")
	//printInstructions(instructions)
	largestValue, _ := executeInstructions(instructions)
	fmt.Printf("Day 8.1: Largest value in any register: %v \n", largestValue)
}

// Run2 is the secondary solution
func Run2() {
	var instructions = readInstructions("day8/input.text")
	_, highestValue := executeInstructions(instructions)
	fmt.Printf("Day 8.2: Highest value in any register at any time: %v \n", highestValue)
}
