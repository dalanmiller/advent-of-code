package main

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

type direction string

const (
	UP    direction = "U"
	LEFT  direction = "L"
	RIGHT direction = "R"
	DOWN  direction = "D"
)

type instruction struct {
	Dir   direction
	Steps int8
}

type coord struct {
	X int
	Y int
}

func (c coord) InSurrounds(oc coord) bool {
	var new coord
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			new = coord{c.X + i, c.Y + j}
			if oc == new {
				return true
			}
		}
	}

	return false
}

func readInput(input io.Reader) []instruction {
	insts := []instruction{}

	s := bufio.NewScanner(input)

	for s.Scan() {
		split := strings.Split(s.Text(), " ")
		dir := direction(split[0])
		steps, _ := strconv.ParseInt(split[1], 10, 8)

		insts = append(insts, instruction{dir, int8(steps)})
	}

	return insts

}

func moveHead(H *coord, dir direction) {
	switch dir {
	case UP:
		H.Y++ // int(inst.Steps)
	case LEFT:
		H.X-- // int(inst.Steps)
	case DOWN:
		H.Y-- // int(inst.Steps)
	case RIGHT:
		H.X++ // int(inst.Steps)
	}
}

func moveTail(H *coord, T *coord, tailKnot bool, positions *map[coord]int) {

	// Is the tail in the 8 positions around the head?, continue
	if *T == *H || H.InSurrounds(*T) {
		return
	}

	// If H moves more than two steps and T is not 'behind' H, then it
	// needs to hop onto same axis of direction.
	xD := int(math.Abs(float64(H.X - T.X)))
	yD := int(math.Abs(float64(H.Y - T.Y)))

	if xD > 1 {
		if yD == 1 {
			T.Y = H.Y
		}

		T.X = incrementTail(H.X, T.X)
	}

	if yD > 1 {
		if xD == 1 {
			T.X = H.X
		}

		T.Y = incrementTail(H.Y, T.Y)
	}

	if tailKnot {
		(*positions)[*T]++
	}
}

func incrementTail(H int, T int) int {
	if H < T {
		return T - 1
	} else {
		return T + 1
	}
}

func run(input io.Reader, numKnots int) int {
	instructions := readInput(input)
	positions := map[coord]int{{0, 0}: 1}
	knots := make([]coord, numKnots)

	// Iterate over instructions
	for _, inst := range instructions {
		// Iterate for each step
		for i := 0; int8(i) < inst.Steps; i++ {
			moveHead(&knots[0], inst.Dir)
			// Iterate each knot
			for j := range knots[1:] {
				moveTail(&knots[j], &knots[j+1], j == len(knots)-2, &positions)
			}
		}
	}

	return len(positions)
}
