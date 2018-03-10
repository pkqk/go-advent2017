package main

import (
	_ "day1"
	_ "day2"
	_ "day3"
	_ "day4"
	_ "day5"
	_ "day6"
	_ "day7"
	_ "day8"
	_ "day9"
	"days"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run runner.go n input")
		os.Exit(1)
	}
	day, input := os.Args[1], os.Args[2]
	fmt.Println("Advent of Code - Day", day)
	days.Call(day, input)
}
