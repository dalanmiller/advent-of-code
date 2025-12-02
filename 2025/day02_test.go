package year2025

import (
	"strings"
	"testing"
)

const day02Example = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func TestDay02PartOneExample(t *testing.T) {
	AssertEqual(t, Day02PartOne, strings.NewReader(day02Example), 1227775554)
}

// 15873079070 (too low)
// 4395990980 (too low)
// 15873079081
func TestDay02PartOneInput(t *testing.T) {
	AssertEqual(t, Day02PartOne, ReaderForInput(2), 15873079081)
}

func TestDay02PartTwoExample(t *testing.T) {
	AssertEqual(t, Day02PartTwo, strings.NewReader(day02Example), 4174379265)
}

func TestDay02PartTwoInput(t *testing.T) {
	AssertEqual(t, Day02PartTwo, ReaderForInput(2), 22617871034)
}
