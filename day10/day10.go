package day10

import "days"
import (
	"strings"
	"fmt"
	"encoding/hex"
	"io/ioutil"
	"strconv"
)

func init() {
	days.Register("10a", Part1)
	days.Register("10b", Part2)
}

func Part1(path string) {
	lengths := readInput(path)
	data := generateData()
	twist(data, lengths)
	fmt.Println("product of [0,1]:", data[0]*data[1])
}

func Part2(path string) {
	lengths := getLengthBytes(path)
	data := generateBytes()
	twistRounds(data, lengths)
	hash := denseHash(data)
	fmt.Println(hex.EncodeToString(hash))
}

func twist(data, lengths []int) {
	position := 0
	skipsize := 0
	for _, length := range(lengths) {
		reverse(data, position, length)
		position += length + skipsize
		skipsize++
	}
}

func twistRounds(data, lengths []byte) {
	position := 0
	skipsize := 0
	for round := 0; round < 64; round++ {
		for _, length := range(lengths) {
			reverseBytes(data, position, length)
			position += int(length) + skipsize
			skipsize++
		}
	}
}

func denseHash(data []byte) []byte {
	result := make([]byte, 16)
	for i := 0; i < 256; i+=16 {
		collector := data[i]
		for count := 1; count < 16; count++ {
			collector ^= data[i+count]
		}
		result[i/16] = collector
	}
	return result
}

func reverse(data []int, position, length int) {
	for i, l := 0, length - 1; i < length/2; i,l = i+1, l-1 {
		a := (position+i) % len(data)
		b := (position+l) % len(data)
		data[a], data[b] = data[b], data[a]
	}
}

func reverseBytes(data []byte, position int, length byte) {
	halfway := int(length/2)
	for i, l := 0, int(length - 1); i < halfway; i,l = i+1, l-1 {
		a := (position+i) % len(data)
		b := (position+l) % len(data)
		data[a], data[b] = data[b], data[a]
	}
}

func readInput(path string) (result []int) {
	input, _ := ioutil.ReadFile(path)
	for _, s := range(strings.Split(string(input), ",")) {
		if i, err := strconv.Atoi(s); err != nil {
			fmt.Println(err)
			return
		} else {
			result = append(result, i)
		}
	}
	return
}

func getInputBytes(path string) ([]byte) {
	if result, err := ioutil.ReadFile(path); err != nil {
		panic("Failed to read input")
	} else {
		return result
	}
}

func getLengthBytes(path string) []byte {
	input := getInputBytes(path);
	input = append(input, 17, 31, 73, 47, 23)
	return input
}

func generateData() (list []int) {
	list = make([]int, 256)
	for idx, _ := range(list) {
		list[idx] = idx
	}
	return
}

func generateBytes() (list []byte) {
	list = make([]byte, 256)
	for idx, _ := range(list) {
		list[idx] = byte(idx)
	}
	return
}
