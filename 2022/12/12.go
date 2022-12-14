package main

import (
	"bufio"
	"container/heap"
	"io"
)

type coord struct {
	X         int
	Y         int
	Elevation int
}

type PathQueue []*coord

func (pq *PathQueue) Len() int {
	return len(*pq)
}

func (pq *PathQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq *PathQueue) Push(c any) {
	co := c.(*coord)
	*pq = append(*pq, co)
}

func (pq *PathQueue) Less(i, j int) bool {
	return (*pq)[i].Elevation > (*pq)[j].Elevation
}

func (pq PathQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func aStar(start *coord) int {

	pq := PathQueue{}
	heap.Init(&pq)
	heap.Push(&pq, start)

	start.PathRisk = 0
	cameFrom := make(map[coord]coord)
	gScore := map[coord]int{
		start.Coord: 0,
	}

	for pq.Len() > 0 {

		// Find min node in the open list
		current := heap.Pop(&pq).(*coord)

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

				heap.Push(&pq, adjNode)
			}
		}
	}

	// JIC
	return -1
}

type elevationMap [][]rune

func readInput(input io.Reader) elevationMap {
	s := bufio.NewScanner(input)

	m := elevationMap{}

	i := 0
	for s.Scan() {
		line := s.Text()

		m = append(m, []rune{})
		for _, r := range line {
			m[i] = append(m[i], r)
		}

		i++
	}

	return m
}

func run(input io.Reader) int {
	eMap := readInput(input)

	return moves
}
