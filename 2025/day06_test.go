package year2025

import (
	"strings"
	"testing"
)

const day06Example = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func TestDay06PartOneExample(t *testing.T) {
	AssertEqual(t, Day06PartOne, strings.NewReader(day06Example), 4277556)
}

func TestDay06PartOneInput(t *testing.T) {
	AssertEqual(t, Day06PartOne, ReaderForInput(6), 5552221122013)
}

func TestDay06PartTwoExample(t *testing.T) {
	AssertEqual(t, Day06PartTwo, strings.NewReader(day06Example), 3263827)
}

func TestDay06PartTwoInput(t *testing.T) {
	AssertEqual(t, Day06PartTwo, ReaderForInput(6), 11371597126232)
}
