package main

import (
	"AdventOfCode2022/common"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Head and tail of rope must be touching
// Adjacent (diagonally allowed too) or even overlapping
// If the head is not in the same row or col as the tail, the tail always moves diagonally toward the head
// How many spots does the tail visit?
// Answer = 5619
func partOne() {
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("9/input.txt")
	common.FatalIf(err)

	var (
		headR, headC, tailR, tailC int
		tailVisitedSpots = map[string]struct{}{} // will store a string of "X Y" of the tail to show it's been there
	)

	for _, line := range lines {
		split := strings.Split(line, " ")
		dist, _ := strconv.Atoi(split[1])

		// For each distance num
		// - move the head in the direction
		// - follow the tail along
		// - add new tail pos to map
		for i := 0; i < dist; i++{
			switch split[0] {
			case "U":
				headR--
				tailR, tailC = newTailPos(headR, headC, tailR, tailC)
				break
			case "D":
				headR++
				tailR, tailC = newTailPos(headR, headC, tailR, tailC)
				break
			case "L":
				headC--
				tailR, tailC = newTailPos(headR, headC, tailR, tailC)
				break
			case "R":
				headC++
				tailR, tailC = newTailPos(headR, headC, tailR, tailC)
				break
			default:
				os.Exit(1)
			}
			tailVisitedSpots[fmt.Sprintf("%d %d", tailR, tailC)] = struct{}{}
		}
	}
	log.Println("Part One - number of locations the tail visited =", len(tailVisitedSpots))
}

func newTailPos(headR int, headC int, tailR int, tailC int) (int, int) {
	// If they're already next to each other, return tailR, tailC
	if abs(headR - tailR) <= 1 && abs(headC - tailC) <= 1 { return tailR, tailC }

	// Get a new tail position that's next to the head
	if headR - tailR >= 2 {
		tailR++
		tailC = headC
	} else if headR - tailR <= -2 {
		tailR--
		tailC = headC
	}
	if headC - tailC >= 2 {
		tailC++
		tailR = headR
	} else if headC - tailC <= -2 {
		tailC--
		tailR = headR
	}

	return tailR, tailC
}

func newTailPosTwo(headR int, headC int, tailR int, tailC int) (int, int) {
	// If they're already next to each other, return tailR, tailC
	if abs(headR - tailR) <= 1 && abs(headC - tailC) <= 1 { return tailR, tailC }

	// if the row and col are both off by two, move the knot diagonally without setting the
	if (headR - tailR >= 2 || headR - tailR <= -2) && (headC - tailC >= 2 || headC - tailC <= -2) {
		if headR - tailR >= 2 {
			tailR++
		} else {
			tailR--
		}
		if headC - tailC >= 2 {
			tailC++
		} else {
			tailC--
		}

		return tailR, tailC
	}

	// Get a new tail position that's next to the head
	if headR - tailR >= 2 {
		tailR++
		tailC = headC
	} else if headR - tailR <= -2 {
		tailR--
		tailC = headC
	}
	if headC - tailC >= 2 {
		tailC++
		tailR = headR
	} else if headC - tailC <= -2 {
		tailC--
		tailR = headR
	}

	return tailR, tailC
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}

// now the rope has length 10
// how many positions does the tail visit
// 2378 is too high
type Knot struct {
	Row int
	Col int
}
func partTwo() {
	lines, err := common.GetLinesFromFile("9/input.txt")
	common.FatalIf(err)

	var (
		tailVisitedSpots = map[string]struct{}{} // will store a string of "X Y" of the tail to show it's been there
		rope []Knot // The first one in the list will be the head and the last will be the tail
	)

	// Add the knots
	for i := 0; i < 10; i++ {
		rope = append(rope, Knot{Row: 0, Col: 0})
	}

	for _, line := range lines {
		split := strings.Split(line, " ")
		dist, _ := strconv.Atoi(split[1])

		// For each distance num
		// - move the head in the direction
		// - loop over the rest of the knots and move them along
		// - add tail pos to seen map
		for i := 0; i < dist; i++{
			switch split[0] {
			case "U":
				rope[0].Row--
				break
			case "D":
				rope[0].Row++
				break
			case "L":
				rope[0].Col--
				break
			case "R":
				rope[0].Col++
				break
			default:
				os.Exit(1)
			}
			for j := 1; j < len(rope); j++ {
				newRow, newCol := newTailPosTwo(rope[j-1].Row, rope[j-1].Col, rope[j].Row, rope[j].Col)
				rope[j].Row = newRow
				rope[j].Col = newCol
			}
			tailVisitedSpots[fmt.Sprintf("%d %d", rope[len(rope)-1].Row, rope[len(rope) -1].Col)] = struct{}{}
		}
	}

	log.Println("Part Two - number of locations the tail visited =", len(tailVisitedSpots))
}

func main() {
	partOne()
	partTwo()
}
