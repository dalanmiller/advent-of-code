package year2025

import (
	"io"
	"strings"
)

type Position [2]int

type P struct {
	X     int
	Y     int
	Paper bool
}

func (p *Position) GetAdjacent() []Position {
	var adj []Position

	mods := [8][2]int{{-1, 0}, {-1, 1}, {-1, -1}, {0, 1}, {0, -1}, {1, 0}, {1, 1}, {1, -1}}

	for _, m := range mods {
		adj = append(adj, Position{p[0] + m[0], p[1] + m[1]})
	}

	return adj
}

func Day04PartOne(r io.Reader) int {
	all, _ := io.ReadAll(r)
	lines := strings.Split(string(all), "\n")

	papers := map[Position]*P{}

	for x, line := range lines {
		for y, chr := range line {
			switch chr {
			case '.':
				continue
			case '@':
				c := Position{x, y}
				p := P{x, y, true}
				papers[c] = &p
			}
		}
	}

	count := 0
	for pos := range papers {
		adj := pos.GetAdjacent()

		paperAround := 0
		for _, a := range adj {

			// Make sure we find those adjacents in the positionMap
			if p, ok := papers[a]; ok {
				// If any of the adjacent positions have Paper there, then we know it's not one of the chosen ones
				// As well, this handles whether or not there are 2 - 4 positions adjacent with the rest being
				// "off board"
				if p.Paper {
					// log.Printf("%d -> %d, %d: %v", pos, p.X, p.Y, p.Paper)
					paperAround++
				}
			}
		}

		if paperAround < 4 {
			count++
		}
	}

	return count
}

func copyMap(src map[Position]*P) map[Position]*P {
	new := make(map[Position]*P, len(src))
	for k, v := range src {
		new[k] = v
	}
	return new
}

func Day04PartTwo(r io.Reader) int {
	all, _ := io.ReadAll(r)
	lines := strings.Split(string(all), "\n")

	papers := map[Position]*P{}

	for x, line := range lines {
		for y, chr := range line {
			switch chr {
			case '.':
				continue
			case '@':
				c := Position{x, y}
				p := P{x, y, true}
				papers[c] = &p
			}
		}
	}

	removalCount := 0
	count := -1

	for count != 0 {

		count = 0
		tmpPapers := copyMap(papers)
		papers = make(map[Position]*P, len(tmpPapers))

		for pos := range tmpPapers {
			adj := pos.GetAdjacent()

			paperAround := 0
			for _, a := range adj {
				if p, ok := tmpPapers[a]; ok {
					if p.Paper {
						paperAround++
					}
				}
			}

			if paperAround < 4 {
				count++
			} else {
				papers[pos] = tmpPapers[pos]
			}
		}
		removalCount += count
	}

	return removalCount
}
