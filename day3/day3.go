package day3

import (
	"fmt"
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

func Part1(input string) {
	n, _ := strconv.Atoi(input)
	point := SpiralPoint(n)
	fmt.Println("Position", point)
	fmt.Println("Steps", abs(point.x)+abs(point.y))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
