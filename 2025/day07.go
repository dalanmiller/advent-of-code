package year2025

import (
	"bufio"
	"io"
	"strings"
)

// .......S.......
// ...............
// .......^.......
// ...............
// ......^.^......
// ...............
// .....^.^.^.....
// ...............
// ....^.^...^....
// ...............
// ...^.^...^.^...
// ...............
// ..^...^.....^..
// ...............
// .^.^.^.^.^...^.
// ...............

type Stream int

func allDots(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != '.' {
			return false
		}
	}
	return true
}

func Day07PartOne(r io.Reader) int {
	all, _ := io.ReadAll(r)
	lines := strings.Split(string(all), "\n")

	streams := make(map[Stream]struct{})
	split := 0
	for _, line := range lines {
		if allDots(line) {
			continue
		}
		for j, chr := range line {
			switch chr {
			case 'S':
				// streams = append(streams, Stream(j-1), Stream(j+2))
				streams[Stream(j)] = struct{}{}
			case '^':
				if _, ok := streams[Stream(j)]; ok {
					split++
					streams[Stream(j-1)] = struct{}{}
					streams[Stream(j+1)] = struct{}{}
					delete(streams, Stream(j))
				}
			case '.':
				continue
			}
		}
	}

	return split
}

func clear(s []int) {
	for i := range s {
		s[i] = 0
	}
}

func Day07PartTwo(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var streams, next []int

	width := -1
	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		if allDots(line) || line == "" {
			continue
		}

		if width == -1 {
			width = len(line)
			streams = make([]int, width)
			next = make([]int, width)
		}

		if row == 0 {
			streams[Stream(strings.Index(line, "S"))] = 1
			row++
			continue
		}

		clear(next)

		for j := 0; j < width; j++ {

			worlds := streams[j]
			if worlds == 0 {
				continue
			}

			switch line[j] {
			case '^':
				next[j-1] += worlds
				next[j+1] += worlds
			default:
				next[j] += worlds
			}
		}

		streams, next = next, streams
		row++
	}

	timelines := 0
	for _, v := range streams {
		timelines += v
	}

	return timelines
}
