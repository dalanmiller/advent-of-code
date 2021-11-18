package main

import (
	"fmt"
)

type Rule interface {
	validate(int) bool
}

type LengthRule struct {
	length int
}

func (lr LengthRule) validate(input int) bool {
	return len(fmt.Sprint(input)) <= lr.length
}

type IncreasingRule struct {
}

func (ir IncreasingRule) validate(input int) bool {
	last := -1
	for _, rune := range fmt.Sprint(input) {
		n := int(rune)
		if n < last {
			return false
		}
		last = n
	}

	return true
}

type DoubleAdjacentDigitRule struct {
}

func (dadr DoubleAdjacentDigitRule) validate(input int) bool {
	string_input := fmt.Sprint(input)
	for i, char := range string_input {
		if i == 0 && rune(string_input[i+1]) == char {
			return true
		} else if i == len(string_input)-1 {
			if rune(string_input[i-1]) == char {
				return true
			}
		} else if i < len(string_input) && i > 0 && rune(string_input[i-1]) == char || char == rune(string_input[i+1]) {
			return true
		}
	}

	return false
}

type DoubleAdjacentDigitRuleTwo struct {
}

func (dadr DoubleAdjacentDigitRuleTwo) validate(input int) bool {
	string_input := fmt.Sprint(input)
	runs := make(map[rune]int)
	for i, char := range string_input {
		if i == 0 {
			continue
		}
		_, ok := runs[char]
		if rune(string_input[i-1]) == char {
			if !ok {
				runs[char] = 2
			} else {
				runs[char]++
			}
		}
	}

	// If we have at least two same chars in a row, rule is approved
	for _, v := range runs {
		if v == 2 {
			return true
		}
	}

	return false
}

func validator(input int, rules []Rule) bool {

	for _, rule := range rules {
		if !rule.validate(input) {
			return false
		}
	}

	return true
}

type Range struct {
	Begin int
	End   int
}

func run(r Range, rules []Rule) int {

	var count int
	for i := r.Begin; i <= r.End; i++ {
		if validator(i, rules) {
			count++
		}
	}

	return count
}
