package sum

func Sum(answer map[string]int) int {
	var sum int
	for _, arg := range answer {
		sum += arg
	}
	return sum
}
