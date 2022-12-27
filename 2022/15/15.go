package main

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strconv"
)

var coordRegex = regexp.MustCompile(`x=(-?\d+), y=(-?\d+)`)

type sensor struct {
	X, Y      int
	Beacon    point
	Manhattan float64
}

type point struct {
	X, Y int
}

func (s *sensor) Within(p point) bool {
	return math.Abs(float64(s.X-p.X))+math.Abs(float64(s.Y-p.Y)) <= s.Manhattan
}

func readInput(input io.Reader, y int) ([]sensor, map[point]bool, map[point]bool, map[point]bool) {

	s := bufio.NewScanner(input)

	sensors := []sensor{}
	pointsInRow := make(map[point]bool)
	bMap := make(map[point]bool)
	sMap := make(map[point]bool)
	var sX, sY, bX, bY int
	var bPoint point
	var sens sensor

	for s.Scan() {
		matches := coordRegex.FindAllStringSubmatch(s.Text(), -1)

		sX, _ = strconv.Atoi(matches[0][1])
		sY, _ = strconv.Atoi(matches[0][2])

		bX, _ = strconv.Atoi(matches[1][1])
		bY, _ = strconv.Atoi(matches[1][2])

		bPoint = point{
			bX, bY,
		}
		sens = sensor{
			sX,
			sY,
			bPoint,
			math.Abs(float64(sX-bX)) + math.Abs(float64(sY-bY)),
		}

		bMap[bPoint] = true
		sMap[point{sX, sY}] = true

		sensors = append(sensors, sens)

		// Experiment

		// First try to determine if sensor + manhattan is even going to cover
		// the target row
		if y == 0 {
			continue
		}

		if sens.Y-int(sens.Manhattan) <= y && sens.Y+int(sens.Manhattan) >= y {

			// Determine starting point
			p := point{
				sens.X - int(sens.Manhattan), y,
			}

			// Determine end point
			endP := point{sens.X + int(sens.Manhattan), y}

			// Until starting point == endinng point we skim across row
			// . attempting to find where row intersects the manhat distance
			// . of the sensor
			for p != endP {
				if sens.Within(p) {
					pointsInRow[p] = true
				}
				p.X++
			}
		}

	}

	return sensors, bMap, sMap, pointsInRow
}

func runPartOne(input io.Reader, y int) int {
	_, beaconMap, sensorMap, pointsInRow := readInput(input, y)

	cleanPointsInRow := make(map[point]bool, len(pointsInRow))

	for point := range pointsInRow {
		_, bOk := beaconMap[point]
		_, sOk := sensorMap[point]

		if !bOk && !sOk {
			cleanPointsInRow[point] = true
		}
	}

	return len(cleanPointsInRow)
}

func (s *sensor) generatePointEdge() []point {
	d := int(s.Manhattan) + 1

	north := point{s.X, s.Y - d}
	east := point{s.X + d, s.Y}
	south := point{s.X, s.Y + d}
	west := point{s.X - d, s.Y}

	pointEdge := []point{north, east, south, west}

	// north to east
	p := north
	for p != east {
		p.X++
		p.Y++
		pointEdge = append(pointEdge, p)
	}

	// east to south
	for p != south {
		p.X--
		p.Y++
		pointEdge = append(pointEdge, p)
	}

	// south to west
	for p != west {
		p.X--
		p.Y--
		pointEdge = append(pointEdge, p)
	}

	// west to north
	for p != north {
		p.X++
		p.Y--
		pointEdge = append(pointEdge, p)
	}

	return pointEdge
}

func runPartTwo(input io.Reader, min int, max int) int {
	sensors, _, _, _ := readInput(input, 0)

	// Need to iterate through each sensor area + 1
	// . if the point is not in any other sensor area
	// . distance, then I think we've found the match

	for _, sensor := range sensors {
		points := sensor.generatePointEdge()

		for _, point := range points {
			if point.X > max || point.Y > max || point.X < min || point.Y < min {
				continue
			}

			winner := true
			for _, sensor := range sensors {
				if sensor.Within(point) {
					winner = false
					break
				}
			}

			if winner {
				return point.X*4_000_000 + point.Y
			}

		}

	}

	return 0
}
