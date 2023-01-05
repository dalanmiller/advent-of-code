package main

import (
	"fmt"
	"io"
)

type jet rune

const (
	LEFT  jet = '<'
	RIGHT jet = '>'
	DOWN  jet = 'v'
)

// ####

// .#.
// ###
// .#.

// ..#
// ..#
// ###

// #
// #
// #
// #

// ##
// ##

type shapeType int

const (
	HORIZ shapeType = iota
	PLUS
	L
	VERT
	SQUARE
)

type point [2]int

type shape struct {
	ShapeType shapeType
	// Point in which will be the top leftmost point of the shape, even if
	// . that point is not occupied by rock.
	TopPoint point
}

func (ch *chamber) canMove(sh *shape, d jet) bool {
	// Wow this was so tedious to write
	switch sh.ShapeType {
	case HORIZ:
		switch d {
		case RIGHT:
			_, isAdjacentRock := (*ch)[point{sh.TopPoint[0] + 4, sh.TopPoint[1]}]
			return !isAdjacentRock && sh.TopPoint[0] < 3

		case LEFT:
			_, isAdjacentRock := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1]}]
			return !isAdjacentRock && sh.TopPoint[0] > 0

		case DOWN:
			if sh.TopPoint[1] == 0 {
				return false
			}

			for _, i := range []int{0, 1, 2, 3} {
				if _, isBelow := (*ch)[point{sh.TopPoint[0] + i, sh.TopPoint[1] - 1}]; isBelow {
					return false
				}
			}
			return true
			// if _, isAdjacentRock := (*ch)[point{sh.TopPoint[0], sh.TopPoint[1] - 1}]; isAdjacentRock {
			// 	return !isAdjacentRock // && sh.TopPoint[1] > 1
			// }
		}
	case PLUS:
		switch d {
		case RIGHT:
			_, topAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1]}]
			_, midAdj := (*ch)[point{sh.TopPoint[0] + 2, sh.TopPoint[1] - 1}]
			_, botAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1] - 2}]

			return !topAdj && !midAdj && !botAdj && sh.TopPoint[0] < 5

		case LEFT:
			_, topAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1]}]
			_, midAdj := (*ch)[point{sh.TopPoint[0] - 2, sh.TopPoint[1] - 1}]
			_, botAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1] - 2}]

			return !topAdj && !midAdj && !botAdj && sh.TopPoint[0] > 1
		case DOWN:
			_, leftAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1] - 2}]
			_, midAdj := (*ch)[point{sh.TopPoint[0], sh.TopPoint[1] - 3}]
			_, rightAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1] - 2}]

			return !leftAdj && !midAdj && !rightAdj && sh.TopPoint[1] > 3
		}
	case L:
		switch d {
		case RIGHT:
			_, topAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1]}]
			_, midAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1] - 1}]
			_, legAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1] - 2}]

			return !topAdj && !midAdj && !legAdj && sh.TopPoint[0] < 6
		case LEFT:
			_, topAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1]}]
			_, midAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1] - 1}]
			_, legAdj := (*ch)[point{sh.TopPoint[0] - 3, sh.TopPoint[1] - 2}]

			return !topAdj && !midAdj && !legAdj && sh.TopPoint[0] > 2
		case DOWN:
			for _, i := range []int{0, 1, 2} {
				if _, botAdj := (*ch)[point{sh.TopPoint[0] - i, sh.TopPoint[1] - 3}]; botAdj {
					return false
				}
			}
			return true
		}
	case VERT:
		switch d {
		case RIGHT:
			if sh.TopPoint[0] >= 6 {
				return false
			}

			for _, i := range []int{0, 1, 2, 3} {
				if _, rightAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1] - i}]; rightAdj {
					return false
				}
			}
			return true

			// return sh.TopPoint[0] < 6
		case LEFT:
			if sh.TopPoint[0] == 0 {
				return false
			}

			for _, i := range []int{0, 1, 2, 3} {
				if _, rightAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1] - i}]; rightAdj {
					return false
				}
			}
			return true

			// return sh.TopPoint[0] > 0
		case DOWN:
			if _, downAdj := (*ch)[point{sh.TopPoint[0], sh.TopPoint[1] - 4}]; downAdj {
				return false
			}
			return true
		}
	case SQUARE:
		switch d {
		case RIGHT:
			topAdj := (*ch)[point{sh.TopPoint[0] + 2, sh.TopPoint[1]}]
			botAdj := (*ch)[point{sh.TopPoint[0] + 2, sh.TopPoint[1] - 1}]

			return !topAdj && !botAdj && sh.TopPoint[0] < 5
		case LEFT:
			topAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1]}]
			botAdj := (*ch)[point{sh.TopPoint[0] - 1, sh.TopPoint[1] - 1}]

			return !topAdj && !botAdj && sh.TopPoint[0] > 0
		case DOWN:
			leftBotAdj := (*ch)[point{sh.TopPoint[0], sh.TopPoint[1] - 2}]
			rightBotAdj := (*ch)[point{sh.TopPoint[0] + 1, sh.TopPoint[1] - 2}]

			return !leftBotAdj && !rightBotAdj
			// return sh.TopPoint[1] > 2
		}
	}

	return false
}

