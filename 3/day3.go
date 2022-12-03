package main

import (
	"AdventOfCode2022/common"
	"log"
	"os"
	"sort"
	"strings"
)

// Returns 1-26 for a-z and 27-52 for A-Z
func numberHelper(r rune) int {
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 1 + 26)
	} else if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}

	return -1
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func partOne() {
	var totalPriority = 0

	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		sack1 := sortString(line[0 : len(line)/2])
		sack2 := sortString(line[len(line)/2 : len(line)])

		sack1Runes := []rune(sack1)
		sack2Runes := []rune(sack2)
		pt1 := 0
		pt2 := 0

		// if the val at one of the pointers is greater, move the other one up then
		// compare
		for {
			// Shouldn't exceed - there must be a match
			if pt1 >= len(sack1Runes) || pt2 >= len(sack2Runes) {
				os.Exit(1)
			}

			if sack1Runes[pt1] > sack2Runes[pt2] {
				pt2++
			} else if sack2Runes[pt2] > sack1Runes[pt1] {
				pt1++
			} else {
				totalPriority += numberHelper(sack1Runes[pt1])
				break
			}
		}
	}

	log.Println("Part One: Total priority:", totalPriority)
}

func commonLetters(a []rune, b []rune) []rune {
	var ret = []rune{}

	pt1 := 0
	pt2 := 0
	for {
		// Shouldn't exceed - there must be a match
		if pt1 >= len(a) || pt2 >= len(b) {
			break
		}

		if a[pt1] > b[pt2] {
			pt2++
			continue
		} else if b[pt2] > a[pt1] {
			pt1++
			continue
		} else {
			// If this isn't a dup char add it to the return list
			if len(ret) == 0 || ret[len(ret)-1] != a[pt1] {
				ret = append(ret, a[pt1])
			}

			// If we just checked the last index of both, exit
			if pt1 == len(a)-1 && pt2 == len(b)-1 {
				break
			}

			// move both pointers if possible
			if pt1 < len(a)-1 {
				pt1++
			}
			if pt2 < len(b)-1 {
				pt2++
			}
			continue
		}
	}

	return ret
}

// Each group of 3 lines is one group
// Find the common item among each 3-line group and add the priorities together
func partTwo() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	totalPriority := 0
	// Get the lines in groups of 3
	for i := 2; i < len(lines); i += 3 {
		threeLines := []string{lines[i-2], lines[i-1], lines[i]}
		commonLettersAB := commonLetters([]rune(sortString(threeLines[0])), []rune(sortString(threeLines[1])))
		commonLettersAC := commonLetters([]rune(sortString(threeLines[0])), []rune(sortString(threeLines[2])))

		for _, ab := range commonLettersAB {
			for _, ac := range commonLettersAC {
				if ab == ac {
					totalPriority += numberHelper(ab)
					break
				}
			}
		}
	}

	log.Println("Part Two - total priority:", totalPriority)
}

func main() {
	partOne()
	partTwo()
}
