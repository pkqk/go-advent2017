package main

import (
  "fmt"
  "os"
)
import (
  "day1"
  "day2"
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
  default:
    fmt.Println("Implement Day", day)
  }
}
