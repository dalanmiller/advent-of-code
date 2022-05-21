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

type State struct {
	RunningCost int
	Amphipods   []Amphipod
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

func (A Amphipod) destinationColumn() int {
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

type World struct {
	Amphipods    []Amphipod
	BurrowHeight int8
	LeafWorlds   []World
	CurrentCost  int
}

func (w World) burrowTypeFull(aType Type) bool {

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
	for i := int8(2); i < w.BurrowHeight+2; i++ {
		for _, amphi := range w.Amphipods {

			// Just need to check if it's in the column and then whether or
			//  not it's the right type or not.
			if amphi.Position.Column == col && amphi.Type != aType {
				return false
			}
		}
	}
	return true
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
	for x := int8(3); x <= 9; x += 2 {

		switch Type(lines[2][x]) {
		case AMBER:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 2},
				Type:     AMBER,
			})
		case BRONZE:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 2},
				Type:     BRONZE,
			})
		case COPPER:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 2},
				Type:     COPPER,
			})
		case DESERT:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 2},
				Type:     DESERT,
			})
		}

		switch Type(lines[3][x]) {
		case AMBER:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 3},
				Type:     AMBER,
			})
		case BRONZE:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 3},
				Type:     BRONZE,
			})
		case COPPER:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 3},
				Type:     COPPER,
			})
		case DESERT:
			amphipods = append(amphipods, Amphipod{
				Position: Position{Column: x, Row: 3},
				Type:     DESERT,
			})
		}
	}

	return amphipods
}


func generate

func optimalPath(w World) int {
	
	

}

func run(input string) int {
	amphipods := parseInput(input)
	initWorld := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
		CurrentCost:  0,
	}

	optimalPath(world)

	return 0
}
