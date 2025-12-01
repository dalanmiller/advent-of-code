package year2025

import (
	"strings"
	"testing"
)

const day01Example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestDay01PartOneExample(t *testing.T) {
	AssertEqual(t, Day01PartOne, strings.NewReader(day01Example), 3)
}

func TestDay01PartOneInput(t *testing.T) {
	AssertEqual(t, Day01PartOne, ReaderForInput(1), 1191)
}

func TestDay01PartTwoExample(t *testing.T) {
	AssertEqual(t, Day01PartTwo, strings.NewReader(day01Example), 6)
}

// 6543 is not correct (too low)
// 7671 is not correct
// 5722 is not correct
func TestDay01PartTwoInput(t *testing.T) {
	AssertEqual(t, Day01PartTwo, ReaderForInput(1), 6858)
}
