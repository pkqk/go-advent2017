package day12

import "days"
import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(\d+) <-> ([\d, ]+)`)
type GroupMap map[string]map[string]bool

func init() {
	days.Register("12a", Part1)
	days.Register("12b", Part2)

}

func Part1(path string) {
	groups := buildGroups(path)
	visited := make(map[string]bool)
	fmt.Println("visits[0]", deepCount("0", groups, visited))
}

func buildGroups(path string) GroupMap {
	groups := make(GroupMap)
	var scanner *bufio.Scanner
	if file, err := os.Open(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	else {
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		src := matches[1]
		pipes := strings.Split(matches[2], ", ")
		for _, dest := range(pipes) {
			if groups[src] == nil {
				groups[src] = make(map[string]bool)
			}
			if groups[dest] == nil {
				groups[dest] = make(map[string]bool)
			}
			groups[src][dest] = true
			groups[dest][src] = true
		}
	}
	return groups
}

func deepCount(start string, groups map[string]map[string]bool, seen map[string]bool) int {
	if _, pass := seen[start]; pass {
		return 0
	}
	seen[start] = true
	count := 1
	for dest, _ := range(groups[start]) {
		count += deepCount(dest, groups, seen)
	}
	return count
}

func Part2(path string) {
}
