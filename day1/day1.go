package day1

import "fmt"
import "strconv"
import "io/ioutil"

func Run(path string) {
	input, _ := ioutil.ReadFile(path)

	fmt.Printf("%d\n", find_sum(input))
}

func find_sum(input []byte) int {
	sum := 0
	step := len(input) / 2
	for i, char := range input {
		next_char := input[(i+step)%len(input)]
		num, _ := strconv.Atoi(string(char))
		if char == next_char {
			sum += num
		}
	}
	return sum
}
