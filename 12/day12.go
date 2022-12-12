package main

import (
	"AdventOfCode2022/common"
	"log"
	"math"
)
/*
`a` is lowest elevation up to `z` the highest
`S` is starting position and `E` is the spot with the best signal
You want to reach `E` in as few steps as possible
You can only move horizonally or vertically adjacent 1 square
Can only move 1 elevation higher
I guess this means you can also move to equal elevation or down any number of levels
What is the fewest steps required to reach E from S
Could just do BFS... maybe start from the E and branch out (can only move to spots that are -1 or greater)
This would find shortest path from each space to the E
 */

type QueueUnit struct {
	Row int
	Col int
	Distance int
	LastHeight int
}

type Position struct {
	Row int
	Col int
}

func buildGrids() ([][]int, [][]int, Position, Position){
	var (
		elevations [][]int
		distances [][]int
		startPos Position
		endPos Position
	)

	lines, err := common.GetLinesFromFile("12/input.txt")
	common.FatalIf(err)

	for r, line := range lines {
		var elevationRow []int
		var distanceRow []int

		for c, char := range line {
			if char == 'S' {
				startPos = Position{ Row: r, Col: c}
				elevationRow = append(elevationRow, 'a')
			} else if char == 'E' {
				endPos = Position{ Row: r, Col: c }
				elevationRow = append(elevationRow, 'z')
			} else {
				elevationRow = append(elevationRow, int(char))
			}
			distanceRow = append(distanceRow, -1)
		}
		elevations = append(elevations, elevationRow)
		distances = append(distances, distanceRow)
	}

	return elevations, distances, startPos, endPos
}


var (
	elevations, distances, startPos, signalPos = buildGrids()
)

func helperOneBFS() {
	// Make a queue and add the signal point to it
	queue := common.NewQueue()
	queue.Push(QueueUnit{
		Row: signalPos.Row,
		Col: signalPos.Col,
		Distance: 0,
	})

	for !queue.Empty() {
		front := queue.Front().(QueueUnit)

		// Up
		if nextSpaceCheck(front, front.Row - 1, front.Col) {
			queue.Push(QueueUnit{
				Row: front.Row - 1,
				Col: front.Col,
				Distance: front.Distance + 1,
			})
			distances[front.Row-1][front.Col] = front.Distance+1
		}

		// Down
		if nextSpaceCheck(front, front.Row + 1, front.Col) {
			queue.Push(QueueUnit{
				Row: front.Row + 1,
				Col: front.Col,
				Distance: front.Distance + 1,
			})
			distances[front.Row+1][front.Col] = front.Distance+1
		}

		// Left
		if nextSpaceCheck(front, front.Row, front.Col - 1) {
			queue.Push(QueueUnit{
				Row: front.Row,
				Col: front.Col - 1,
				Distance: front.Distance + 1,
			})
			distances[front.Row][front.Col-1] = front.Distance+1
		}

		// Right
		if nextSpaceCheck(front, front.Row, front.Col + 1) {
			queue.Push(QueueUnit{
				Row: front.Row,
				Col: front.Col + 1,
				Distance: front.Distance + 1,
			})
			distances[front.Row][front.Col+1] = front.Distance+1
		}
	}
}

func nextSpaceCheck(front QueueUnit, newRow int, newCol int) bool {
	// if we're out of the grid exit
	if newRow < 0 || newCol < 0 || newRow >= len(distances) || newCol >= len(distances[0]) { return false }

	// If this space is not reachable based on height (it's 2 or more elevations lower than the source)
	heightDiff := elevations[front.Row][front.Col] - elevations[newRow][newCol]
	if heightDiff >= 2 {
		return false
	}

	// If we're looking at the signalPos exit
	if newRow == signalPos.Row && newCol == signalPos.Col {
		return false
	}

	// If this space already has a smaller (non -1) distance than ours, exit
	if distances[newRow][newCol] != -1 {
		return false
	}

	return true
}

func partOne() {
	helperOneBFS()
	log.Println("Part One - shortest number of steps:", distances[startPos.Row][startPos.Col])
}

func partTwo() {
	var (
		shortestPathSteps = math.MaxInt64
	)

	// Now loop through the distance matrix and find the smallest one where it's 'a'
	for r, elevationRow := range elevations {
		for c, elevation := range elevationRow {
			if elevation == 'a' && distances[r][c] != -1 && distances[r][c] < shortestPathSteps {
					shortestPathSteps = distances[r][c]
			}
		}
	}

	log.Println("Part Two: shortest path that starts from elevation 'a':", shortestPathSteps)
}

func main() {
	partOne()
	partTwo()
}
