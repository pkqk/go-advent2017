package day10

import "days"
import (
	"strings"
	"fmt"
	"io/ioutil"
	"strconv"
)

func init() {
	days.Register("10a", Part1)
	days.Register("10b", Part2)
}

func Part1(path string) {
	lengths := readInput(path)
	data := generateData()
	twist(data, lengths)
	fmt.Println("product of [0,1]:", data[0]*data[1])
}

func Part2(path string) {
}

func twist(data, lengths []int) {
	position := 0
	skipsize := 0
	for _, length := range(lengths) {
		reverse(data, position, length)
		position += length + skipsize
		skipsize++
	}
}

func reverse(data []int, position, length int) {
	for i, l := 0, length - 1; i < length/2; i,l = i+1, l-1 {
		a := (position+i) % len(data)
		b := (position+l) % len(data)
		data[a], data[b] = data[b], data[a]
	}
}

func readInput(path string) (result []int) {
	input, _ := ioutil.ReadFile(path)
	for _, s := range(strings.Split(string(input), ",")) {
		if i, err := strconv.Atoi(s); err != nil {
			fmt.Println(err)
			return
		} else {
			result = append(result, i)
		}
	}
	return
}

func generateData() (list []int) {
	list = make([]int, 256)
	for idx, _ := range(list) {
		list[idx] = idx
	}
	return
}
