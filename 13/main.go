package main

import (
	"AdventOfCode2022/common"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"sort"
)

/*
Each line is a list
can contain ints and lists
check if they're in the right order
If two entries are ints, the lower one should be on the left
If both are lists, compare the left-most value, and then the next one if they're equal and so on.
If the left list runs out of items first, they're in the right order
If only one of the vals is a list, convert the integer to a list containing only that value

Pt 1: What is the SUM of all indices where pairs are in the correct order?
 */

func buildInput(filepath string) [][]interface{} {
	var (
		arrays [][]interface{}
	)

	lines, err := common.GetLinesFromFile(filepath)
	common.FatalIf(err)

	for _, line := range lines {
		if line == "" { continue }
		jsonArr := make([]interface{}, 0)
		err = json.Unmarshal([]byte(line), &jsonArr)
		common.FatalIf(err)
		arrays = append(arrays, jsonArr)
	}

	return arrays
}

func partOne() {
	var (
		arrays = buildInput("13/input.txt")
		index = 1
		validIndexes []int
	)

	for i := 0; i <= len(arrays) - 2; i += 2 {
		helper := listComparisonHelper(arrays[i], arrays[i+1])
		if helper == 1 {
			validIndexes = append(validIndexes, index)
		}
		index++
	}

	total := 0
	for _, v := range validIndexes {
		total += v
	}

	log.Println("Part one - total:", total)
}

// Returns whether the lists are in the right order
func listComparisonHelper(listOne []interface{}, listTwo []interface{}) int {
	// If both lists are empty, just continue. Can't make a decision yet
	if len(listOne) == 0 && len(listTwo) == 0 {
		return 0
	}

	// Base case, if list one is empty and list two isn't then they're in the right order
	// If they're both 0 then we can assume it's ok
	if len(listOne) == 0 { return 1 }
	// If list one isn't empty but list two is, then these are not in the right order
	if len(listTwo) == 0 { return -1 }

	loopNum := math.Max(float64(len(listOne)), float64(len(listTwo)))

	// Compare each element of both lists
	// If they can both be converted to ints, compare them
	// If at least one is a list, call recursively. May need to convert one to a list
	for i := 0; i < int(loopNum); i++ {
		// Another base case, if we're at an index in list one that isn't in two, we've
		// gone far enough and can assume they're not in the right order
		if i >= len(listTwo) { return -1 }
		if i >= len(listOne) { return 1 }

		oneInt, okOne := listOne[i].(float64)
		twoInt, okTwo := listTwo[i].(float64)

		// Both are integers. Can just compare them
		// If the int in list one is larger, then these are not in the right order
		if okOne && okTwo {
			if oneInt == twoInt {
				continue
			} else if oneInt > twoInt {
				return -1
			} else {
				return 1
			}
		} else if !okOne && !okTwo { // both are lists, call recursively
			ret := listComparisonHelper(listOne[i].([]interface{}), listTwo[i].([]interface{}))
			if ret == 0 {continue}
			return ret
		} else if !okOne { // Only two is an integer
			ret := listComparisonHelper(listOne[i].([]interface{}), []interface{}{twoInt})
			if ret == 0 {continue}
			return ret
		} else if !okTwo { // only one is an integer
			ret := listComparisonHelper([]interface{}{oneInt}, listTwo[i].([]interface{}))
			if ret == 0 {continue}
			return ret
		}
	}

	// Lists are exactly the same
	return 0
}

func partTwo() {
	var (
		arrays = buildInput("13/input2.txt")
	)

	sort.Slice(arrays, func(i int, j int) bool {
		return listComparisonHelper(arrays[i], arrays[j]) == 1
	})

	var (
		bufferOne = "[[2]]"
		bufferTwo = "[[6]]"

		bufferOneIdx int
		bufferTwoIdx int
	)

	for i, arr := range arrays {
		arrStr := fmt.Sprintf("%+v", arr)
		if arrStr == bufferOne {
			bufferOneIdx = i+1
		} else if arrStr == bufferTwo {
			bufferTwoIdx = i+1
		}
	}

	log.Println("Part Two - locations of both buffers after sorting multiplied together:", bufferOneIdx * bufferTwoIdx)
}

func main() {
	partOne()
	partTwo()
}
