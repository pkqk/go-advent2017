package day6

import (
	"fmt"
	"io/ioutil"
)

var seenCounts map[string]int

func Part2(path string) {
	seenCounts = make(map[string]int)

	input, _ := ioutil.ReadFile(path)
	state := scanLine(string(input))

	steps := 0
	var distance int

	for ; ; steps++ {
		fmt.Println(state)
		if visited, last_step := visitAndGetLastStep(state, steps); visited {
			distance = steps - last_step
			break
		}
		redistribute(state)
	}

	fmt.Println("state", state, "steps", steps, "distance", distance)
}

func visitAndGetLastStep(state []int, step int) (bool, int) {
	key := keyFor(state)
	last_step, visited := seenCounts[key]
	seenCounts[key] = step
	return visited, last_step
}
