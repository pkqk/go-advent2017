package day3

import (
	"fmt"
	"strconv"
)

func Part2(input string) {
	n, _ := strconv.Atoi(input)
	max := 4 * pointOnLine(SpiralPoint(n))
	grid := make([]int, max)
	grid[pointOnLine(Point{0, 0})] = 1
	for i := 2; i < n; i++ {
		point := SpiralPoint(i)
		idx := pointOnLine(point)
		sum := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}
				new_p := Point{point.x + x, point.y + y}
				p := pointOnLine(new_p)
				sum += grid[p]
			}
		}
		grid[idx] = sum
		if sum > n {
			fmt.Println("point", point, "sum", sum)
			break
		}
	}
}

// this was overkill, I could've just used a map[point]int
// using idea from
// https://math.stackexchange.com/questions/374694/mapping-from-1d-line-to-2d-plane-an-infinite-piece-of-rope-covering-2d-plane-wi#comment806690_374700
func cantorOffset(point Point) int {
	offset := 0
	if point.x < 0 {
		offset += 1
	}
	if point.y < 0 {
		offset += 2
	}
	return offset
}

func cantorPair(point Point) int {
	k1 := abs(point.x)
	k2 := abs(point.y)
	num := (2 * (k1 + k2) * (k1 + k2 + 1)) + 4*k2
	return num
}

func pointOnLine(point Point) int {
	return cantorPair(point) + cantorOffset(point)
}
