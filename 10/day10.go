package main

import (
	"AdventOfCode2022/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Clock circuit ticks at constant rate
// each tick called a `cycle`
// single register X - starts with value 1
// 2 instructions:
// - `addx V` takes two cycles. after the 2 cylces, X is incremented by the number
// - `noop` takes one cycle and does nothing
// Signal strength:
// cycle number * val of X
// can evaluate every 20 cycles

var (
	partOneCycleNumsToCheck = map[int]struct{}{
		20: struct{}{},
		60: struct{}{},
		100: struct{}{},
		140: struct{}{},
		180: struct{}{},
		220: struct{}{},
	}
)

//Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles. What is the sum of these six signal strengths?
func partOne() {
	var (
		cycle = 0
		X = 1
		signalStrengthSum = 0
	)
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("10/input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		cycle++
		signalStrengthSum += getSignalStrength(cycle, X)

		// Do nothing and take 1 cycle
		if line == "noop" {
			continue
		} else {
			// say in cycle 2
			// during 2 and 3, X doesn't incr
			// after 3 it does so during 4 is when it's added
			split := strings.Split(line, " ")
			num, _ := strconv.Atoi(split[1])

			cycle++
			signalStrengthSum += getSignalStrength(cycle, X)

			X += num // add the number after 2 cycles complete (current one and next one)
		}
	}

	log.Println("Part one - signal strength at benchmarks added:", signalStrengthSum)
}

func getSignalStrength(cycle int, X int) int {
	if _, ok := partOneCycleNumsToCheck[cycle]; ok {
		log.Println("FOUND CYCLE", cycle)
		return cycle * X
	}

	return 0
}

// sprite is 3 pixels wide
// X sets the horizontal pos of the middle of the sprite
// Monitor is 40 wide and 6 high
// The left-most pixel in each row is in position 0, and the right-most pixel in each row is in position 39.
// CRT draws a single pixel during each cycle. From left to right, row by row (0-39 in each row)
func partTwo() {
	_, err := common.GetLinesFromFile("10/input.txt")
	common.FatalIf(err)

	var (
		cycle = 0
		X = 1
		rowPosition = 0
		characterBuffer [][]string
		rowChars []string
	)

	lines, err := common.GetLinesFromFile("10/input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		cycle++

		// Do nothing and take 1 cycle
		if line == "noop" {
			// Draw the pixel
			if rowPosition >= X - 1 && rowPosition <= X + 1 {
				rowChars = append(rowChars, "#")
				rowPosition++
			} else {
				rowChars = append(rowChars, ".")
				rowPosition++
			}
			if len(rowChars) == 40 {
				charsCopy := make([]string, len(rowChars))
				copy(charsCopy, rowChars)
				characterBuffer = append(characterBuffer, charsCopy)
				rowPosition = 0
				rowChars = []string{}
			}
		} else {
			// say in cycle 2
			// during 2 and 3, X doesn't incr
			// after 3 it does so during 4 is when it's added
			split := strings.Split(line, " ")
			num, _ := strconv.Atoi(split[1])

			// Draw the pixel
			if rowPosition >= X - 1 && rowPosition <= X + 1 {
				rowChars = append(rowChars, "#")
				rowPosition++
			} else {
				rowChars = append(rowChars, ".")
				rowPosition++
			}
			if len(rowChars) == 40 {
				charsCopy := make([]string, len(rowChars))
				copy(charsCopy, rowChars)
				characterBuffer = append(characterBuffer, charsCopy)
				rowPosition = 0
				rowChars = []string{}
			}

			cycle++

			//X += num // add the number after 2 cycles complete (current one and next one)
			// Draw the pixel
			if rowPosition >= X - 1 && rowPosition <= X + 1 {
				rowChars = append(rowChars, "#")
				rowPosition++
			} else {
				rowChars = append(rowChars, ".")
				rowPosition++
			}
			if len(rowChars) == 40 {
				charsCopy := make([]string, len(rowChars))
				copy(charsCopy, rowChars)
				characterBuffer = append(characterBuffer, charsCopy)
				rowPosition = 0
				rowChars = []string{}
			}
			X += num
		}
	}

	fmt.Println("Part 2")
	for _, row := range characterBuffer {
		for _, c := range row {
			fmt.Printf(c)
		}
		fmt.Printf("\n")
	}
}

func main() {
	partOne() // 15580 too high
	partTwo()
}
