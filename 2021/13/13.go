package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Coord struct {
	X int
	Y int
}

type Fold struct {
	Axis string
	N    int
}

type Paper struct {
	Height int
	Width  int
	Dots   []Coord
}

func (p *Paper) foldUp(y int) {
	newPaper := Paper{
		Height: p.Height - y - 1,
		Width:  p.Width,
	}

	newDots := []Coord{}
	seen := make(map[Coord]bool)
	for _, dot := range p.Dots {
		_, ok := seen[dot]
		if !ok && dot.Y > y {
			newDot := Coord{X: p.Width - dot.X - 1, Y: dot.Y}
			_, newOk := seen[newDot]
			if !newOk {
				newDots = append(newDots, Coord{X: dot.X, Y: p.Height - dot.Y - 1})
				seen[dot] = true
			}
		} else if !ok {
			newDots = append(newDots, dot)
			seen[dot] = true
		}
	}

	p.Dots = newDots
	p.Height = newPaper.Height
	p.Width = newPaper.Width
}

func (p *Paper) foldLeft(x int) {
	newPaper := Paper{
		Height: p.Height,
		Width:  p.Width - x - 1,
	}

	newDots := []Coord{}
	seen := make(map[Coord]bool)
	for _, dot := range p.Dots {
		_, ok := seen[dot]
		if !ok && dot.X > x {
			newDot := Coord{X: p.Width - dot.X - 1, Y: dot.Y}
			_, newOk := seen[newDot]
			if !newOk {
				newDots = append(newDots, newDot)
				seen[newDot] = true
			}
		} else if !ok {
			newDots = append(newDots, dot)
			seen[dot] = true
		}
	}

	p.Dots = newDots
	p.Height = newPaper.Height
	p.Width = newPaper.Width
}

func parseInput(input string) (*Paper, []Fold) {
	lines := strings.Split(input, "\n")

	dots := []Coord{}
	folds := []Fold{}
	var maxY, maxX int
	for _, line := range lines {
		if line == "" {
			continue
		}

		if unicode.IsDigit(rune(line[0])) {

			rawCoords := strings.Split(line, ",")
			x, _ := strconv.Atoi(rawCoords[0])
			y, _ := strconv.Atoi(rawCoords[1])
			dots = append(dots, Coord{
				X: x,
				Y: y,
			})

			if x > maxX {
				maxX = x
			}

			if y > maxY {
				maxY = y
			}

		} else {

			foldInstruction := strings.Split(line, "=")
			axis := string(foldInstruction[0][len(foldInstruction[0])-1])
			n, _ := strconv.Atoi(foldInstruction[1])

			folds = append(folds, Fold{
				Axis: axis,
				N:    n,
			})
		}
	}

	paper := Paper{
		Height: maxY + 1,
		Width:  maxX + 1,
		Dots:   dots,
	}

	return &paper, folds
}

func run(input string, nFolds int) int {
	paper, folds := parseInput(input)

	for i, fold := range folds {
		if i+1 > nFolds {
			return len(paper.Dots)
		}
		switch fold.Axis {
		case "x":
			paper.foldLeft(fold.N)
		case "y":
			paper.foldUp(fold.N)
		}
	}

	coordMap := make(map[Coord]bool)
	for _, coord := range paper.Dots {
		coordMap[coord] = true
	}
	for i := 0; i < paper.Height; i++ {
		for j := 0; j < paper.Width; j++ {
			_, ok := coordMap[Coord{X: j, Y: i}]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	return len(paper.Dots)
}
