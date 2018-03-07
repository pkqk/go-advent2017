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
  treeWeight int
  children []string
  root bool
}

// fwft (72) -> ktlj, cntj, xhth
var re = regexp.MustCompile(`(\w+) \((\d+)\)( -> (.*))?`)

func parseTree(lines []string) (string, map[string]*Node) {
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
    nodes[name] = &Node{name: name, weight: weight, children: children, root: true}
  }

  for _, node := range nodes {
    for _, childName := range node.children {
      nodes[childName].root = false
    }
  }

  var root string
  for _, node := range nodes {
    if node.root {
      root = node.name
    }
  }

  return root, nodes
}

func calcWeight(root *Node, nodes map[string]*Node) int {
  weight := root.weight
  for _, node := range root.children {
    weight += calcWeight(nodes[node], nodes)
  }
  root.treeWeight = weight
  return weight
}

func printWeights(root *Node, nodes map[string]*Node, indent string) {
  fmt.Println(indent, root.name, root.weight + root.treeWeight)
  for _, node := range root.children {
    printWeights(nodes[node], nodes, indent + " ")
  }
}

func findUnbalance(root *Node, nodes map[string]*Node, indent string) bool {
  if len(root.children) == 0 {
    return true
  }
  balanced := true
  firstWeight := nodes[root.children[0]].treeWeight
  for i := 1; i < len(root.children); i++ {
    if nodes[root.children[i]].treeWeight != firstWeight {
      balanced = false
    }
  }
  if !balanced {
    for _, name := range root.children {
      fmt.Println(indent, name, "weight", nodes[name].treeWeight, nodes[name].weight)
      findUnbalance(nodes[name], nodes, indent + "  ")
    }
  }
  return balanced
}

func main() {
	fmt.Println("Advent of Code - Day 7")

	input, _ := ioutil.ReadFile("day7.txt")
  lines := strings.Split(string(input), "\n")
  root, nodes := parseTree(lines)
  fmt.Println("root", nodes[root].name)
  calcWeight(nodes[root], nodes)
  fmt.Println("weight", nodes[root].treeWeight)
  //printWeights(nodes[root], nodes, "")
  findUnbalance(nodes[root], nodes, "")
}
