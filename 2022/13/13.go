package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Outcome int

const (
	LESS Outcome = iota - 1
	EQUAL
	GREATER
)

func compare(left, right []interface{}) Outcome {

	// 1. If both values are integers, the lowest one should come first. If otherwise, the values
	// . are not in the right order. If same, continue.
	// 2. If both values are lists, compare the first value of each list, then the second value, and so on.
	// . If the left list runs out of items first, then the inputs are in the right order
	// . If right runs out first, then not right order.
	// . If same length -- continue.
	// 3. If exactly one value is an integer, convert the integer to a list which contains
	// . that integer as its only value then retry the comparison.

	// https://github.com/alexchao26/advent-of-code-go/blob/main/2022/day13/main.go#L103
	// Had no idea you could effectively abuse interface in golang in this way
	// Basically parse using json.Unmarshal then hope for the best when casting?

	for i := 0; i < len(left); i++ {
		if i > len(right)-1 {
			return GREATER
		}

		// Have to use float64 and not int as this is what the json Unmarshal uses
		leftN, isLeftInt := left[i].(float64)
		rightN, isRightInt := right[i].(float64)

		leftList, _ := left[i].([]interface{})
		rightList, _ := right[i].([]interface{})

		if isLeftInt && isRightInt {
			// log.Println(leftN, rightN)
			if leftN != rightN {
				switch leftN < rightN {
				case true:
					return LESS
				case false:
					return GREATER
				}
			} else {
				continue
			}
		} else if isLeftInt || isRightInt {
			if isLeftInt {
				leftList = []interface{}{leftN}
			} else {
				rightList = []interface{}{rightN}
			}
		}
		outcome := compare(leftList, rightList)
		switch outcome {
		case LESS, GREATER:
			return outcome
		case EQUAL:
			continue

		}
	}

	if len(left) < len(right) {
		return LESS
	}

	return EQUAL
}

func readInput(input io.Reader) (pairs [][2][]interface{}) {

	b, _ := io.ReadAll(input)

	for _, packetPair := range strings.Split(string(b), "\n\n") {
		packets := strings.Split(packetPair, "\n")
		pairs = append(pairs, [2][]interface{}{
			parsePacketString(packets[0]),
			parsePacketString(packets[1]),
		})
	}

	return pairs
}

func parsePacketString(p string) []interface{} {
	// [1,[2,[3,[4,[5,6,7]]]],8,9]

	packet := []interface{}{}
	json.Unmarshal([]byte(p), &packet)
	return packet
}

func runPartOne(input io.Reader) int {
	pairs := readInput(input)

	s := 0
	for i, p := range pairs {
		// log.Println(p[0], p[1])
		if compare(p[0], p[1]) == LESS {
			s += i + 1
		}
	}

	return s
}

func runPartTwo(input io.Reader) int {
	b, _ := io.ReadAll(input)

	lines := strings.Split(string(b), "\n")

	packets := make([][]interface{}, 0, len(lines))
	for _, line := range lines {
		if line == "\n" || line == "" {
			continue
		}

		packets = append(packets, parsePacketString(line))
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == LESS
	})

	s := 1
	for i, packet := range packets {
		if fmt.Sprint(packet) == "[[2]]" || fmt.Sprint(packet) == "[[6]]" {
			s *= i + 1
		}
	}

	return s
}
