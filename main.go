package main

import (
	"fmt"
	"os"

	"github.com/juansc/go-aoc2020/days"
	"github.com/juansc/go-aoc2020/lib"
)

// As you add more days add more solvers here
var runners = map[string]days.DaySolver{
	"1": days.Day01Solver{},
}

func main() {
	day :=os.Args[1]
	runner, ok := runners[day]
	if !ok {
		fmt.Println("no solution for day %s", day)
		return
	}
	lines, err := lib.ReadLines(fmt.Sprintf("./days/inputs/day%s.txt", day))
	if err != nil {
		fmt.Printf("could not read file: %v\n", err)
	}

	part1, err := runner.Part1(lines)
	if err != nil{
		fmt.Printf("encountered error with part1: %v\n", err)
		os.Exit(1)
		return
	}
	part2, err := runner.Part2(lines)
	if err != nil{
		fmt.Printf("encountered error with part2: %v\n", err)
		os.Exit(1)
		return
	}
	fmt.Println("Answer for Part1: ", part1)
	fmt.Println("Answer for Part2: ", part2)
}