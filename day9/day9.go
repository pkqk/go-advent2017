package day9

import "days"
import (
	"bufio"
	"io"
	"fmt"
	"os"
)

func init() {
	days.Register("9a", Part1)
	// days.Register("9b", Part2)
}

func Part1(path string) {
	score := make(chan int, 1)
	instream := make(chan rune)
	safestream := make(chan rune)
	cleanstream := make(chan rune)
	go readInput(path, instream)
	go bangEscape(instream, safestream)
	go garbageCollect(safestream, cleanstream)
	go keepScore(cleanstream, score)
	total := totalScore(score)
	fmt.Println("total score:", total)
}

func readInput(path string, output chan rune) {
	var input *bufio.Reader
	if path == "-" {
		input = bufio.NewReader(os.Stdin)
	} else {
		file, _ := os.Open(path)
		input = bufio.NewReader(file)
	}
	for {
		char, _, err := input.ReadRune()
		if err == io.EOF {
			close(output)
			break
		}
		output <- char
	}
}

func bangEscape(input chan rune, output chan rune) {
	skip := false
	for {
		if r, more := <-input; more {
			if skip {
				skip = false
			} else if r == '!' {
				skip = true
			} else {
				output <- r
			}
		} else {
			close(output)
			break
		}
	}
}

func garbageCollect(input chan rune, output chan rune) {
	inGarbage := false
	for {
		if r, more := <-input; more {
			if r == '<' {
				inGarbage = true
			} else if r == '>' {
				inGarbage = false
			} else if !inGarbage {
				output <- r
			}
		} else {
			close(output)
			break
		}
	}
}

func keepScore(input chan rune, output chan int) {
	depth := 0
	for {
		if r, more := <-input; more {
			switch r {
			case '{':
				depth++
				output <- depth
			case '}':
				depth--
			}
		} else {
			close(output)
			break
		}
	}
}

func totalScore(input chan int) int {
	score := 0
	for {
		if val, more := <-input; more {
			score += val
		} else {
			return score
		}
	}
}
