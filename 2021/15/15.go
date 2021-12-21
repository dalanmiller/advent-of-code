package main

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Grid [][]*Node

type Node struct {
	Coord    Coord
	Risk     int
	PathRisk int
	Adj      []*Node
	Terminal bool
}

type Coord struct {
	X int
	Y int
}

type LowRiskQueue []*Node

func (lrq LowRiskQueue) Len() int {
	return len(lrq)
}

func (lrq LowRiskQueue) Less(i, j int) bool {
	return lrq[i].PathRisk < lrq[j].PathRisk
}

func (lrq LowRiskQueue) Swap(i, j int) {
	lrq[i], lrq[j] = lrq[j], lrq[i]
}

func (lrq *LowRiskQueue) Push(x interface{}) {
	node := x.(*Node)
	*lrq = append(*lrq, node)
}

func (lrq *LowRiskQueue) Pop() interface{} {
	old := *lrq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*lrq = old[0 : n-1]
	return item
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

// func retracePath(m map[Coord]Coord) {
// 	ok := true
// 	var c Coord
// 	oc := Coord{X: 0, Y: 0}
// 	for ok {
// 		c, ok = m[oc]
// 		log.Printf("(%d, %d) => (%d, %d)", oc.X, oc.Y, c.X, c.Y)
// 		oc = c
// 	}
// }

func aStar(start *Node) int {

	lrq := LowRiskQueue{}
	heap.Init(&lrq)
	heap.Push(&lrq, start)

	start.PathRisk = 0
	cameFrom := make(map[Coord]Coord)
	gScore := map[Coord]int{
		start.Coord: 0,
	}

	for lrq.Len() > 0 {

		// Find min node in the open list
		current := heap.Pop(&lrq).(*Node)

		for _, adjNode := range current.Adj {

			if adjNode.Terminal {
				v := gScore[current.Coord]
				// retracePath(cameFrom)
				return v + adjNode.Risk
			}

			tentativeRisk := gScore[current.Coord] + adjNode.Risk
			_, ok := gScore[adjNode.Coord]
			// !ok maps to the default value of nil and thus
			// the tentative score is better
			if !ok || tentativeRisk < gScore[adjNode.Coord] {
				cameFrom[current.Coord] = adjNode.Coord
				gScore[adjNode.Coord] = tentativeRisk
				adjNode.PathRisk = tentativeRisk

				heap.Push(&lrq, adjNode)
			}
		}
	}

	// JIC
	return -1
}

func parseInput(input string, sizeMultiplier int) Grid {
	lines := strings.Split(input, "\n")

	caveMap := make(Grid, len(lines)*sizeMultiplier)

	for i := 0; i < len(lines)*sizeMultiplier; i++ {
		line := lines[i%len(lines)]
		split := strings.Split(line, "")

		row := []*Node{}
		yAddAmount := int(math.Floor(float64(i) / float64(len(split))))
		for j := 0; j < len(split)*sizeMultiplier; j++ {
			xAddAmount := int(math.Floor(float64(j) / float64(len(split))))

			str := split[j%len(split)]
			n, _ := strconv.Atoi(str)
			v := n + xAddAmount + yAddAmount
			if v > 9 {
				v %= 9
			}
			row = append(row, &Node{
				Coord: Coord{X: j, Y: i},
				Risk:  v,
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
		}
	}

	caveMap[len(caveMap)-1][len(caveMap[0])-1].Terminal = true

	return caveMap
}

func run(input string, sizeMultiplier int) int {
	caveMap := parseInput(input, sizeMultiplier)

	n := aStar(caveMap[0][0])

	return n
}
