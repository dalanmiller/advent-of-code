package main

import (
	"bufio"
	"strconv"
	"strings"
)

type pair struct {
	lhStart int
	lhEnd   int
	rhStart int
	rhEnd   int
}

func readInput(input string) (pairs []pair) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		commaSplit := strings.Split(scanner.Text(), ",")

		left := strings.Split(commaSplit[0], "-")
		right := strings.Split(commaSplit[1], "-")

		leftStart, _ := strconv.Atoi(left[0])
		leftEnd, _ := strconv.Atoi(left[1])
		rightStart, _ := strconv.Atoi(right[0])
		rightEnd, _ := strconv.Atoi(right[1])

		pairs = append(pairs, pair{
			leftStart,
			leftEnd,
			rightStart,
			rightEnd,
		})
	}

	return pairs
}

func run(input string) (int, int) {
	pairs := readInput(input)

	partOne, partTwo := 0, 0
	for _, pair := range pairs {

		// Part one calculation
		if pair.lhStart == pair.rhStart {
			if pair.rhEnd <= pair.lhEnd || pair.lhEnd <= pair.rhEnd {
				partOne++
			}
		} else if pair.lhStart < pair.rhStart {
			if pair.rhStart >= pair.lhStart && pair.rhEnd <= pair.lhEnd {
				partOne++
			}

		} else {
			if pair.lhStart >= pair.rhStart && pair.lhEnd <= pair.rhEnd {
				partOne++
			}
		}

		// Part two checks
		// ! 5-7,1-4
		if pair.lhEnd >= pair.rhStart && pair.rhEnd >= pair.lhStart {
			partTwo++
		}

	}
	return partOne, partTwo
}
