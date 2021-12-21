package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type Grid [][]*Node

type Node struct {
	Coord Coord
	Risk  int
	// PathRisk int
	Adj []*Node

	Terminal bool
	// Parent   *Node
}

type Coord struct {
	X int
	Y int
}

func adjacentCoords(i, j, width, height int) []Coord {
	var fullCoords []Coord

	fullCoords = append(fullCoords, []Coord{
		// top
		{X: j, Y: i - 1},
		// middle
		{X: j - 1, Y: i},
		{X: j + 1, Y: i},
		// bottom
		{X: j, Y: i + 1},
	}...)

	var validCoords []Coord
	for _, coord := range fullCoords {
		if coord.X >= 0 && coord.X < width && coord.Y >= 0 && coord.Y < height {
			validCoords = append(validCoords, coord)
		}
	}

	return validCoords
}

func retracePath(m map[Coord]Coord) {
	ok := true
	var c Coord
	oc := Coord{X: 0, Y: 0}
	for ok {
		c, ok = m[oc]
		log.Printf("(%d, %d) => (%d, %d)", oc.X, oc.Y, c.X, c.Y)
		oc = c
	}
}

func aStar(start *Node) int {

	open := []*Node{start}
	openContents := map[Coord]bool{
		start.Coord: true,
	}
	cameFrom := make(map[Coord]Coord)
	gScore := map[Coord]int{
		start.Coord: 0,
	}

	for len(open) > 0 {

		// Find min node in the open list
		current := open[0]

		// Pop min node from open list
		open = open[1:]
		delete(openContents, current.Coord)
		cx := Coord{2, 2}
		if current.Coord == cx {
			log.Printf("Welp")
		}

		for _, adjNode := range current.Adj {

			cy := Coord{4, 2}
			cyy := Coord{3, 2}
			if adjNode.Coord == cy || current.Coord == cyy {
				log.Printf("Welp2")
			}
			if adjNode.Terminal {
				v := gScore[current.Coord]
				retracePath(cameFrom)
				return v + adjNode.Risk
			}

			tentativeRisk := gScore[current.Coord] + adjNode.Risk
			_, ok := gScore[adjNode.Coord]
			// !ok maps to the default value of nil and thus
			// the tentative score is better
			if !ok || tentativeRisk < gScore[adjNode.Coord] {
				cameFrom[current.Coord] = adjNode.Coord
				gScore[adjNode.Coord] = tentativeRisk
				if _, ok := openContents[adjNode.Coord]; !ok {
					open = append(open, adjNode)
					openContents[adjNode.Coord] = true

					// Overkill to sort with one entry appended but hey
					sort.Slice(open, func(a, b int) bool {
						return gScore[open[a].Coord] < gScore[open[b].Coord]
					})
				}
			}
		}
	}

	// JIC
	return -1
}

func parseInput(input string, sizeMultiplier int) Grid {
	lines := strings.Split(input, "\n")

	caveMap := make(Grid, len(lines)*sizeMultiplier)

	for i, line := range lines {
		split := strings.Split(line, "")

		row := []*Node{}
		for j, str := range split {
			n, _ := strconv.Atoi(str)

			row = append(row, &Node{
				Coord: Coord{X: j, Y: i},
				Risk:  n,
			})
		}
		caveMap[i] = row
	}

	for i, row := range caveMap {
		for j, node := range row {
			coords := adjacentCoords(i, j, len(row), len(caveMap))
			for _, c := range coords {
				adjNode := caveMap[c.Y][c.X]
				node.Adj = append(node.Adj, adjNode)
			}

			sort.Slice(node.Adj, func(i, j int) bool {
				return node.Adj[i].Risk < node.Adj[j].Risk
			})

			if i == len(caveMap)-1 && j == len(row)-1 {
				node.Terminal = true
			} else {
				node.Terminal = false
			}
		}
	}

	return caveMap
}

func run(input string, sizeMultiplier int) int {
	caveMap := parseInput(input, sizeMultiplier)

	n := aStar(caveMap[0][0])

	return n
}
