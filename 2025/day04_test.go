package year2025

import (
	"strings"
	"testing"
)

const day04Example = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

const day04Example01 = `@@@
@@@
@@@`

func TestDay04PartOneCheckOne(t *testing.T) {
	AssertEqual(t, Day04PartOne, strings.NewReader(day04Example01), 4)
}

const day04Example02 = `@@@
@@.
@.@`

func TestDay04PartOneCheckTwo(t *testing.T) {
	AssertEqual(t, Day04PartOne, strings.NewReader(day04Example02), 4)
}

func TestDay04PartOneExample(t *testing.T) {
	AssertEqual(t, Day04PartOne, strings.NewReader(day04Example), 13)
}

func TestDay04PartOneInput(t *testing.T) {
	AssertEqual(t, Day04PartOne, ReaderForInput(4), 1553)
}

func TestDay04PartTwoExample(t *testing.T) {
	AssertEqual(t, Day04PartTwo, strings.NewReader(day04Example), 43)
}

func TestDay04PartTwoInput(t *testing.T) {
	AssertEqual(t, Day04PartTwo, ReaderForInput(4), 8442)
}
