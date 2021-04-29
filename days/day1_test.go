package days

import (
	"testing"

	"github.com/juansc/go-aoc2020/lib"
)

// Test using array of strings from example
func TestDay01Solver_Part1Simple(t *testing.T) {
	exampleInput := []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}

	actual, err := Day01Solver{}.Part1(exampleInput)
	if err != nil {
		t.Errorf("failed part1: encountered error %w", err)
	}
	expected := "514579"
	if actual != expected {
		t.Errorf("failed part1: expected %s, got %s", expected, actual)
	}
}

// Test using input file
func TestDay01Solver_Part1(t *testing.T) {
	lines, err := lib.ReadLines("../days/inputs/day1.txt")
	if err != nil{
		t.Errorf("failed part1: could not read input file %v", err)
		return
	}
	actual, err := Day01Solver{}.Part1(lines)
	if err != nil {
		t.Errorf("failed part1: encountered error %v", err)
		return
	}
	expected := "876459"
	if actual != expected {
		t.Errorf("failed part1: expected %s, got %s", expected, actual)
		return
	}
}

// TODO: Add tests for part2
