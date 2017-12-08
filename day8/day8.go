package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tower struct {
	refcount int
	weight   int
	name     string
	towers   []string
	children []*Tower
	parent   *Tower
}

// Recursive function to get weight of a tower
func (t *Tower) getWeight() int {
	w := t.weight
	for _, child := range t.children {
		w += child.getWeight()
	}
	return w
}

func (t *Tower) print() {
	if len(t.towers) > 0 {
		fmt.Printf("Tower(%s) : weight(%d) -> ", t.name, t.weight)
		for i, tower := range t.children {
			if i == 0 {
				fmt.Printf("%v(%d)", tower.name, tower.weight)
			} else {
				fmt.Printf(", %v(%d)", tower.name, tower.weight)
			}
		}
	} else {
		fmt.Printf("Tower(%s) : weight(%d)", t.name, t.weight)
	}
	fmt.Printf("\n")
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

func decodeTower(line string) (t *Tower, ok bool) {
	t = &Tower{}
	t.refcount = 0
	t.towers = []string{}
	t.weight = 0
	t.parent = nil

	// parse tower
	// example: ehsqyyb (174) -> xtcdt, tujcuy, wiqohmb, cxdwmu
	result := strings.Split(line, "->")
	fmt.Sscanf(result[0], "%s (%d)", &t.name, &t.weight)

	if len(result) > 1 {
		//fmt.Printf("%s\n", result[1])
		children := strings.Split(result[1], ",")
		for _, child := range children {
			child = strings.TrimSpace(child)
			t.towers = append(t.towers, child)
		}
	}
	return t, true
}

func readTowers(filename string) (towers map[string]*Tower) {
	towers = map[string]*Tower{}

	reader := func(line string) {
		tower, ok := decodeTower(line)
		if ok {
			//tower.print()
			towers[tower.name] = tower
		}
	}
	iterateOverLinesInTextFile(filename, reader)
	return towers
}

func findBottomTower(towers map[string]*Tower) (bottomName string) {
	// Update ref-count and children array in every Tower
	for _, tower := range towers {
		for _, subName := range tower.towers {
			subTower := towers[subName]
			tower.children = append(tower.children, subTower)
			subTower.refcount++
			subTower.parent = tower
		}
	}

	// Find the tower which is not referenced, this one should be
	// the bottom tower
	for _, tower := range towers {
		if tower.refcount == 0 {
			bottomName = tower.name
		}
	}

	for _, tower := range towers {
		subWeights := map[int]int{}
		for _, subName := range tower.towers {
			subTower := towers[subName]
			weight := subTower.getWeight()
			weightCount, hasWeight := subWeights[weight]
			if hasWeight {
				subWeights[weight] = weightCount + 1
			} else {
				subWeights[weight] = 1
			}
		}

		if len(subWeights) == 1 {
			//fmt.Printf("program (%v) is balanced \n", tower.name)
		} else if len(subWeights) == 2 {
			fmt.Printf("Found a program (%v) with one unbalanced weight \n", tower.name)
			tower.print()
			for _, subName := range tower.towers {
				subTower := towers[subName]
				fmt.Printf("%d ", subTower.getWeight())
			}
			fmt.Printf("\n")
		} else if len(subWeights) > 1 {
			fmt.Printf("Found a sub program (%v) with unbalanced weights \n", tower.name)
			tower.print()
			for _, subName := range tower.towers {
				subTower := towers[subName]
				fmt.Printf("%d ", subTower.getWeight())
			}
			fmt.Printf("\n")
		}
	}

	return
}

// Run1 is the primary solution
func Run1() {
	var towers = readTowers("day8/input.text")
	name := findBottomTower(towers)
	fmt.Printf("Day 8.1: Name of bottom tower: %v \n", name)
}

// Run2 is the secondary solution
func Run2() {
	var towers = readTowers("day8/input.text")
	name := findBottomTower(towers)
	fmt.Printf("Day 8.1: Name of bottom tower: %v \n", name)
}
