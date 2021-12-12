package main

import (
	"strings"
	"unicode"
)

type Cave struct {
	Size        rune
	Name        string
	Connections []*Cave
}

type Path []*Cave

func parseInput(input string, caveMap map[string]*Cave) *Cave {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		system := strings.Split(line, "-")
		s := system[0]
		e := system[1]

		sSize := unicode.IsUpper(rune(s[0]))
		eSize := unicode.IsUpper(rune(e[0]))
		var sSizeRune, eSizeRune rune

		if sSize {
			sSizeRune = 'L'
		} else {
			sSizeRune = 'S'
		}

		if eSize {
			eSizeRune = 'L'
		} else {
			eSizeRune = 'S'
		}

		var start, end *Cave
		ev, ev_ok := caveMap[e]
		sv, sv_ok := caveMap[s]

		if !ev_ok {
			end = &Cave{
				Size:        eSizeRune,
				Name:        e,
				Connections: []*Cave{},
			}
			caveMap[e] = end
		} else {
			end = ev
		}

		if !sv_ok {
			start = &Cave{
				Size:        sSizeRune,
				Name:        s,
				Connections: []*Cave{},
			}
			caveMap[s] = start
		} else {
			start = sv
		}

		start.Connections = append(start.Connections, end)
		end.Connections = append(end.Connections, start)

	}

	return caveMap["start"]
}

func spelunk(c Cave, p Path, ps *[]Path, visitsAllowed int) {
	// If previous path was also a small cave, we don't
	// need to continue this recursion

	count := make(map[string]int, len(p))
	for _, visitedCave := range p {

		// Keep to just little guys
		if visitedCave.Size == 'L' {
			continue
		}

		_, ok := count[visitedCave.Name]
		if !ok {
			count[visitedCave.Name] = 1
		} else {
			count[visitedCave.Name]++
		}
	}

	visitGate := false
	for k, v := range count {
		if k == "start" && v > 1 || k == "end" && v > 1 {
			return
		}

		if visitsAllowed > 1 && v == visitsAllowed {
			if !visitGate {
				visitGate = true
			} else {
				return
			}
		}

		if v > visitsAllowed {
			return
		}
	}

	// Append current cave to Path
	// var newPath Path
	p = append(p, &c)

	// If we are at the end, time to add to
	//  path slice and end
	if c.Name == "end" {
		*ps = append(*ps, append(Path(nil), p...))
		return
	}

	// Lastly, iterate through all connections and start
	// recursions
	for _, cave := range c.Connections {
		spelunk(*cave, p, ps, visitsAllowed)
	}
}

func run(input string, visitsAllowed int) int {

	caveMap := map[string]*Cave{}

	startCave := parseInput(input, caveMap)

	// Valid paths to end
	var validPaths []Path

	spelunk(*startCave, Path{}, &validPaths, visitsAllowed)

	return len(validPaths)
}