func (ch *chamber) moveShape(sh *shape, d jet) {
	// Check if can move and break if not
	if !ch.canMove(sh, d) {
		return
	}

	switch d {
	case LEFT:
		sh.TopPoint[0]--
	case RIGHT:
		sh.TopPoint[0]++
	case DOWN:
		sh.TopPoint[1]--
	}
}

func (ch *chamber) print() {
	highestY := 0
	for p, _ := range *ch {
		if p[1] > highestY {
			highestY = p[1]
		}
	}

	rows := make([]string, highestY+1)
	for y := 0; y <= highestY; y++ {
		for x := 0; x <= 6; x++ {
			if z := (*ch)[point{x, y}]; z {
				rows[y] += "#"
			} else {
				rows[y] += "."
			}
		}
	}

	for i := len(rows) - 1; i >= 0; i-- {
		fmt.Println(rows[i])
	}
}

type chamber map[point]bool

// Returns true if the given shape is 'at rest' ontop of another shape or floor
// func (ch *chamber) placed(sh *shape) bool {
// 	sX, sY := sh.TopPoint[0], sh.TopPoint[1]

// 	switch (*sh).ShapeType {
// 	case HORIZ:
// 		for _, dX := range []int{0, 1, 2, 3} {
// 			if _, ok := (*ch)[point{sX + dX, sY - 1}]; ok {
// 				return true
// 			}
// 		}
// 	case PLUS:
// 		// Left arm, bottom arm, or right arm of '+'
// 		return (*ch)[point{sX - 1, sY - 2}] || (*ch)[point{sX, sY - 3}] || (*ch)[point{sX + 1, sY - 2}]
// 	case L:
// 		return (*ch)[point{sX, sY - 3}] || (*ch)[point{sX - 1, sY - 3}] || (*ch)[point{sX - 2, sY - 3}]
// 	case VERT:
// 		return (*ch)[point{sX, sY - 4}]
// 	case SQUARE:
// 		return (*ch)[point{sX, sY - 2}] || (*ch)[point{sX + 1, sY - 2}]
// 	}

// 	return false
// }

func readInput(input io.Reader) []jet {
	buffer, _ := io.ReadAll(input)

	jetStream := make([]jet, 0, len(buffer))
	for _, r := range buffer {
		jetStream = append(jetStream, jet(r))
	}

	return jetStream
}

type key struct {
	rockIndex int
	jetIndex  int
}

type value struct {
	rockStep int
	highestY int
}

