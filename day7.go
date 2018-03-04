package main

import (
	"fmt"
	"io/ioutil"
  "regexp"
  "strconv"
	"strings"
)

type Node struct {
  name string
  weight int
  children []string
  root bool
}

// fwft (72) -> ktlj, cntj, xhth
var re = regexp.MustCompile(`(\w+) \((\d+)\)( -> (.*))?`)

func parseLines(lines []string) {
  nodes := make(map[string]*Node)

  for _, line := range lines {
    matches := re.FindStringSubmatch(line)
    name, weightString, childList := matches[1], matches[2], matches[4]
    weight, _ := strconv.Atoi(weightString)
    var children []string
    if len(childList) > 0 {
      children = strings.Split(childList, ", ")
    } else {
      children = []string{}
    }
    nodes[name] = &Node{name, weight, children, true}
  }

  for _, node := range nodes {
    for _, childName := range node.children {
      nodes[childName].root = false
    }
  }

  for _, node := range nodes {
    if node.root {
      fmt.Println("root", node.name)
    }
  }
}

func main() {
	fmt.Println("Advent of Code - Day 7")

	input, _ := ioutil.ReadFile("day7.txt")
  lines := strings.Split(string(input), "\n")
  parseLines(lines)
}
