package main

import "fmt"
import "strconv"
import "io/ioutil"

func main() {
	fmt.Println("Advent of Code - Day 1")

	input, _ := ioutil.ReadFile("day1.txt")

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