func run(input io.Reader, rockLimit int) int {
	jetStream := readInput(input)

	// The tall, vertical chamber is exactly seven units wide. Each rock appears so that its
	// left edge is two units away from the left wall and its bottom edge is three units above
	// the highest rock in the room (or the floor, if there isn't one).

	ch := chamber{}
	for x := 0; x <= 6; x++ {
		ch[point{x, 0}] = true
	}

	var currentRock *shape
	highestRock := &shape{TopPoint: point{0, 0}}
	startPoint := point{2, 4}
	shapeTypeList := []shapeType{HORIZ, PLUS, L, VERT, SQUARE}
	jetStreamIndex := 0

	cache := make(map[key]value, 1000)

	for i := 0; i < rockLimit; i++ {
		currentRock = &shape{
			// ShapeType: shapeTypeList[(i-1)%5],
			ShapeType: shapeTypeList[i%5],
		}

		// Need to fine tune starting point as the real definition is "leftmost part of rock",
		// . is two parts from wall
		switch currentRock.ShapeType {
		case HORIZ:
			currentRock.TopPoint = startPoint
		case PLUS:
			currentRock.TopPoint = point{startPoint[0] + 1, startPoint[1] + 2}
		case L:
			currentRock.TopPoint = point{startPoint[0] + 2, startPoint[1] + 2}
		case VERT:
			currentRock.TopPoint = point{startPoint[0], startPoint[1] + 3}
		case SQUARE:
			currentRock.TopPoint = point{startPoint[0], startPoint[1] + 1}
		}

		if v, cacheHit := cache[key{i % 5, jetStreamIndex}]; cacheHit && i > 20 {
			delta := (rockLimit - i) / (i - v.rockStep)
			mod := (rockLimit - i) % (i - v.rockStep)

			if mod == 0 {
				return highestRock.TopPoint[1] + (highestRock.TopPoint[1]-v.highestY)*delta
			}

		} else {
			cache[key{i % 5, jetStreamIndex}] = value{i, highestRock.TopPoint[1]}
		}

		for {

			// Move by jet stream
			ch.moveShape(currentRock, jetStream[jetStreamIndex])
			jetStreamIndex = (jetStreamIndex + 1) % len(jetStream)

			// Move down
			if !ch.canMove(currentRock, DOWN) {
				if currentRock.TopPoint[1] > highestRock.TopPoint[1] {
					highestRock = currentRock
				}

				// Make sure that the top points are added to chamber
				tp := currentRock.TopPoint
				switch currentRock.ShapeType {
				case HORIZ:
					for _, dX := range []int{0, 1, 2, 3} {
						ch[point{tp[0] + dX, tp[1]}] = true
					}
				case PLUS:
					ch[tp] = true                          // top
					ch[point{tp[0] - 1, tp[1] - 1}] = true // left
					ch[point{tp[0] + 1, tp[1] - 1}] = true // right
					ch[point{tp[0], tp[1] - 2}] = true     // bot
				case L:
					// Top
					ch[tp] = true

					// Mid
					ch[point{tp[0], tp[1] - 1}] = true

					// Bot
					ch[point{tp[0], tp[1] - 2}] = true
					ch[point{tp[0] - 1, tp[1] - 2}] = true
					ch[point{tp[0] - 2, tp[1] - 2}] = true

				case VERT:
					for _, dY := range []int{0, 1, 2, 3} {
						ch[point{tp[0], tp[1] - dY}] = true
					}
				case SQUARE:
					// Top row
					ch[tp] = true
					ch[point{tp[0] + 1, tp[1]}] = true

					// Bot row
					ch[point{tp[0], tp[1] - 1}] = true
					ch[point{tp[0] + 1, tp[1] - 1}] = true
				}

				// ch.print()

				// Break moving
				break
			} else {
				ch.moveShape(currentRock, DOWN)
			}
		}

		startPoint = point{2, highestRock.TopPoint[1] + 4}
	}

	return highestRock.TopPoint[1]
}
