package year2025

import (
	"strings"
	"testing"
)

const day08Example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestDay08PartOneExample(t *testing.T) {
	AssertEqual8(t, Day08PartOne, strings.NewReader(day08Example), 10, 40)
}

func TestDay08PartOneInput(t *testing.T) {
	AssertEqual8(t, Day08PartOne, ReaderForInput(8), 1000, 102816)
}

func TestDay08PartTwoExample(t *testing.T) {
	AssertEqual(t, Day08PartTwo, strings.NewReader(day08Example), 25272)
}

func TestDay08PartTwoInput(t *testing.T) {
	AssertEqual(t, Day08PartTwo, ReaderForInput(8), 100011612)
}
