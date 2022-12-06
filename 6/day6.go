package main

import (
	"AdventOfCode2022/common"
	"log"
)

func helper(part string, windowSize int) {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	line := lines[0]
	for i := 0; i < len(line)-(windowSize-1); i++ {
		subStr := line[i : i+windowSize]
		repeat := false
		for j := 0; j < len(subStr)-1; j++ {
			for k := 1; k < len(subStr); k++ {
				if j == k {
					continue
				}
				if subStr[j] == subStr[k] {
					repeat = true
				}
			}
		}

		if !repeat {
			log.Printf("Part %s - number of characters seen to hit a 4-len window of no repeats = %d", part, i+windowSize)
			break
		}
	}
}

func partOne() {
	helper("One", 4)
}

func partTwo() {
	helper("Two", 14)
}

func main() {
	partOne()
	partTwo()
}
