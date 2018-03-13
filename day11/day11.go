package day11

import "days"
import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func init() {
	days.Register("11a", Part1)
	days.Register("11b", Part2)
}

func Part1(path string) {
	input, _ := os.Open(path)
	moves := readMoves(input)
	start := HexLoc{0,0,0}
	var loc HexLoc
	for _, move := range moves {
		loc = loc.move(move)
	}
	fmt.Println(loc.distance(start))
}

func Part2(path string) {
}

// using
// http://keekerdc.com/2011/03/hexagon-grids-coordinate-systems-and-distance-calculations/
// we can represent a hexgrid using 3d co-ords with some constraints on movement
// x+y+z = 0
// x -> E
// y -> NW
// z -> SW
type HexLoc struct {
	x,y,z int
}

func (src HexLoc) move(dir string) (dst HexLoc) {
	dst = src
	switch dir {
	case "n":
		dst = HexLoc{src.x, src.y+1, src.z-1}
	case "ne":
		dst = HexLoc{src.x+1, src.y, src.z-1}
	case "se":
		dst = HexLoc{src.x+1, src.y-1, src.z}
	case "s":
		dst = HexLoc{src.x, src.y-1, src.z+1}
	case "sw":
		dst = HexLoc{src.x-1, src.y, src.z+1}
	case "nw":
		dst = HexLoc{src.x-1, src.y+1, src.z}
	}
	return
}

func (point HexLoc) valid() bool {
	return (point.x + point.y + point.z) == 0
}

func (point HexLoc) distance(target HexLoc) int {
	dx := abs(point.x - target.x)
	max := dx
	dy := abs(point.y - target.y)
	if dy > max {
		max = dy
	}
	dz := abs(point.z - target.z)
	if dz > max {
		max = dz
	}
	return max
}

func readMoves(input *os.File) []string {
	scanner := bufio.NewScanner(bufio.NewReader(input))
	scanner.Split(scanCommas)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

// copied from
// https://golang.org/src/bufio/scan.go?s=11522:11600#L330
func scanCommas(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
	 		return 0, nil, nil
	 	}
	 	if i := bytes.IndexByte(data, ','); i >= 0 {
	 		// Up to a comma
	 		return i + 1, data[0:i], nil
	 	}
	 	// If we're at EOF, we have the last element after a comma
	 	if atEOF {
	 		return len(data), data, nil
	 	}
	 	// Request more data.
	 	return 0, nil, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
