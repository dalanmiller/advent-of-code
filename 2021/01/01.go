package main

import (
	"strconv"
	"strings"
)

func parse_string(input string) []int {
	lines := strings.Split(input, "\n")

	var int_lines []int
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		int_lines = append(int_lines, n)
	}

	return int_lines
}

func run(input string) (int, int) {
	lines := parse_string(input)

	increases := 0
	rolling_increases := 0
	sums := make([]int, len(lines)-2)
	for i := 1; i < len(lines); i++ {

		// Count Part 1 increases
		if lines[i-1] < lines[i] {
			increases++
		}

		// Count Part 2 sums
		if i == 1 {
			sums[0] = lines[0] + lines[1] + lines[2]
		}

		// Only want to proceed if we have space remaining
		// . in the array
		if i < len(lines)-2 {
			sums[i] = lines[i] + lines[i+1] + lines[i+2]

			// looking back increase the count
			// needs to be encapsulated in above IF
			// or else would continue to check when the length
			// of the rolling sums array is shorter.
			if sums[i-1] < sums[i] {
				rolling_increases++
			}
		}
	}

	return increases, rolling_increases
}
