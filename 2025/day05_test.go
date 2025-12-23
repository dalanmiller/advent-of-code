package year2025

import (
	"strings"
	"testing"
)

const day05Example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestDay05PartOneExample(t *testing.T) {
	AssertEqual(t, Day05PartOne, strings.NewReader(day05Example), 3)
}

func TestDay05PartOneInput(t *testing.T) {
	AssertEqual(t, Day05PartOne, ReaderForInput(5), 707)
}

func TestDay05PartTwoExample(t *testing.T) {
	AssertEqual(t, Day05PartTwo, strings.NewReader(day05Example), 14)
}

func TestDay05PartTwoInput(t *testing.T) {
	AssertEqual(t, Day05PartTwo, ReaderForInput(5), 361615643045059)
}
