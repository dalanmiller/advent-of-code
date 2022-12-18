package main

import (
	"bufio"
	"container/heap"
	"io"
	"log"
	"math"
	"sort"
)

type coord struct {
	X         int
	Y         int
	Elevation int
	Terminal  bool
	Steps     int
	AdjCoords []*coord
}

type pathQueue []*coord

func (pq *pathQueue) Len() int {
	return len(*pq)
}

func (pq *pathQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq *pathQueue) Push(c any) {
	co := c.(*coord)
	*pq = append(*pq, co)
}

func (pq *pathQueue) Less(i, j int) bool {
	return (*pq)[i].Steps < (*pq)[j].Steps
}

func (pq pathQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func traverse(start, end *coord, eGrid elevationGrid) int {
	visited := map[*coord]bool{start: true}
	start.Steps = 0
	pq := pathQueue{}

	heap.Init(&pq)
	heap.Push(&pq, start)

	for _, row := range eGrid {
		for _, coord := range row {
			if coord == start {
				continue
			}
			heap.Push(&pq, coord)
		}
	}

	for pq.Len() > 0 {
		if _, visitedEnd := visited[end]; visitedEnd {
			break
		}
		// Find min node in the open list
		current := heap.Pop(&pq).(*coord)
		log.Printf("(%d, %d)", current.X, current.Y)

		for _, adjCoord := range current.AdjCoords {
			if _, visited := visited[adjCoord]; !visited {
				if current.Steps+1 < adjCoord.Steps {
					adjCoord.Steps = current.Steps + 1

					for i, coord := range pq {
						if coord == adjCoord {

							// The heap doesn't rebalance so you have to tell it to do so
							heap.Fix(&pq, i)
							break
						}
					}
				}
			}
		}

		visited[current] = true
	}

	// JIC
	return end.Steps
}

func adjacentCoords(i, j, width, height int) []coord {
	fullCoords := []coord{
		// top
		{X: j, Y: i - 1},
		// middle
		{X: j - 1, Y: i},
		{X: j + 1, Y: i},
		// bottom
		{X: j, Y: i + 1},
	}

	validCoords := make([]coord, 0, 4)
	for _, c := range fullCoords {
		if c.X >= 0 && c.X < width && c.Y >= 0 && c.Y < height {
			validCoords = append(validCoords, c)
		}
	}

	return validCoords
}

type elevationGrid [][]*coord

func readInput(input io.Reader) (*coord, *coord, elevationGrid) {

	var start, end *coord
	s := bufio.NewScanner(input)

	m := elevationGrid{}

	i := 0
	for s.Scan() {
		line := s.Text()

		m = append(m, []*coord{})
		for j, r := range line {
			c := &coord{
				X:         j,
				Y:         i,
				Elevation: int(r),
				Terminal:  r == 'E',
				Steps:     math.MaxInt,
			}

			if r == 'E' {
				c.Elevation = int('z')
				end = c
			}

			if r == 'S' {
				c.Elevation = int('a')
				c.Steps = 0
				start = c
			}

			m[i] = append(m[i], c)
		}

		i++
	}

	// Generate coords and then
	// determine if real and then
	// sort by height
	for y, row := range m {
		for x, c := range row {
			coords := adjacentCoords(y, x, len(row), len(m))
			for _, co := range coords {
				adjCoord := m[co.Y][co.X]
				if adjCoord.Elevation > c.Elevation+1 {
					continue
				}

				c.AdjCoords = append(c.AdjCoords, adjCoord)
			}

			sort.Slice(c.AdjCoords, func(i, j int) bool {
				return c.AdjCoords[i].Elevation < c.AdjCoords[j].Elevation
			})
		}
	}

	return start, end, m
}

func run(input io.Reader) int {
	start, end, eGrid := readInput(input)

	steps := traverse(start, end, eGrid)

	return steps
}
