package main

import (
	"strings"
)

func readInput(input string) ([]rune, []rune) {

	lines := strings.Split(input, "\n")

	commonItems := []rune{}
	badgeTypes := []rune{}

	badgeMap := map[rune]int{}
	for i, line := range lines {
		common := make(map[rune]bool, len(lines[0]))
		addedForThisLine := map[rune]bool{}
		for j, char := range line {

			// Next if statements check if we need to first
			// insert into the badgeMap and init at 1 otherwise
			// increment but only if it's the first time for this line
			if _, ok := badgeMap[char]; !ok {
				badgeMap[char] = 1
				addedForThisLine[char] = true
			} else if _, ok := addedForThisLine[char]; !ok {
				badgeMap[char] += 1
				addedForThisLine[char] = true
			}

			// If we are on the left half, then add to map
			if j < len(line)/2 {
				common[char] = true

				// If we are in second half of line, then check
				//  if we've already seen this one
			} else if _, ok := common[char]; ok && len(commonItems) < i+1 {
				commonItems = append(commonItems, char)
			}
		}

		// If we are on the last line of a trio
		// Then we need to iterate through map
		// and see if we caught any items with three
		// then reset map
		if i%3 == 2 {

			for k, v := range badgeMap {
				if v == 3 {
					badgeTypes = append(badgeTypes, k)
				}
			}
			badgeMap = make(map[rune]int)
		}
	}

	return commonItems, badgeTypes
}

func calcRuneValue(r rune) int {
	if int(r) < 97 { // Capital letters are before lowercase
		return int(r)%65 + 27
	} else {
		return int(r)%97 + 1
	}
}

func run(input string) (int, int) {
	commonItems, badgeTypes := readInput(input)

	s := 0

	for _, item := range commonItems {
		s += calcRuneValue(item)
	}

	r := 0
	for _, badge := range badgeTypes {
		r += calcRuneValue(badge)
	}

	return s, r
}