package main

import (
	"AdventOfCode2022/common"
	"log"
	"strconv"
)

func partOne() {
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("8/input.txt")
	common.FatalIf(err)

	// build the grid
	var grid  [][]int
	var visibleGrid [][]bool
	for _, line := range lines {
		row := []int{}
		var visibleRow []bool
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
			visibleRow = append(visibleRow, false)
		}
		grid = append(grid, row)
		visibleGrid = append(visibleGrid, visibleRow)
	}


	// Get total number of visible trees from the outside

	// Check from left to right for each row
	highestSeen := -1
	for i, row := range grid {
		for j, _ := range row {
			gridVal := grid[i][j]
			if gridVal > highestSeen {
				highestSeen = gridVal
				visibleGrid[i][j] = true
			}
			if gridVal == 9 { break }
		}
		highestSeen = -1
	}

	// Check from right to left for each row
	highestSeen = -1
	for i := 0; i < len(grid); i++ {
		for j := len(grid[0])-1; j >= 0; j-- {
			gridVal := grid[i][j]
			if gridVal > highestSeen {
				highestSeen = gridVal
				visibleGrid[i][j] = true
			}
			if gridVal == 9 { break }
		}
		highestSeen = -1
	}

	// Check from top to bottom for each column
	highestSeen = -1
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			gridVal := grid[i][j]
			if gridVal > highestSeen {
				highestSeen = gridVal
				visibleGrid[i][j] = true
			}
			if gridVal == 9 { break }
		}
		highestSeen = -1
	}

	// Check from bottom to top for each column
	highestSeen = -1
	for j := 0; j < len(grid[0]); j++ {
		for i := len(grid)-1; i >= 0; i-- {
			gridVal := grid[i][j]
			if gridVal > highestSeen {
				highestSeen = gridVal
				visibleGrid[i][j] = true
			}
			if gridVal == 9 { break }
		}
		highestSeen = -1
	}

	totalVisible := 0
	for _, row := range visibleGrid {
		for _, val := range row {
			if val {
				totalVisible++
			}
		}
	}

	log.Println("Part One - total number of trees visible:", totalVisible)
}

func partTwo() {
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("8/input.txt")
	common.FatalIf(err)

	// build the grid
	var grid  [][]int
	for _, line := range lines {
		row := []int{}
		var visibleRow []bool
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
			visibleRow = append(visibleRow, false)
		}
		grid = append(grid, row)
	}

	highestScenicScore := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			var (
				upScore int
				downScore int
				leftScore int
				rightScore int
			)

			treeHeight := grid[i][j]

			// upScore
			up := i
			for {
				if up == 0 { break }
				up--
				upScore++
				if grid[up][j] >= treeHeight { break }
			}

			// downScore
			down := i
			for {
				if down == len(grid)-1 { break }
				down++
				downScore++
				if grid[down][j] >= treeHeight { break }
			}

			// leftScore
			left := j
			for {
				if left == 0 { break }
				left--
				leftScore++
				if grid[i][left] >= treeHeight { break }
			}

			// rightScore
			right := j
			for {
				if right == len(grid[0])-1 { break }
				right++
				rightScore++
				if grid[i][right] >= treeHeight { break }
			}

			thisScore := upScore * downScore * leftScore * rightScore
			if thisScore > highestScenicScore { highestScenicScore = thisScore }
		}
	}
	log.Println("Part 2 - highest scenic score:", highestScenicScore)
}

func main() {
	partOne()
	partTwo()
}
