package main

import "fmt"

import "github.com/jurgen-kluft/AoC_2017/download"

import "github.com/jurgen-kluft/AoC_2017/day1"
import "github.com/jurgen-kluft/AoC_2017/day2"
import "github.com/jurgen-kluft/AoC_2017/day3"
import "github.com/jurgen-kluft/AoC_2017/day4"
import "github.com/jurgen-kluft/AoC_2017/day5"
import "github.com/jurgen-kluft/AoC_2017/day6"
import "github.com/jurgen-kluft/AoC_2017/day7"
import "github.com/jurgen-kluft/AoC_2017/day8"
import "github.com/jurgen-kluft/AoC_2017/day9"
import "github.com/jurgen-kluft/AoC_2017/day10"
import "github.com/jurgen-kluft/AoC_2017/day11"
import "github.com/jurgen-kluft/AoC_2017/day12"
import "github.com/jurgen-kluft/AoC_2017/day13"
import "github.com/jurgen-kluft/AoC_2017/day14"
import "github.com/jurgen-kluft/AoC_2017/day15"
import "github.com/jurgen-kluft/AoC_2017/day16"
import "github.com/jurgen-kluft/AoC_2017/day17"
import "github.com/jurgen-kluft/AoC_2017/day18"
import "github.com/jurgen-kluft/AoC_2017/day19"
import "github.com/jurgen-kluft/AoC_2017/day20"
import "github.com/jurgen-kluft/AoC_2017/day21"
import "github.com/jurgen-kluft/AoC_2017/day22"
import "github.com/jurgen-kluft/AoC_2017/day23"
import "github.com/jurgen-kluft/AoC_2017/day24"
import "github.com/jurgen-kluft/AoC_2017/day25"

func day(day int) {
	download.GetInputForDay(day)

	switch day {
	case 1:
		day1.Run1()
		day1.Run2()
		break
	case 2:
		day2.Run1()
		day2.Run2()
		break
	case 3:
		day3.Run1()
		day3.Run2()
		break
	case 4:
		day4.Run1()
		day4.Run2()
		break
	case 5:
		day5.Run1()
		day5.Run2()
		break
	case 6:
		day6.Run1()
		day6.Run2()
		break
	case 7:
		day7.Run1()
		day7.Run2()
		break
	case 8:
		day8.Run1()
		day8.Run2()
		break
	case 9:
		day9.Run1()
		day9.Run2()
		break
	case 10:
		day10.Run1()
		day10.Run2()
		break
	case 11:
		day11.Run1()
		day11.Run2()
		break
	case 12:
		day12.Run1()
		day12.Run2()
		break
	case 13:
		day13.Run1()
		day13.Run2()
		break
	case 14:
		day14.Run1()
		day14.Run2()
		break
	case 15:
		day15.Run1()
		day15.Run2()
		break
	case 16:
		day16.Run1(day)
		day16.Run2(day)
		break
	case 17:
		day17.Run1(day)
		day17.Run2(day)
		break
	case 18:
		day18.Run1(day)
		day18.Run2(day)
		break
	case 19:
		day19.Run1(day)
		day19.Run2(day)
		break
	case 20:
		day20.Run1(day)
		day20.Run2(day)
		break
	case 21:
		day21.Run1(day)
		day21.Run2(day)
		break
	case 22:
		day22.Run1(day)
		day22.Run2(day)
		break
	case 23:
		day23.Run1(day)
		day23.Run2(day)
		break
	case 24:
		day24.Run1(day)
		day24.Run2(day)
		break
	case 25:
		day25.Run1(day)
		day25.Run2(day)
		break
	}
}

func main() {
	fmt.Printf("Advent of Code - 2017.\n")
	day(25)
}
