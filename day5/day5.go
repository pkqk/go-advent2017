package day5

import (
	"bufio"
	"days"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	days.Register("5a", Part1)
	days.Register("5b", Part2)
}


func Part1(path string) {
	input, _ := ioutil.ReadFile(path)
	jumps := readInstructions(string(input))

	fmt.Printf("%d\n", runMachine(jumps, false))
}
func Part2(path string) {
	input, _ := ioutil.ReadFile(path)
	jumps := readInstructions(string(input))

	fmt.Printf("%d\n", runMachine(jumps, true))
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

func runMachine(jumps []int, part2 bool) int {
	steps := 0
	pos := 0
	var jump int
	for {
		if pos < 0 || pos >= len(jumps) {
			break
		}
		jump = jumps[pos]
		if part2 && jump > 2 {
			jumps[pos]--
		} else {
			jumps[pos]++
		}
		pos += jump
		steps++
	}
	return steps
}
