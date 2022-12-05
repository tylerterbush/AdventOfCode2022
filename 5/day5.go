package main

import (
	"AdventOfCode2022/common"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func helper(stacks [][]string, part string, reverse bool) {
	r := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		regexMatchGroups := r.FindStringSubmatch(line)
		amountToMove, _ := strconv.Atoi(regexMatchGroups[1])
		fromStack, _ := strconv.Atoi(regexMatchGroups[2])
		fromStack-- // to make it line up with our index
		toStack, _ := strconv.Atoi(regexMatchGroups[3])
		toStack--

		elementsToMove := stacks[fromStack][len(stacks[fromStack])-amountToMove : len(stacks[fromStack])]
		// now reverse them (if reverse param is true) and then append to the "toStack"
		// and then delete them from the "fromStack"
		if reverse {
			reverseSlice(elementsToMove)
		}
		stacks[toStack] = append(stacks[toStack], elementsToMove...)
		stacks[fromStack] = stacks[fromStack][0 : len(stacks[fromStack])-amountToMove]
	}

	// Now concatenate the top letter of each stack into a string
	outStr := ""
	for _, stack := range stacks {
		outStr = fmt.Sprintf("%s%s", outStr, stack[len(stack)-1])
	}

	log.Printf("Part %s: top of each stack after all moves - %s\n", part, outStr)
}

func partOne() {
	helper([][]string{
		[]string{"W", "M", "L", "F"},
		[]string{"B", "Z", "V", "M", "F"},
		[]string{"H", "V", "R", "S", "L", "Q"},
		[]string{"F", "S", "V", "Q", "P", "M", "T", "J"},
		[]string{"L", "S", "W"},
		[]string{"F", "V", "P", "M", "R", "J", "W"},
		[]string{"J", "Q", "C", "P", "N", "R", "F"},
		[]string{"V", "H", "P", "S", "Z", "W", "R", "B"},
		[]string{"B", "M", "J", "C", "G", "H", "Z", "W"},
	}, "One", true)
}

func partTwo() {
	helper([][]string{
		[]string{"W", "M", "L", "F"},
		[]string{"B", "Z", "V", "M", "F"},
		[]string{"H", "V", "R", "S", "L", "Q"},
		[]string{"F", "S", "V", "Q", "P", "M", "T", "J"},
		[]string{"L", "S", "W"},
		[]string{"F", "V", "P", "M", "R", "J", "W"},
		[]string{"J", "Q", "C", "P", "N", "R", "F"},
		[]string{"V", "H", "P", "S", "Z", "W", "R", "B"},
		[]string{"B", "M", "J", "C", "G", "H", "Z", "W"},
	}, "Two", false)
}

func reverseSlice(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func main() {
	partOne()
	partTwo()
}
