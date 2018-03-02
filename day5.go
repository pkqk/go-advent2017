package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 5")

	input, _ := ioutil.ReadFile("day5.txt")
	jumps := readInstructions(string(input))

	fmt.Printf("%d\n", runMachine(jumps))
}

func readInstructions(input string) []int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result
		}
		result = append(result, x)
	}
	return result
}

func runMachine(jumps []int) int {
	steps := 0
	pos := 0
	var jump int
	for {
		if pos < 0 || pos >= len(jumps) {
			break
		}
		jump = jumps[pos]
		if jump > 2 {
			jumps[pos]--
		} else {
			jumps[pos]++
		}
		pos += jump
		steps++
	}
	return steps
}
