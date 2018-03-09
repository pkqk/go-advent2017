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

type Instruction struct {
	register, condition, check_register string
	amount, check_value int
}

func (instr Instruction) Execute(registers map[string]int) {
	run := false
	switch instr.condition {
	case ">":
		run = registers[instr.check_register] > instr.check_value
	case "<":
		run = registers[instr.check_register] < instr.check_value
	case ">=":
		run = registers[instr.check_register] >= instr.check_value
	case "<=":
		run = registers[instr.check_register] <= instr.check_value
	case "==":
		run = registers[instr.check_register] == instr.check_value
	case "!=":
		run = registers[instr.check_register] != instr.check_value
	default:
		fmt.Println("Missing operation", instr.condition)
		panic("Missing operation")
	}
	if run {
		registers[instr.register] += instr.amount
	}
}

func Part1(path string) {
	input, _ := os.Open(path)
	instructions := readInstructions(input)
	registers := runMachine(instructions)
	maxRegister(registers)
}

func maxRegister(registers map[string]int) {
	var key string
	var max int
	for reg, val := range registers {
		key = reg
		max = val
		break
	}
	for reg, val := range registers {
		if val > max {
			max = val
			key = reg
		}
	}
	fmt.Println("Max register", key, max)
}

func readInstructions(input *os.File) []Instruction {
	re := regexp.MustCompile(`(\w+) (inc|dec) ((-)?\d+) if (\w+) ([><!=]+) ((-)?\d+)`)
	instructions := make([]Instruction, 0)
	scanner := bufio.NewScanner(bufio.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		amount, _ := strconv.Atoi(matches[3])
		if matches[2] == "dec" {
			amount = -amount
		}
		check_value, _ := strconv.Atoi(matches[7])
		instructions = append(instructions, Instruction{
			register: matches[1],
			amount: amount,
			check_register: matches[5],
			condition: matches[6],
			check_value: check_value,
		})
	}
	return instructions
}

func runMachine(instructions []Instruction) map[string]int {
	registers := make(map[string]int)
	for _, instr := range instructions {
		instr.Execute(registers)
	}
	return registers
}
