package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Outcome int

const (
	LESS Outcome = iota - 1
	EQUAL
	GREATER
)

type DataType string

const (
	LIST    DataType = "LIST"
	INTEGER DataType = "INTEGER"
)

type pair struct {
	Left  *packetData
	Right *packetData
	// Raw   string
}

type packetData struct {
	DataType DataType
	List     []*packetData
	Integer  int
}

func (p pair) Compare() (order bool) {

	// 1. If both values are integers, the lowest one should come first. If otherwise, the values
	// . are not in the right order. If same, continue.
	// 2. If both values are lists, compare the first value of each list, then the second value, and so on.
	// . If the left list runs out of items first, then the inputs are in the right order
	// . If right runs out first, then not right order.
	// . If same length -- continue.
	// 3. If exactly one value is an integer, convert the integer to a list which contains
	// . that integer as its only value then retry the comparison.

	var left, right *packetData
	// log.Println(p.Raw)

	for i := 0; i < len(p.Left.List); i++ {

		if i > len(p.Right.List)-1 {
			return false
		}

		left = p.Left.List[i]
		right = p.Right.List[i]
		switch {

		case left.DataType == INTEGER && right.DataType == INTEGER:
			switch {
			case left.Integer < right.Integer:
				return true
			case left.Integer == right.Integer:
				continue
			case left.Integer > right.Integer:
				return false
			}

		case left.DataType == LIST && right.DataType == LIST:
			switch compareLists(left, right) {
			case LESS:
				return true
			case EQUAL:
				continue
			case GREATER:
				return false
			}

		case (left.DataType == LIST && right.DataType == INTEGER) || (left.DataType == INTEGER && right.DataType == LIST):
			if left.DataType == INTEGER {
				left = &packetData{
					DataType: LIST,
					List:     []*packetData{{DataType: INTEGER, Integer: left.Integer}},
				}
			} else {
				right = &packetData{
					DataType: LIST,
					List:     []*packetData{{DataType: INTEGER, Integer: right.Integer}},
				}
			}

			switch compareLists(left, right) {
			case LESS:
				return true
			case EQUAL:
				continue
			case GREATER:
				return false
			}
		}
	}

	if len(p.Left.List) < len(p.Right.List) {
		return true
	} else if len(p.Left.List) > len(p.Right.List) {
		return false
	}

	return order
}

func compareLists(a *packetData, b *packetData) Outcome {
	left, right := *a, *b

	for i := range left.List {
		if i > len(left.List)-1 {
			return LESS
		}

		if i > len(right.List)-1 {
			return GREATER
		}

		// log.Println(left.List[i], right.List[i])

		if left.List[i].DataType == LIST && right.List[i].DataType == LIST {
			switch compareLists(left.List[i], right.List[i]) {
			case LESS:
				return LESS
			case EQUAL:
				continue
			case GREATER:
				return GREATER
			}
		}

		if left.List[i].Integer < right.List[i].Integer {
			return LESS
		}
		if left.List[i].Integer > right.List[i].Integer {
			return GREATER
		}
	}

	return EQUAL
}

func readInput(input io.Reader) (pairs []pair) {

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()

		pair := pair{
			Left:  parsePacketString(first),
			Right: parsePacketString(second),
			// Raw:   fmt.Sprintf("%s | %s", first, second),
		}
		pairs = append(pairs, pair)

		// Jump blank line
		scanner.Scan()
	}

	return pairs
}

func parsePacketString(p string) *packetData {
	// [1,[2,[3,[4,[5,6,7]]]],8,9]

	s := bufio.NewScanner(strings.NewReader(strings.TrimLeft(p, " ")))
	s.Split(bufio.ScanRunes)

	var current *packetData
	depth := []*packetData{}

	// first = current
	for s.Scan() {
		r := []rune(s.Text())[0]
		switch {
		case r == '[':
			current = &packetData{
				DataType: LIST,
				List:     []*packetData{},
			}

			// Make sure that we update the last layer
			if len(depth) > 0 {
				depth[len(depth)-1].List = append(depth[len(depth)-1].List, current)
			}

			depth = append(depth, current)
		case r == ']':
			if len(depth) > 1 {
				depth = depth[:len(depth)-1]
			}
			current = depth[len(depth)-1]
		case unicode.IsDigit(r):
			v, _ := strconv.Atoi(string(r))
			current.List = append(current.List, &packetData{
				DataType: INTEGER,
				Integer:  v,
			})
		}
	}

	return depth[0]
}

func run(input io.Reader) int {
	pairs := readInput(input)

	s := 0
	for i, p := range pairs {
		if p.Compare() {
			s += i + 1
		}
	}

	return s
}
