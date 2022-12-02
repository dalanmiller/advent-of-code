package main

import (
	"strings"
)

func readInputOne(input string) int {
	lines := strings.Split(input, "\n")

	score := 0
	for _, line := range lines {
		vals := strings.Split(line, " ")
		o := vals[0]
		m := vals[1]

		switch o {
		case "A":
			if m == "Y" { // win
				score += 8
			} else if m == "Z" { // lose
				score += 3
			} else {
				score += 4 // draw
			}
		case "B":
			if m == "Z" { //win
				score += 9
			} else if m == "X" { // lose
				score += 1
			} else { // draw
				score += 5
			}
		case "C":
			if m == "X" {
				score += 7
			} else if m == "Y" {
				score += 2
			} else {
				score += 6
			}
		}
	}

	return score
}

func readInputTwo(input string) int {
	lines := strings.Split(input, "\n")

	// X == lose
	// y == draw
	// z == win

	score := 0
	for _, line := range lines {
		vals := strings.Split(line, " ")
		o := vals[0]
		m := vals[1]

		switch o {
		case "A":
			switch m {
			case "X":
				score += 3
			case "Y":
				score += 4
			case "Z":
				score += 8
			}
		case "B":
			switch m {
			case "X":
				score += 1
			case "Y":
				score += 5
			case "Z":
				score += 9
			}
		case "C":
			switch m {
			case "X":
				score += 2
			case "Y":
				score += 6
			case "Z":
				score += 7
			}
		}
	}

	return score
}

func run(input string) (int, int) {
	scoreOne := readInputOne(input)
	scoreTwo := readInputTwo(input)

	return scoreOne, scoreTwo

}