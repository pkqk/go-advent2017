package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 4")

	input, _ := ioutil.ReadFile("day4.txt")

	fmt.Printf("%d\n", checkValidity(input))
}

func checkValidity(input []byte) int {
	sum := 0
	for _, line := range strings.Split(string(input), "\n") {
		words := strings.Split(line, " ")
		count := make(map[string]int)
		valid := true
		for _, word := range words {
			// part2
			// sort word letters
			letters := strings.Split(word, "")
			sort.Strings(letters)
			word := strings.Join(letters, "")
			fmt.Println(word)
			// end part2
			count[word] += 1
			if count[word] > 1 {
				valid = false
				fmt.Println("not valid", line)
			}
		}
		if valid {
			fmt.Println("valid", line)
			sum += 1
		}
	}
	return sum
}
