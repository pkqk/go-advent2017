package day2

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func TestRowCheck() {
	input := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	output := []int{8, 4, 6}

	for i, row := range input {
		check := rowCheckSum(row)
		fmt.Printf("%v => %d == %d %v\n", row, output[i], check, check == output[i])
	}
}

func rowCheckSum(row []int) int {
	min := row[0]
	max := row[0]
	for _, val := range row {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return max - min
}

func rowEvenDivsor(row []int) int {
	for i, top := range row {
		for j, bottom := range row {
			if i == j {
				continue
			}
			if top%bottom == 0 {
				return top / bottom
			}
		}
	}
	return 0
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

func Part1(path string) {
  run(path, rowCheckSum)
}

func Part2(path string) {
  run(path, rowEvenDivsor)
}

func run(path string, checksum_func func([]int) int) {
	input, _ := ioutil.ReadFile(path)
	checksum := 0
	for _, line := range strings.Split(string(input), "\n") {
		row := scanLine(line)
		check := checksum_func(row)
		checksum += check
	}
	fmt.Println(checksum)
}
