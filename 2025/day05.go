package year2025

import (
	"io"
	"slices"
	"strconv"
	"strings"
)

type FreshRange struct{ Begin, End int }

type FreshRanges []FreshRange

func (fr *FreshRange) InRange(i int) bool {
	return (i >= fr.Begin && i <= fr.End)
}

func (frs *FreshRanges) Check(i int) bool {
	for _, fr := range *frs {
		if fr.InRange(i) {
			return true
		}
	}
	return false
}

func Day05PartOne(r io.Reader) int {
	all, _ := io.ReadAll(r)
	lines := strings.Split(string(all), "\n")

	ranges := FreshRanges{}
	count := 0

	for _, line := range lines {
		if strings.Contains(line, "-") {
			split := strings.Split(line, "-")
			begin, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, FreshRange{begin, end})
			continue
		}

		if line == "" {
			continue
		}

		d, _ := strconv.Atoi(line)
		if ranges.Check(d) {
			count++
		}

	}

	return count
}

func Day05PartTwo(r io.Reader) int {
	all, _ := io.ReadAll(r)
	lines := strings.Split(string(all), "\n")

	ranges := FreshRanges{}
	length := 0

	for _, line := range lines {
		if strings.Contains(line, "-") {
			split := strings.Split(line, "-")
			begin, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, FreshRange{begin, end})
			length += end - begin
		}
	}

	// Sort all the ranges based on Begin and then on End
	slices.SortFunc(ranges, func(a, b FreshRange) int {
		if a.Begin != b.Begin {
			return a.Begin - b.Begin
		}
		return a.End - b.End
	})

	// generate slice of where we will put merged ranges
	out := make(FreshRanges, 0, len(ranges))
	cur := ranges[0]

	for _, r := range ranges[1:] {
		if r.Begin <= cur.End+1 {
			if r.End > cur.End {
				cur.End = r.End // We just push out the end of current to whatever we currently have.
			}
			continue
		}
		out = append(out, cur)
		cur = r
	}
	out = append(out, cur)

	total := 0
	for _, s := range out {
		total += (s.End - s.Begin) + 1
	}

	return total
}
