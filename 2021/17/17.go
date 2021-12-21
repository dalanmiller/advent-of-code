package main

import (
	"math"
	"strconv"
	"strings"
)

type TargetBox struct {
	XStart, XEnd, YStart, YEnd int
}

type Coord struct {
	X int
	Y int
}

type VelocityEstimate struct {
	DeltaV int
	DeltaY int
	MaxY   int
}

type Probe struct {
	DeltaV       int
	DeltaY       int
	CurrentCoord Coord
}

func (p *Probe) Step() {
	p.CurrentCoord = Coord{
		X: p.CurrentCoord.X + p.DeltaV,
		Y: p.CurrentCoord.Y + p.DeltaY,
	}

	if p.DeltaV > 0 {
		p.DeltaV--
	}
	p.DeltaY--
}

func (p Probe) BeyondTarget(t TargetBox) bool {
	if p.CurrentCoord.X > t.XEnd || // If we are beyond X
		p.CurrentCoord.Y < t.YEnd || // If we are beyond Y
		(p.CurrentCoord.X < t.XStart && p.DeltaV == 0) { // If we are stalled and not yet at minX
		return true
	}
	return false
}

func (p Probe) WithinTarget(t TargetBox) bool {
	if t.XEnd >= p.CurrentCoord.X &&
		p.CurrentCoord.X >= t.XStart &&
		t.YEnd <= p.CurrentCoord.Y &&
		p.CurrentCoord.Y <= t.YStart {
		return true
	}
	return false
}

type Trajectory []Coord

func parseInput(input string) *TargetBox {
	v := strings.Split(input, ": ")
	coords := strings.Split(v[1], ", ")
	xString := strings.TrimPrefix(coords[0], "x=")
	yString := strings.TrimPrefix(coords[1], "y=")

	xStartEnd := strings.Split(xString, "..")
	yStartEnd := strings.Split(yString, "..")

	xStart, _ := strconv.Atoi(xStartEnd[0])
	xEnd, _ := strconv.Atoi(xStartEnd[1])

	yEnd, _ := strconv.Atoi(yStartEnd[0])
	yStart, _ := strconv.Atoi(yStartEnd[1])

	t := TargetBox{
		XStart: xStart,
		XEnd:   xEnd,
		YStart: yStart,
		YEnd:   yEnd,
	}

	return &t
}

func generatePossibleVelocities(lowerX int, upperX int, lowerY int, upperY int) []VelocityEstimate {
	max := upperX * int(math.Abs(float64(upperY)-float64(lowerY)))
	possibilities := make([]VelocityEstimate, 0, max)
	for i := lowerX; i <= upperX; i++ {
		for j := lowerY; j <= upperY; j++ {
			possibilities = append(possibilities, VelocityEstimate{
				DeltaV: i,
				DeltaY: j,
			})
		}
	}

	return possibilities
}

func run(input string) (int, int) {
	targetBox := parseInput(input)

	possibleVelocities := generatePossibleVelocities(
		int(math.Pow(float64(targetBox.XStart*2), 0.5)),
		targetBox.XEnd,
		-90,
		90,
	)

	validVelocities := make([]VelocityEstimate, 0, len(possibleVelocities)/2)
	for _, e := range possibleVelocities {
		p := Probe{
			DeltaV:       e.DeltaV,
			DeltaY:       e.DeltaY,
			CurrentCoord: Coord{X: 0, Y: 0},
		}

		for !p.BeyondTarget(*targetBox) {
			p.Step()

			if p.CurrentCoord.Y > e.MaxY {
				e.MaxY = p.CurrentCoord.Y
			}

			if p.WithinTarget(*targetBox) {
				validVelocities = append(validVelocities, e)
				break
			}
		}
	}

	max := 0
	for _, e := range validVelocities {
		if e.MaxY > max {
			max = e.MaxY
		}
	}

	return max, len(validVelocities)
}
