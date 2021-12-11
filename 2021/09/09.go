package main

import (
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	N          int
	Sink       bool
	NextLowest *Point
	FinalSink  *Point
	Up         *Point
	Down       *Point
	Left       *Point
	Right      *Point
}

func descendToSink(p *Point) *Point {
	if p.Sink || p.NextLowest == nil {
		return p
	} else if p.FinalSink != nil {
		return p.FinalSink
	} else {
		return descendToSink(p.NextLowest)
	}
}

func parseInput(input string) [][]*Point {
	rows := strings.Split(input, "\n")
	height := len(rows)
	width := len(rows[0])

	grid := make([][]*Point, height)
	for i := range grid {
		grid[i] = make([]*Point, width)
	}

	for i, row := range rows {
		// rows[i] = make([]*Point, len(row))
		for j, char := range row {
			n, _ := strconv.Atoi(string(char))
			grid[i][j] = &Point{
				N:    n,
				Sink: false,
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cur := grid[i][j]
			if i > 0 && j > 0 && i < height-1 && j < width-1 {
				// inner cells
				cur.Up = grid[i-1][j]
				cur.Right = grid[i][j+1]
				cur.Down = grid[i+1][j]
				cur.Left = grid[i][j-1]
			} else if i-1 == -1 {
				// top row

				if j-1 == -1 {
					// Top left corner
					cur.Right = grid[i][j+1]
					cur.Down = grid[i+1][j]
				} else if j+1 == width {
					// Top right corner
					cur.Left = grid[i][j-1]
					cur.Down = grid[i+1][j]
				} else {
					// Inner top row
					cur.Right = grid[i][j+1]
					cur.Down = grid[i+1][j]
					cur.Left = grid[i][j-1]
				}

			} else if j+1 == width && i+1 < height {
				// right edge
				cur.Up = grid[i-1][j]
				cur.Down = grid[i+1][j]
				cur.Left = grid[i][j-1]
			} else if i+1 == height {
				//bottom edge

				if j-1 == -1 {
					// Bottom left corner
					cur.Right = grid[i][j+1]
					cur.Up = grid[i-1][j]
				} else if j+1 == width {
					// Bottom right corner
					cur.Left = grid[i][j-1]
					cur.Up = grid[i-1][j]
				} else {
					// Inner bottom row
					cur.Up = grid[i-1][j]
					cur.Right = grid[i][j+1]
					cur.Left = grid[i][j-1]
				}

			} else if j-1 == -1 {
				//left edge
				cur.Up = grid[i-1][j]
				cur.Right = grid[i][j+1]
				cur.Down = grid[i+1][j]
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cur := grid[i][j]
			allAdjLarger := true
			lowest := 9
			for _, adj := range []*Point{cur.Up, cur.Right, cur.Down, cur.Left} {
				if adj != nil {
					if cur.N >= adj.N {
						allAdjLarger = false
					}

					if adj.N < lowest && adj.N < cur.N {
						cur.NextLowest = adj
						lowest = adj.N
					}
				}
			}

			cur.Sink = allAdjLarger
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cur := grid[i][j]

			if cur.NextLowest != nil && cur.N != 9 {
				cur.FinalSink = descendToSink(cur)
			} else {
				cur.FinalSink = nil
			}
		}
	}

	return grid
}

type basins []int

func (b basins) Len() int {
	return len(b)
}

func (b basins) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b basins) Less(i, j int) bool {
	return b[i] > b[j]
}

func run(input string) (int, int) {
	grid := parseInput(input)
	height, width := len(grid), len(grid[0])

	var sumPartOne int
	basinSize := make(map[*Point]int)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j].Sink {
				sumPartOne += 1 + grid[i][j].N
			}

			if grid[i][j].FinalSink != nil {
				basinSize[grid[i][j].FinalSink]++
			}
		}
	}

	var basins basins
	for _, v := range basinSize {
		basins = append(basins, v+1)
	}

	sort.Sort(basins)

	return sumPartOne, basins[0] * basins[1] * basins[2]
}
