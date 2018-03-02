package main

import (
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

func SpiralPoint(n int) Point {
	return FollowSpiral(n, 1, Point{0, 0})
}

func FollowSpiral(n, step int, point Point) Point {
	if n <= 1 {
		return point
	}
	for r := 0; r < step; r++ {
		n -= 1
		point.x += 1
		if n == 1 {
			return point
		}
	}
	for u := 0; u < step; u++ {
		n -= 1
		point.y += 1
		if n == 1 {
			return point
		}
	}
	step += 1
	for l := 0; l < step; l++ {
		n -= 1
		point.x -= 1
		if n == 1 {
			return point
		}
	}
	for d := 0; d < step; d++ {
		n -= 1
		point.y -= 1
		if n == 1 {
			return point
		}
	}
	return FollowSpiral(n, step+1, point)
}

func main() {
	fmt.Println("Advent of Code - Day 3")
	n, _ := strconv.Atoi(os.Args[1])
	max := 4 * pointOnLine(SpiralPoint(n))
	grid := make([]int, max)
	grid[pointOnLine(Point{0, 0})] = 1
	for i := 2; i < n; i++ {
		point := SpiralPoint(i)
		idx := pointOnLine(point)
		sum := 0
		fmt.Println("----\naddition for", point)
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}
				new_p := Point{point.x + x, point.y + y}
				p := pointOnLine(new_p)
				sum += grid[p]
				fmt.Println("new_point", new_p, "p", p, "add", grid[p])
			}
		}
		grid[idx] = sum
		if sum > n {
			fmt.Println("point", point, "sum", sum)
			break
		}
		fmt.Println("n", i, "point", point, "value", grid[idx])
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

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
