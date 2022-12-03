package main

import (
	"AdventOfCode2022/common"
	"log"
	"strings"
)

// Follow the exact strat in the guide and see what your score is
func partOne() {
	// A = Rock, B = Paper, C = Scissors
	// X = Rock, Y = Paper, Z = Scissors
	// Win = 6, Tie = 3, Loss = 0
	var scoreMap = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	var winMap = map[string]map[string]int{
		"A": map[string]int{
			"X": 3,
			"Y": 6,
			"Z": 0,
		},
		"B": map[string]int{
			"X": 0,
			"Y": 3,
			"Z": 6,
		},
		"C": map[string]int{
			"X": 6,
			"Y": 0,
			"Z": 3,
		},
	}

	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	totalScore := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		totalScore += scoreMap[split[1]] + winMap[split[0]][split[1]]
	}

	log.Println("Part 1 - total score:", totalScore)
}

// Now X means you need to lose, Y means you need to draw, Z means you need to win
func partTwo() {
	var chooseMap = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	var chooseList = []int{1, 2, 3}
	var winMap = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	totalScore := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		outcome := split[1]
		totalScore += winMap[outcome] // we already know if we're tying, winning or losing

		val := split[0]
		index := chooseMap[val] - 1
		if outcome == "X" { // Lose
			newIndex := (index - 1 + 3) % 3
			totalScore += chooseList[newIndex]
		} else if outcome == "Y" { // Draw
			totalScore += chooseMap[split[0]]
		} else if outcome == "Z" { // Win
			newIndex := (index + 1 + 3) % 3
			totalScore += chooseList[newIndex]
		}
	}

	log.Println("Part 2 - total score:", totalScore)

}

func main() {
	partOne()
	partTwo()
}
