package main

import (
  "fmt"
  "os"
)
import (
  "day1"
  "day2"
  "day3"
  "day4"
  "day5"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("usage: go run runner.go n input")
    os.Exit(1)
  }
  day, input := os.Args[1], os.Args[2]
  fmt.Println("Advent of Code - Day", day)
  switch day {
  case "1a":
    day1.Part1(input)
  case "1b":
    day1.Part2(input)
  case "2a":
    day2.Part1(input)
  case "2b":
    day2.Part2(input)
  case "3a":
    day3.Part1(input)
  case "3b":
    day3.Part2(input)
  case "4a":
    day4.Part1(input)
  case "4b":
    day4.Part2(input)
  case "5a":
    day5.Part1(input)
  case "5b":
    day5.Part2(input)
  default:
    fmt.Println("Implement Day", day)
  }
}
