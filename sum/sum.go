package sum

import "strconv"

func Sum(answer map[string]int) int {
	var sum int
	for _, arg := range answer {
		sum += arg
	}
	return sum
}

func SumStrings(numbers []string) (int, error) {
	var sum int
	for _, str := range numbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		sum += num
	}
	return sum, nil
}
