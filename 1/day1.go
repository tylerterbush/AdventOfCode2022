package main

import (
	"AdventOfCode2022/common"
	"log"
	"sort"
	"strconv"
)

func partOne() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	curTotal := 0
	maxTotal := 0
	for _, line := range lines {
		if line == "" {
			if curTotal > maxTotal {
				maxTotal = curTotal
			}
			curTotal = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		common.FatalIf(err)
		curTotal += calories
	}

	// We can reach end of input and need to check calories then
	if curTotal > maxTotal {
		maxTotal = curTotal
	}

	log.Println("Part 1 - max calories for one elf is:", maxTotal)
}

func partTwo() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	curTotal := 0
	totals := []int{}
	for _, line := range lines {
		if line == "" {
			totals = append(totals, curTotal)
			curTotal = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		common.FatalIf(err)
		curTotal += calories
	}

	// We can reach end of input and need to check calories then
	if curTotal > 0 {
		totals = append(totals, curTotal)
	}

	sort.Ints(totals)
	top3 := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]
	log.Println("Part 2 - calories total of top 3 elves:", top3)
}

func main() {
	partOne()
	partTwo()
}
