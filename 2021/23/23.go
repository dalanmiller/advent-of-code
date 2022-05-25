package main

import (
	"fmt"
	// "log"
	"math"
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

func (A Amphipod) costToMove(p2 Position) int {
	p1 := A.Position
	cost := amphipodCostMap[A.Type]
	xDelta := int(math.Abs(float64(p1.Column - p2.Column)))

	if p1.Row > 1 && p2.Row == 1 {
		yDelta := int(p1.Row - 1)
		return (yDelta + xDelta) * cost
	} else if p1.Row == 1 && p2.Row > 1 {
		yDelta := int(p2.Row - 1)
		return (yDelta + xDelta) * cost
	} else {

		// Lastly, we are moving from a burrow to a burrow
		y1Delta := int(p1.Row - 1)
		y2Delta := int(p2.Row - 1)
		return (y1Delta + y2Delta + xDelta) * cost
	}

}

var amphipodDestinationColumnMap = map[Type]int8{
	AMBER:  3,
	BRONZE: 5,
	COPPER: 7,
	DESERT: 9,
}

var amphipodTypeMap = map[string]Type{
	"A": AMBER,
	"B": BRONZE,
	"C": COPPER,
	"D": DESERT,
}

var amphipodCostMap = map[Type]int{
	AMBER:  1,
	BRONZE: 10,
	COPPER: 100,
	DESERT: 1000,
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
	for _, col := range []int8{1, 2, 4, 6, 8, 10, 11} {
		p := Position{col, 1}
		// log.Printf("Checking %v", p)
		if w.canStopHere(amphipod, p) {
			moves = append(moves, p)
		}
	}

	// ~Then we iterate over each of the spaces in the burrows~
	// Incorrect! We just need to go through the column for this
	// . particular amphipod.
	destColumn := amphipodDestinationColumnMap[amphipod.Type]
	deepestBurrowPoint := Position{0, 0}
	for row := int8(2); row < w.BurrowHeight+2; row++ {
		p := Position{destColumn, row}
		// log.Printf("Checking %v", p)
		if w.canStopHere(amphipod, p) {
			deepestBurrowPoint = p
		}
	}

	// Prepend as we likely want to explore the path where
	// the amphipod moves directly into the burrow it belongs
	// to first.
	moves = append([]Position{deepestBurrowPoint}, moves...)

	return moves
}

func (w World) canStopHere(a Amphipod, proposedDestination Position) bool {
	Row := proposedDestination.Row
	Col := proposedDestination.Column

	// Firstly, can't stop on self
	if a.Position.Row == Row && a.Position.Column == Col {
		return false
	}

	// Can't stop there if another a is there!
	for _, b := range w.Amphipods {
		// log.Printf("Checking %v against %v", proposedDestination, b.Position)
		if Row == b.Position.Row && Col == b.Position.Column {
			return false
		}
	}

	if Row == 1 {
		// Can't stop above a burrow
		if Col == 3 || Col == 5 || Col == 7 || Col == 9 {
			return false
		}

		if a.Position.Row == 1 {
			// a can't move from hallway to hallway again
			return false
		}

		// Can't move to a location in hallway if blocked
		// by another amphipod
		for _, b := range w.Amphipods {
			// Somewhat uncertain if this excludes all valid cases, but I think it's good

			// Check if proposed destination is rightwards and
			// whether there is amphipod in the way
			if a.Position.Column < b.Position.Column && b.Position.Column < Col {
				return false
			}

			// Now check if the proposed destination is leftwards and
			// if there is amphipod in way.
			if a.Position.Column > b.Position.Column && b.Position.Column > Col {
				return false
			}
		}

	} else {

		// Need to check if burrow is empty
		// . if empty we can only go to the bottom of the burrow
		return true
	}

	return true
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
