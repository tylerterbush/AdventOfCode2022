package main

import (
	"log"
	"os"
	"sort"
)

// Each item is the worry level of that item
//

type Monkey struct {
	Items []int
	Operation string
	OperationNum int
	TestDivisibleBy int
	IfTrue int
	IfFalse int
	NumberOfInspections int
}

func getMonkeyList() []*Monkey {
	return []*Monkey {
		{ // 0
			Items: []int{98, 70, 75, 80, 84, 89, 55, 98},
			Operation: "multiply",
			OperationNum: 2,
			TestDivisibleBy: 11,
			IfTrue: 1,
			IfFalse: 4,
			NumberOfInspections: 0,
		},
		{ // 1
			Items: []int{59},
			Operation: "square",
			TestDivisibleBy: 19,
			IfTrue: 7,
			IfFalse: 3,
			NumberOfInspections: 0,
		},
		{ // 2
			Items: []int{77, 95, 54, 65, 89},
			Operation: "add",
			OperationNum: 6,
			TestDivisibleBy: 7,
			IfTrue: 0,
			IfFalse: 5,
			NumberOfInspections: 0,
		},
		{ // 3
			Items: []int{71, 64, 75},
			Operation: "add",
			OperationNum: 2,
			TestDivisibleBy: 17,
			IfTrue: 6,
			IfFalse: 2,
			NumberOfInspections: 0,
		},
		{ // 4
			Items: []int{74, 55, 87, 98},
			Operation: "multiply",
			OperationNum: 11,
			TestDivisibleBy: 3,
			IfTrue: 1,
			IfFalse: 7,
			NumberOfInspections: 0,
		},
		{ // 5
			Items: []int{90, 98, 85, 52, 91, 60},
			Operation: "add",
			OperationNum: 7,
			TestDivisibleBy: 5,
			IfTrue: 0,
			IfFalse: 4,
			NumberOfInspections: 0,
		},
		{ // 6
			Items: []int{99, 51},
			Operation: "add",
			OperationNum: 1,
			TestDivisibleBy: 13,
			IfTrue: 5,
			IfFalse: 2,
			NumberOfInspections: 0,
		},
		{ // 7
			Items: []int{98, 94, 59, 76, 51, 65, 75},
			Operation: "add",
			OperationNum: 5,
			TestDivisibleBy: 2,
			IfTrue: 3,
			IfFalse: 6,
			NumberOfInspections: 0,
		},
	}
}

// Perform the operation, then divide it by 3 rounding down to the nearest int
// Multiple two most active monkeys' number of inspections together
func partOne() {
	var (
		numRounds = 20
		monkeys = getMonkeyList()
	)

	for i := 0; i < numRounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				// Perform the inspection calculation
				switch monkey.Operation {
				case "multiply":
					item *= monkey.OperationNum
				case "add":
					item += monkey.OperationNum
				case "square":
					item *= item
				default:
					os.Exit(1)
				}

				// Now divide by 3 and round to the nearest int
				item /= 3

				// Perform the divisible check to see where to throw it
				// and then move it
				if item % monkey.TestDivisibleBy == 0 {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, item)
				} else {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, item)
				}
				monkey.NumberOfInspections++
			}
			// Clear out this monkey's items. They threw them all
			monkey.Items = []int{}
		}
	}

	// Sort the monkeys by number of inspections
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].NumberOfInspections < monkeys[j].NumberOfInspections
	})

	log.Println("Part One - monkey business:", monkeys[len(monkeys)-1].NumberOfInspections * monkeys[len(monkeys)-2].NumberOfInspections)
}

// We no longer divide our worry levels by 3 after the monkey inspects it
// What's the monkey business level after 10000 rounds?
func partTwo() {
	var (
		numRounds = 10000
		monkeys = getMonkeyList()
		// I asked "What is the smallest number that is divisible by all of the following numbers (11, 19, 7, 17, 3, 5, 13, 2)"
		// to https://chat.openai.com/chat for this :)
		leastCommonMultiple = 232792560
	)

	for i := 0; i < numRounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				// Perform the inspection calculation
				switch monkey.Operation {
				case "multiply":
					item *= monkey.OperationNum
				case "add":
					item += monkey.OperationNum
				case "square":
					item *= item
				default:
					os.Exit(1)
				}

				item = item % leastCommonMultiple

				// Perform the divisible check to see where to throw it
				// and then move it
				if item % monkey.TestDivisibleBy == 0 {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, item)
				} else {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, item)
				}
				monkey.NumberOfInspections++
			}
			// Clear out this monkey's items. They threw them all
			monkey.Items = []int{}
		}
	}

	// Sort the monkeys by number of inspections
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].NumberOfInspections < monkeys[j].NumberOfInspections
	})

	log.Println("Part Two - monkey business:", monkeys[len(monkeys)-1].NumberOfInspections * monkeys[len(monkeys)-2].NumberOfInspections)
}

func main() {
	partOne()
	partTwo()
}
