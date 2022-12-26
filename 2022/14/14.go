package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
}

type outcome int

const (
	CONTINUE outcome = iota
	STOP
)

type rockMap map[point]bool

func (r *rockMap) fall(p point, maxY int, part int) outcome {

	finalPoint := p
	var tempPoint, newPoint point

	for {

		tempPoint = finalPoint
		// Iterate through directly below, down and to left, down and to right
		for _, i := range [][2]int{{0, 1}, {-1, 1}, {1, 1}} {
			newPoint = point{finalPoint.X + i[0], finalPoint.Y + i[1]}

			// Part 2: need to check if we've hit the floor
			if part == 2 && newPoint.Y == maxY+2 {
				break
			}

			// Check if something is at the point it's trying to go to
			if _, ok := (*r)[newPoint]; !ok {
				finalPoint = newPoint
				break
			}
		}

		// Check if we are headed towards the void
		if part == 1 && finalPoint.Y > maxY {
			return STOP
		}

		if part == 2 && finalPoint.X == 500 && finalPoint.Y == 0 {
			return STOP
		}

		// Check if we moved if not, break
		if tempPoint == finalPoint {
			break
		}
	}

	(*r)[finalPoint] = true
	return CONTINUE
}

func readInput(input io.Reader) (rockMap, int) {
	s := bufio.NewScanner(input)

	rockMap := rockMap{}
	maxY := 0
	for s.Scan() {
		coords := strings.Split(s.Text(), " -> ")
		points := []point{}

		for _, coord := range coords {
			rockPoint := strings.Split(coord, ",")
			x, _ := strconv.Atoi(rockPoint[0])
			y, _ := strconv.Atoi(rockPoint[1])
			points = append(points, point{x, y})
		}

		morePoints := []point{}
		for i := 0; i+1 <= len(points)-1; i++ {

			var n int
			// If it's on the same Y axis
			if points[i].X == points[i+1].X {
				if points[i+1].Y < points[i].Y {
					n = points[i].Y - 1
				} else {
					n = points[i].Y + 1
				}

				for n != points[i+1].Y {
					morePoints = append(morePoints, point{
						points[i].X, n,
					})

					if points[i+1].Y < n {
						n--
					} else {
						n++
					}
				}
			} else { // else we are on the same X axis
				if points[i+1].X < points[i].X {
					n = points[i].X - 1
				} else {
					n = points[i].X + 1
				}
				for n != points[i+1].X {
					morePoints = append(morePoints, point{
						n, points[i].Y,
					})

					if points[i+1].X < n {
						n--
					} else {
						n++
					}
				}
			}
		}

		points = append(points, morePoints...)

		for _, point := range points {
			rockMap[point] = true

			if point.Y > maxY {
				maxY = point.Y
			}
		}
	}

	return rockMap, maxY
}

func run(input io.Reader, part int) int {
	rockMap, maxY := readInput(input)

	n := 0
	for rockMap.fall(point{500, 0}, maxY, part) == CONTINUE {
		n++
	}

	return n + (part - 1)
}
