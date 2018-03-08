package day1

import (
  "fmt"
  "strconv"
  "io/ioutil"
)

func Part1(path string) {
	input, _ := ioutil.ReadFile(path)

	sum := 0
	for i, char := range input {
		next_char := input[(i+1)%len(input)]
		num, _ := strconv.Atoi(string(char))
		if char == next_char {
			sum += num
		}
	}

	fmt.Printf("%d\n", sum)
}

func Part2(path string) {
	input, _ := ioutil.ReadFile(path)

	sum := 0
	step := len(input) / 2
	for i, char := range input {
		next_char := input[(i+step)%len(input)]
		num, _ := strconv.Atoi(string(char))
		if char == next_char {
			sum += num
		}
	}

	fmt.Printf("%d\n", sum)
}
