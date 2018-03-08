package day8

import "days"
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"regexp"
)

func init() {
	days.Register("8a", Part1)
	// days.Register("8b", Part2)
}

type instruction struct {
	register, condition, check_register string
	inc, amount, check_value int
}

func Part1(path string) {
	input, _ := os.Open(path)
	instructions := readInstructions(input)
	registers := runMachine(instructions)

	fmt.Printf("%d\n", registers)
}

func readInstructions(input *os.File) []instruction {
	re := regexp.MustCompile(`(\w+) (inc|dec) ((-)?\d+) if (\w+) ([><=]+) ((-)?\d+)`)
	instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(bufio.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		fmt.Println(line)
		register := matches[1]
		inc, _ := strconv.Atoi(matches[3])
		if matches[2] == "dec" {
			inc = -inc
		}
		fmt.Println(register, inc)
	}
	return instructions
}

func runMachine(instructions []instruction) map[string]int {
	return make(map[string]int)
}
