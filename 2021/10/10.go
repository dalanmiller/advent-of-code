package main

import (
	"sort"
	"strings"
)

var syntaxErrorMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocompleteMap = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

const (
	PAREN    = '('
	BRACKET  = '{'
	SQUARE   = '['
	CHEVRON  = '<'
	CPAREN   = ')'
	CBRACKET = '}'
	CSQUARE  = ']'
	CCHEVRON = '>'
)

var openerMap = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
	'>': '<',
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

type Stack struct {
	Stack []rune
}

func (s *Stack) Push(r rune) {
	s.Stack = append(s.Stack, r)
}

func (s *Stack) Pop() rune {
	last := len(s.Stack) - 1
	r := s.Stack[last]
	s.Stack = s.Stack[:last]
	return r
}

func (s *Stack) Peek() rune {
	return s.Stack[len(s.Stack)-1]
}

func (s *Stack) HasNext() bool {
	return len(s.Stack) > 0
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func parseLine(line string, filter bool) (int, []rune) {

	s := Stack{
		Stack: []rune{},
	}

	for _, char := range line {
		switch char {
		case PAREN, BRACKET, SQUARE, CHEVRON:
			s.Push(char)
		case CPAREN, CBRACKET, CSQUARE, CCHEVRON:
			if openerMap[char] != s.Peek() {
				return syntaxErrorMap[char], nil
			} else {
				s.Pop()
			}
		}
	}

	var completion []rune
	for s.HasNext() {
		completion = append(completion, openerMap[s.Pop()])
	}

	return 0, completion
}

func run(input string) int {
	lines := parseInput(input)

	var sum int
	for _, line := range lines {
		v, _ := parseLine(line, false)
		sum += v
	}
	return sum
}

type Score []int

func (s Score) Len() int {
	return len(s)
}

func (s Score) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Score) Less(i, j int) bool {
	return s[i] < s[j]
}

func run_two(input string) int {
	lines := parseInput(input)

	var scores Score

	for _, line := range lines {
		_, completion := parseLine(line, true)

		if completion == nil {
			continue
		}

		var sum int
		for _, char := range completion {
			sum = sum*5 + autocompleteMap[char]
		}

		scores = append(scores, sum)
	}

	sort.Sort(scores)

	return scores[(len(scores) / 2)]
}
