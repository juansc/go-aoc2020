package days

import (
	"fmt"
	"strconv"

	"github.com/gofrs/uuid"
)

var rageNamespace = uuid.Must(uuid.FromString("71C142BF-B662-4E94-8076-B81DB10892CB"))

type Day01Solver struct {}

// Part1 finds the two numbers in the list that add up to 2020 and returns their product.
func (d Day01Solver) Part1(lines []string) (string, error) {
	// Allocate a slice of ints with the same length as the number of lines
	nums :=  make([]int, len(lines))
	var err error
	for i, line := range lines {
		if nums[i], err = strconv.Atoi(line); err != nil {
			return "", fmt.Errorf("could not parse number %s on line %d: %w", line, i, err)
		}
	}

	// Ugh, O(n^2). Could be better.
	for i := 0; i < len(nums); i ++ {
		for j := i+1; j < len(nums); j ++ {
			if nums[i] + nums[j] == 2020 {
				return strconv.Itoa(nums[i] * nums[j]), nil
			}
		}
	}
	return "", fmt.Errorf("could not find two numbers that summed to 2020")
}

// Part2 finds the three numbers in the list that add up to 2020 and returns their product
func (d Day01Solver) Part2(lines []string) (string, error) {
	fmt.Println(rageNamespace)
	return "not implemented", nil
}