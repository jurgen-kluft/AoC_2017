package day4

import (
	"bufio"
	"fmt"
	"github.com/jurgen-kluft/AoC_2017/permutation"
	"os"
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

func validatePassphrase(pp string, perms bool) (ok bool) {
	words := strings.Split(pp, " ")
	exists := map[string]bool{}
	for _, word := range words {
		has := false
		if perms {
			perm, _ := permutation.NewPerm([]rune(word), nil)
			for i, err := perm.Next(); err == nil; i, err = perm.Next() {
				palin := string(i.([]rune))
				if exists[palin] {
					has = true
					break
				}
			}
		} else {
			has = exists[word]
		}
		if !has {
			exists[word] = true
		} else {
			return false
		}
	}
	return true
}

func readValidPassphrases(filename string, perms bool) (numberOfValidPassphrases int) {
	numberOfValidPassphrases = 0

	computator := func(passphrase string) {
		ok := validatePassphrase(passphrase, perms)
		if ok {
			numberOfValidPassphrases++
		}
	}

	iterateOverLinesInTextFile(filename, computator)
	return
}

// Run1 is the primary solution
func Run1() {
	var ok = readValidPassphrases("day4/input.text", false)
	fmt.Printf("Day 4.1: Valid passphrases are %v \n", ok)
}

// Run2 is the secondary solution
func Run2() {
	var ok = readValidPassphrases("day4/input.text", true)
	fmt.Printf("Day 4.2: Valid passphrases are %v \n", ok)
}
