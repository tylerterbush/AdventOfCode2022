package main

import (
	"AdventOfCode2022/common"
	"log"
	"strconv"
	"strings"
)

// In how many assignment pairs does one contain the other
// i.e. 27-27,27-90 - the first range is fully contained within the second
func partOne() {
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	total := 0
	for _, line := range lines {
		split := strings.Split(line, ",")
		range1Strs := strings.Split(split[0], "-")
		range2Strs := strings.Split(split[1], "-")
		range1Ints := make([]int, len(range1Strs))
		range2Ints := make([]int, len(range2Strs))
		for i, r1s := range range1Strs {
			num, _ := strconv.Atoi(r1s)
			range1Ints[i] = num
		}
		for i, r2s := range range2Strs {
			num, _ := strconv.Atoi(r2s)
			range2Ints[i] = num
		}

		// If range 1 is within range 2
		if range1Ints[0] >= range2Ints[0] && range1Ints[1] <= range2Ints[1] {
			total++
		} else if range2Ints[0] >= range1Ints[0] && range2Ints[1] <= range1Ints[1] {
			total++
		}
	}

	log.Println("Part 1 - number of sections that fully contain another section:", total)
}

func partTwo() {
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	total := 0
	for _, line := range lines {
		split := strings.Split(line, ",")
		range1Strs := strings.Split(split[0], "-")
		range2Strs := strings.Split(split[1], "-")
		range1Ints := make([]int, len(range1Strs))
		range2Ints := make([]int, len(range2Strs))
		for i, r1s := range range1Strs {
			num, _ := strconv.Atoi(r1s)
			range1Ints[i] = num
		}
		for i, r2s := range range2Strs {
			num, _ := strconv.Atoi(r2s)
			range2Ints[i] = num
		}

		// Check if the start/end of each range lies within the other range
		if range1Ints[0] >= range2Ints[0] && range1Ints[0] <= range2Ints[1] {
			total++
			continue
		} else if range1Ints[1] >= range2Ints[0] && range1Ints[1] <= range2Ints[1] {
			total++
			continue
		} else if range2Ints[0] >= range1Ints[0] && range2Ints[0] <= range1Ints[1] {
			total++
			continue
		} else if range2Ints[1] >= range1Ints[0] && range2Ints[1] <= range1Ints[1] {
			total++
		}
	}

	log.Println("Part 2 - number of sections that overlap at all:", total)
}

func main() {
	partOne()
	partTwo()
}
