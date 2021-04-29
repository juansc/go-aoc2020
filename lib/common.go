package lib

import (
	"io/ioutil"
	"strings"
)

// ReadLines returns the lines from a file as a slice of strings.
func ReadLines(fileName string) ([]string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), "\n"), nil
}
