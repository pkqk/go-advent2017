package main

import (
  "day1"
  "fmt"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("usage: go run runner.go n input")
    os.Exit(1)
  }
  if n, err := strconv.Atoi(os.Args[1]); err == nil {
    fmt.Println("Advent of Code - Day", n)
    switch n {
    case 1:
      day1.Run(os.Args[2])
    default:
      fmt.Println("Implement Day", n)
    }
  } else {
    fmt.Println(os.Args[1], "should be the number of the Day")
    os.Exit(1)
  }
}
