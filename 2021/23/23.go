package main

import (
	"fmt"
	"strings"
)

//0 #############
//1 #...........#
//2 ###A#B#C#D###
//3   #A#B#C#D#
//4   #########

type Type string

const (
	AMBER  Type = "A"
	BRONZE Type = "B"
	COPPER Type = "C"
	DESERT Type = "D"
)

type Position struct {
	Column, Row int8
}

type Amphipod struct {
	InHallway bool
	Position  Position
	Type      Type
}

func (A Amphipod) movementCost() int {
	switch A.Type {
	case AMBER:
		return 1
	case BRONZE:
		return 10
	case COPPER:
		return 100
	case DESERT:
		return 1000
	}

	return 0
}

func (A Amphipod) destinationColumn() int8 {
	switch A.Type {
	case AMBER:
		return 3
	case BRONZE:
		return 5
	case COPPER:
		return 7
	case DESERT:
		return 9
	}

	return 0
}

var amphipodTypeMap = map[string]Type{
	"A": AMBER,
	"B": BRONZE,
	"C": COPPER,
	"D": DESERT,
}

type World struct {
	Amphipods    []Amphipod
	BurrowHeight int8
	LeafWorlds   []World
	CurrentCost  int
}

func (w World) burrowTypeCompleteAndOccupants(aType Type) (bool, []bool) {
	var col int8
	switch aType {
	case AMBER:
		col = 3
	case BRONZE:
		col = 5
	case COPPER:
		col = 7
	case DESERT:
		col = 9
	}

	// Okay, we need to check if all the spaces in the column are
	//  filled with the right Amphipod type
	occupancy := make([]bool, w.BurrowHeight)
	// fmt.Printf(occupancy)
	result := true
	for i := int8(2); i < w.BurrowHeight+2; i++ {
		for _, amphi := range w.Amphipods {

			// Just need to check if it's in the column and then whether or
			//  not it's the right type or not.
			if amphi.Position.Column == col {
				occupancy[amphi.Position.Column-2] = true
				if amphi.Type != aType {
					result = false
				}
			}
		}
	}
	return result, occupancy
}

func (w World) possibleMoves(amphipod Amphipod) []Position {
	moves := make([]Position, 0, 12)

	// First we iterate over hallway
	for _, x := range []int8{1, 2, 4, 6, 8, 10, 11} {
		p := Position{x, 2}
		if w.canStopHere(amphipod, p) {
			moves = append(moves, p)
		}
	}

	// ~Then we iterate over each of the spaces in the burrows~
	// Incorrect! We just need to go through the column for this
	// . particular amphipod.
	destColumn := amphipod.destinationColumn()
	deepestBurrowPoint := Position{0, 0}
	for y := int8(2); y < w.BurrowHeight+2; y++ {
		p := Position{destColumn, y}
		if w.canStopHere(amphipod, p) {
			deepestBurrowPoint = p
		}
	}
	moves = append(moves, deepestBurrowPoint)

	return moves
}

func (w World) canStopHere(amphipod Amphipod, proposedDestination Position) bool {
	Row := proposedDestination.Row
	Col := proposedDestination.Column

	if Row == 1 {
		// Can't stop above a burrow
		if Col == 3 || Col == 5 || Col == 7 || Col == 9 {
			return false
		} else {
			// Amphipod can't move from hallway to hallway again
			return amphipod.Position.Row != 1
		}
	} else {

		// Need to check if burrow is empty
		// . if empty we can only go to the bottom of the burrow
		//		if
		// We're trying to get to a burrow... hopefully
		return true
	}
}

func (w World) print() {

	//0 #############
	//1 #...........#
	//2 ###A#B#C#D###
	//3   #A#B#C#D#
	//4   #########

	var output strings.Builder
	output.WriteString("#############\n")
	hallway := []string{"#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#", "\n"}
	for _, amphi := range w.Amphipods {
		if amphi.InHallway {
			hallway[amphi.Position.Column] = string(amphi.Type)
		}
	}
	output.WriteString(strings.Join(hallway, ""))

	// Now the burrow rows
	for i := 0; i < int(w.BurrowHeight); i++ {
		var burrowRow []string
		if i == 0 {
			burrowRow = []string{"#", "#", "#", ".", "#", ".", "#", ".", "#", ".", "###\n"}
		} else {
			burrowRow = []string{" ", " ", "#", ".", "#", ".", "#", ".", "#", ".", "#  \n"}
		}

		// Iterate over all the Amphipods and if they are in this row, pop them in
		//  the correct column
		for _, amphi := range w.Amphipods {
			if !amphi.InHallway && amphi.Position.Row == int8(2+i) {
				burrowRow[amphi.Position.Column] = string(amphi.Type)
			}
		}
		output.WriteString(strings.Join(burrowRow, ""))
	}

	// Now the last row
	output.WriteString("  #########  \n")
	fmt.Print(output.String())
}

func parseInput(input string) []Amphipod {
	lines := strings.Split(input, "\n")
	amphipods := make([]Amphipod, 0, 8)

	// Iterate over hallway (for testing purposes)
	for _, x := range []int8{1, 2, 4, 6, 8, 10, 11} {
		char := string(lines[1][x])
		if char != "." && char != "#" {
			amphipods = append(amphipods, Amphipod{
				Position: Position{x, 1},
				Type:     amphipodTypeMap[char],
			})
		}
	}

	// Iterate over burrows
	for x := int8(3); x <= 9; x += 2 {
		for y := int8(2); y < int8(len(lines)-1); y++ {
			char := string(lines[y][x])
			if char != "." && char != "#" {
				amphipods = append(amphipods, Amphipod{
					Position: Position{x, y},
					Type:     amphipodTypeMap[char],
				})
			}
		}
	}

	return amphipods
}

func run(input string) int {
	amphipods := parseInput(input)
	initWorld := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
		CurrentCost:  0,
	}

	initWorld.print()
	return 0
}
