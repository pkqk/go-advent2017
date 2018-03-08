package day7

import (
	"fmt"
	"io/ioutil"
  "strconv"
	"strings"
)

type LoadNode struct {
  name string
  weight int
  treeWeight int
  children []string
  root bool
}

type TreeNode struct {
	name string
	weight int
	children []*TreeNode
	totalWeight int
	balanced bool
}

func parseTree(lines []string) *TreeNode {
  nodes := make(map[string]*LoadNode)

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
    nodes[name] = &LoadNode{name: name, weight: weight, children: children, root: true}
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

  return buildTree(root, nodes)
}

func buildTree(name string, nodes map[string]*LoadNode) *TreeNode {
	node := nodes[name]
	children := make([]*TreeNode, 0)

	for _, child := range node.children {
		children = append(children, buildTree(nodes[child].name, nodes))
	}
	return &TreeNode{name: node.name, weight: node.weight, children: children}
}

func calcWeight(node *TreeNode) int {
	sum := node.weight
	for _, child := range node.children {
		sum += calcWeight(child)
	}
	node.totalWeight = sum
	return sum
}

func calcBalance(node *TreeNode) bool {
	if len(node.children) < 2 {
		node.balanced = true
		return true
	}
	balanced := true
	weight := node.children[0].totalWeight
	for i := 1; i < len(node.children); i++ {
		balanced = balanced && weight == node.children[i].totalWeight
	}
	node.balanced = balanced
	return balanced
}

func findProblem(node *TreeNode, goalWeight int) {
	if !node.balanced {
		balanced := true
		weights := make(map[int]int)
		for _, child := range node.children {
			balanced = balanced && child.balanced
			weights[child.totalWeight]++
		}
		if !balanced {
			var badWeight, goodWeight int
			for weight, count := range(weights) {
				if count == 1 {
					badWeight = weight
				} else {
					goodWeight = weight
				}
			}
			if badWeight == 0 {
				fmt.Println("found problem", node.name)
				diff := goalWeight - node.totalWeight
				fmt.Println("answer", node.weight + diff)
			} else {
				for _, child := range node.children {
					if child.totalWeight == badWeight {
						findProblem(child, goodWeight)
					}
				}
			}
		}
	}
}

func Part2(path string) {
	input, _ := ioutil.ReadFile(path)
  lines := strings.Split(string(input), "\n")
  root := parseTree(lines)
	calcWeight(root)
	calcBalance(root)
	findProblem(root, 0)
}
