package day6

import "days"
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	days.Register("6a", Part1)
	days.Register("6b", Part2)
}

var seen map[string]bool

func Part1(path string) {
	seen = make(map[string]bool)

	input, _ := ioutil.ReadFile(path)
	state := scanLine(string(input))
	//state = []int{0, 2, 7, 0}

	steps := 0

	for ; ; steps++ {
		fmt.Println(state)
		if visit(state) {
			break
		}
		redistribute(state)
	}

	fmt.Println("state", state, "steps", steps)
}

func visit(state []int) bool {
	key := keyFor(state)
	_, visited := seen[key]
	seen[key] = true
	return visited
}

func keyFor(state []int) string {
	blocks := make([]string, len(state))
	for i, n := range state {
		blocks[i] = strconv.Itoa(n)
	}
	return strings.Join(blocks, ":")
}

func redistribute(state []int) {
	max := findMax(state)
	val := state[max]
	state[max] = 0
	i := (max + 1) % len(state)
	for val > 0 {
		state[i]++
		val--
		i = (i + 1) % len(state)
	}
}

func findMax(state []int) int {
	max := 0
	for i := 1; i < len(state); i++ {
		if state[i] > state[max] {
			max = i
		}
	}
	return max
}

func scanLine(input string) []int {
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
