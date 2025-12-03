package year2025

import (
	"strings"
	"testing"
)

const day03Example = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestDay03PartOneExample(t *testing.T) {
	AssertEqual(t, Day03PartOne, strings.NewReader(day03Example), 357)
}

func TestDay03PartOneInput(t *testing.T) {
	AssertEqual(t, Day03PartOne, ReaderForInput(3), 17179)
}

func TestDay03PartTwoExample(t *testing.T) {
	AssertEqual64(t, Day03PartTwo, strings.NewReader(day03Example), 3121910778619)
}

// 17835271775054454722 (too high)
func TestDay03PartTwoInput(t *testing.T) {
	AssertEqual64(t, Day03PartTwo, ReaderForInput(3), 170025781683941)
}
