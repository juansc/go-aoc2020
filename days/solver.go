package days

// DaySolver is a type that has two methods: Part1 and Part2. Both of these expect an array of strings
// and return a string and an error.
type DaySolver interface{
	Part1(lines []string) (string, error)
	Part2(lines []string) (string, error)
}
