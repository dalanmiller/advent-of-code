package main

import (
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	raw_fish := strings.Split(input, ",")

	fishes := make([]int, len(raw_fish))
	for i, fish := range raw_fish {
		n, _ := strconv.Atoi(fish)
		fishes[i] = n
	}

	return fishes
}

func run(input string, days int) int {
	fishes := parseInput(input)

	// Extra length to store room for the new fishjp
	fishCounts := make([]int, 10)

	for _, fish := range fishes {
		fishCounts[fish]++
	}

	// It took me a stupidly and embarassingly
	// long time to get this right
	for i := 0; i < days; i++ {
		fishCounts[9] = fishCounts[0]

		// Much slower than the following for loop for shifting
		// fishCounts = append(fishCounts[1:], fishCounts[0])

		for j := 1; j < len(fishCounts); j++ {
			fishCounts[j-1] = fishCounts[j]
		}

		fishCounts[6] += fishCounts[9]
		fishCounts[8] = fishCounts[9]
		fishCounts[9] = 0
	}

	sum := 0
	for _, count := range fishCounts {
		sum += count
	}
	return sum
}
