package days

import (
  "fmt"
)

type Solution func(string)

var registry map[string]Solution

func init() {
  registry = make(map[string]Solution)
}

func Register(name string, solution Solution) {
  registry[name] = solution
}

func Call(name, input string) {
  if fn, ok := registry[name]; ok {
    fn(input)
  } else {
    fmt.Println("Implement Day", name)
  }
}
