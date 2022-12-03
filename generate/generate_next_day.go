package main

import (
	"AdventOfCode2022/common"
	"flag"
	"fmt"
	"os"
)

var day int

func init() {
	flag.IntVar(&day, "day", -1, "day to generate")
}

func main() {
	flag.Parse()

	if day <= 0 {
		os.Exit(1)
	}

	err := os.Mkdir(fmt.Sprint(day), os.ModePerm)
	common.FatalIf(err)

	filename := fmt.Sprintf("%d/day%d.go", day, day)
	os.Create(filename)
	dat, err := os.ReadFile("generate/template.txt")
	common.FatalIf(err)
	err = os.WriteFile(filename, dat, 0644)
	common.FatalIf(err)

	os.Create(fmt.Sprintf("%d/input.txt", day))
}
