package main

import (
	"fmt"
	"math"
	"strings"
)

// CH -> B
// CB, BH

func parseInput(input string) ([]string, map[string]string, map[string][]string) {
	lines := strings.Split(input, "\n")
	polymerTemplate := strings.Split(lines[0], "")

	insertionMap := make(map[string]string, len(lines[2:]))
	for _, line := range lines[2:] {
		pair := strings.Split(line, " -> ")
		insertionMap[pair[0]] = pair[1]
	}

	genMap := make(map[string][]string)
	for k, v := range insertionMap {
		genMap[k] = []string{
			fmt.Sprintf("%s%s", string(k[0]), v),
			fmt.Sprintf("%s%s", v, string(k[1])),
		}
	}

	return polymerTemplate, insertionMap, genMap
}

func run(input string, steps int) uint {
	template, insMap, genMap := parseInput(input)

	pairCount := make(map[string]uint, len(insMap))
	for i := 1; i < len(template); i++ {
		pair := strings.Join(template[i-1:i+1], "")
		pairCount[pair]++
	}

	for i := 0; i < steps; i++ {
		newPairCount := make(map[string]uint, len(insMap))
		for pair, n := range pairCount {

			newPairs := genMap[pair]

			newPairCount[newPairs[0]] += n
			newPairCount[newPairs[1]] += n
		}
		pairCount = newPairCount
	}

	count := make(map[string]uint)

	for k, v := range pairCount {
		for _, side := range []string{string(k[0]), string(k[1])} {
			count[side] += uint(v)
		}
	}

	max := uint(0)
	min := uint(math.MaxUint)
	for _, v := range count {

		if v%2 == 1 {
			v = v/2 + 1
		} else {
			v = v / 2
		}

		if max < v {
			max = v
		}

		if min > v {
			min = v
		}
	}

	return max - min
}
