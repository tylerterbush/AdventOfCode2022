package main

import (
	"AdventOfCode2022/common"
	"log"
	"strconv"
	"strings"
)

type Directory struct {
	size            int
	parentDirectory *Directory
	name            string
	subDirectories  []*Directory
}

// For each line
// if it's a cd
// - if it's a directory then make a new subdirecotry
// - if it's .. then change current dir pointer to the parent
// if it's an ls
// - take all files and add to current weight of the directory
// - go up the chain and add that file weight to all parent directories

// NOTES
// 1019191 - first guess was wrong
func partOne() {
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("7/input.txt")
	common.FatalIf(err)

	var sizeMap = map[string]int{}

	var root = Directory{
		name:            "root",
		size:            0,
		subDirectories:  []*Directory{},
		parentDirectory: nil,
	}

	curDirectory := &root
	for _, line := range lines {
		if line == "$ cd /" {
			continue
		}

		split := strings.Split(line, " ")
		// cd command
		if split[0] == "$" && split[1] == "cd" {
			if split[2] == ".." {
				// go up a directory to one we've seen already
				curDirectory = curDirectory.parentDirectory
			} else {
				// Go down a level into a new directory (assuming no repeats in puzzle input)
				newChildDirectory := Directory{
					name:            curDirectory.name + "-" + split[2],
					size:            0,
					subDirectories:  []*Directory{},
					parentDirectory: curDirectory,
				}
				curDirectory.subDirectories = append(curDirectory.subDirectories, &newChildDirectory)
				curDirectory = &newChildDirectory
			}
		} else if split[0] == "$" && split[1] == "ls" {
			continue
		} else {
			// The base case is either a file with size or a directory being "ls"-ed.
			if split[0] == "dir" {
				continue
			} else {
				// It's a file of format `12345 file.name` where the num is the size
				fileSize, _ := strconv.Atoi(split[0])
				curDirectory.size += fileSize

				// update size map
				_, ok := sizeMap[curDirectory.name]
				if curDirectory.size > 100000 {
					if ok {
						delete(sizeMap, curDirectory.name)
					}
				} else {
					sizeMap[curDirectory.name] = curDirectory.size
				}

				// Need to increase size of all directories above this
				pointer := curDirectory.parentDirectory
				for {
					if pointer == nil {
						break
					}
					pointer.size += fileSize

					// update size map
					_, ok := sizeMap[pointer.name]
					if pointer.size > 100000 {
						if ok {
							delete(sizeMap, pointer.name)
						}
					} else {
						sizeMap[pointer.name] = pointer.size
					}

					if pointer.parentDirectory == nil {
						break
					} else {
						pointer = pointer.parentDirectory
					}
				}
			}
		}
	}

	total := 0
	for _, val := range sizeMap {
		total += val
	}

	log.Println("Part One - total size of all directories with size less than 100000 - ", total)
}

func partTwo() {
	// Replace _ with `lines`
	lines, err := common.GetLinesFromFile("7/input.txt")
	common.FatalIf(err)

	var sizeMap = map[string]int{}

	var root = Directory{
		name:            "root",
		size:            0,
		subDirectories:  []*Directory{},
		parentDirectory: nil,
	}

	totalSpaceUsed := 0 // Add each individual file to this

	curDirectory := &root
	for _, line := range lines {
		if line == "$ cd /" {
			continue
		}

		split := strings.Split(line, " ")
		// cd command
		if split[0] == "$" && split[1] == "cd" {
			if split[2] == ".." {
				// go up a directory to one we've seen already
				curDirectory = curDirectory.parentDirectory
			} else {
				// Go down a level into a new directory (assuming no repeats in puzzle input)
				newChildDirectory := Directory{
					name:            curDirectory.name + "-" + split[2],
					size:            0,
					subDirectories:  []*Directory{},
					parentDirectory: curDirectory,
				}
				curDirectory.subDirectories = append(curDirectory.subDirectories, &newChildDirectory)
				curDirectory = &newChildDirectory
			}
		} else if split[0] == "$" && split[1] == "ls" {
			continue
		} else {
			// The base case is either a file with size or a directory being "ls"-ed.
			if split[0] == "dir" {
				continue
			} else {
				// It's a file of format `12345 file.name` where the num is the size
				fileSize, _ := strconv.Atoi(split[0])
				curDirectory.size += fileSize
				totalSpaceUsed += fileSize

				// update size map
				sizeMap[curDirectory.name] = curDirectory.size

				// Need to increase size of all directories above this
				pointer := curDirectory.parentDirectory
				for {
					if pointer == nil {
						break
					}
					pointer.size += fileSize

					// update size map
					sizeMap[pointer.name] = pointer.size

					if pointer.parentDirectory == nil {
						break
					} else {
						pointer = pointer.parentDirectory
					}
				}
			}
		}
	}

	totalSpaceAvailable := 70000000
	unusedSpaceNeeded := 30000000
	actualFreeSpace := 70000000 - totalSpaceUsed
	amountToDelete := unusedSpaceNeeded - actualFreeSpace
	actualDelete := totalSpaceAvailable
	for _, val := range sizeMap {
		if val >= amountToDelete && val < actualDelete {
			actualDelete = val
		}
	}

	log.Println("Part 2 - directory size to delete:", actualDelete)
}

func main() {
	partOne()
	partTwo()
}
