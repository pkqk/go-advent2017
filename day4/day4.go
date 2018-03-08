package day4

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func Part1(path string) {
	input, _ := ioutil.ReadFile(path)

	run(input, false)
}

func Part2(path string) {
	input, _ := ioutil.ReadFile(path)

	run(input, true)
}

func run(input []byte, part2 bool) {
	sum := 0
	for _, line := range strings.Split(string(input), "\n") {
		words := strings.Split(line, " ")
		count := make(map[string]int)
		valid := true
		for _, word := range words {
			if part2 {
				// sort word letters
				letters := strings.Split(word, "")
				sort.Strings(letters)
				word = strings.Join(letters, "")
			}
			count[word] += 1
			if count[word] > 1 {
				valid = false
			}
		}
		if valid {
			sum += 1
		}
	}
	fmt.Printf("%d\n", sum)
}
