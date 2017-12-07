package day6

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

func decodeMemoryBanks(line string) (banks []int, ok bool) {
	banks = []int{}
	result := strings.Split(line, "\t")
	for _, bankStr := range result {
		s32, err := strconv.ParseInt(bankStr, 10, 32)
		if err == nil {
			banks = append(banks, int(s32))
		}
	}
	return banks, true
}

func readInstructions(filename string) (banks []int) {
	banks = []int{}

	reader := func(line string) {
		banks, _ = decodeMemoryBanks(line)
	}

	iterateOverLinesInTextFile(filename, reader)
	return banks
}

func chooseBankForRedistribution(banks []int) (bank int) {
	bank = 0
	most := 0
	for i, blocks := range banks {
		if blocks > most {
			most = blocks
			bank = i
		}
	}
	return
}

func balanceMemoryBanks(banks []int) (redistributionCycles int, loopCycles int) {
	redistributionCycles = 0
	loopCycles = 0
	loopConfiguration := ""
	loopCycleIncrement := 0

	sformat := ""
	for i := range banks {
		if i == 0 {
			sformat = "%v"
		} else {
			sformat = sformat + ":%v"
		}
	}

	configurations := map[string]int{}
	for true {
		bank := chooseBankForRedistribution(banks)
		redistributionCycles++
		loopCycles += loopCycleIncrement

		blocks := banks[bank]
		banks[bank] = 0

		bank = (bank + 1) % len(banks)
		for blocks > 0 {
			banks[bank]++
			bank = (bank + 1) % len(banks)
			blocks--
		}
		configuration := ""
		for i, b := range banks {
			if i == 0 {
				configuration = fmt.Sprintf("%v", b)
			} else {
				configuration = configuration + fmt.Sprintf(":%v", b)
			}
		}

		//fmt.Printf("%v\n", configuration)

		_, hasConfig := configurations[configuration]
		if hasConfig {
			configOccurances := configurations[configuration]
			configOccurances++
			configurations[configuration] = configOccurances

			if configOccurances == 2 {
				//fmt.Printf("Have seen this configuration \"%v\" twice!\n", configuration)
				if loopCycles == 0 {
					loopConfiguration = configuration
					loopCycleIncrement = 1
				}
			} else if configOccurances == 3 {
				if configuration == loopConfiguration {
					break
				}
			}
		} else {
			configurations[configuration] = 1
		}
	}
	return
}

// Run1 is the primary solution
func Run1() {
	var program = readInstructions("day6/input.text")
	var cycles, _ = balanceMemoryBanks(program)
	fmt.Printf("Day 6.1: Number of redistribution cycles: %v \n", cycles)
}

// Run2 is the secondary solution
func Run2() {
	var program = readInstructions("day6/input.text")
	var _, cycles = balanceMemoryBanks(program)
	fmt.Printf("Day 6.2: Number of loop cycles: %v \n", cycles)
}
