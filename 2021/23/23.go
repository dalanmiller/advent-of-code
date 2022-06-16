package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/fatih/color"
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
	Type     Type
	Position Position
}

type Journey struct {
	FromTo [2]Position
	Type   Type
}

var costToMoveMemoization = make(map[Journey]uint16, 48)

func (A Amphipod) costToMove(p2 Position) uint16 {
	p1 := A.Position
	j := Journey{[2]Position{p1, p2}, A.Type}
	if cost, ok := costToMoveMemoization[j]; ok {
		return cost
	}

	energyCost := amphipodCostMap[A.Type]
	xDelta := uint16(math.Abs(float64(p1.Column - p2.Column)))
	cost := uint16(0)

	if p1.Row > 1 && p2.Row == 1 {
		yDelta := uint16(p1.Row - 1)

		cost = (yDelta + xDelta) * energyCost
	} else if p1.Row == 1 && p2.Row > 1 {
		yDelta := uint16(p2.Row - 1)
		cost = (yDelta + xDelta) * energyCost
	} else {

		// Lastly, we are moving from a burrow to a burrow
		y1Delta := uint16(p1.Row - 1)
		y2Delta := uint16(p2.Row - 1)
		cost = (y1Delta + y2Delta + xDelta) * energyCost
	}

	costToMoveMemoization[j] = cost
	return cost
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

var amphipodCostMap = map[Type]uint16{
	AMBER:  1,
	BRONZE: 10,
	COPPER: 100,
	DESERT: 1000,
}

var amphipodColorMap = map[Type]func(format string, a ...interface{}) string{
	AMBER:  color.YellowString,
	BRONZE: color.RedString,
	COPPER: color.BlueString,
	DESERT: color.GreenString,
}

type World struct {
	LeafWorlds     [48]*World
	Amphipods      [16]Amphipod
	AdditionalCost uint16
	CurrentCost    uint16
	AmphipodsLeft  int8
	BurrowHeight   int8
	Complete       bool
}

func (w World) h() uint16 {
	return w.CurrentCost + w.costToCompletion()
}

func (w World) burrowTypeCompleteAndOccupants(aType Type) (bool, []Type) {
	col := amphipodDestinationColumnMap[aType]

	// Okay, we need to check if all the spaces in the column are
	//  filled with the right Amphipod type
	occupancy := make([]Type, w.BurrowHeight)
	result := true
	for i := int8(2); i < w.BurrowHeight+2; i++ {

		// Now iterate through the Amphipods and check if
		// any are in this Position
		for _, amphi := range w.Amphipods {
			p := amphi.Position

			// If we are examining an amphi whose type
			// . is the target type for this burrow, but
			// . they are somewhere else, result is false
			if aType == amphi.Type && p.Column != col {
				result = false
			}

			// If the amphi is in the target column and
			// . is in the row we are looking for
			if p.Column == col && p.Row == i {

				// Ensure we mark the occupancy
				occupancy[p.Row-2] = amphi.Type

				// And we check if the type matches that
				// of the target type for the burrow.
				// If it doesn't match the result is false
				if amphi.Type != aType {
					result = false
				}
			}
		}
	}

	// Case handling of if no amphipods are in the space then all the occupancy slots
	// . are actually 'false'
	for _, occu := range occupancy {
		if occu == "" {
			result = false
		}
	}

	return result, occupancy
}

func (w World) possibleMoves(amphipod Amphipod) []Position {

	// Amphipod is where it should be, then return no possible moves
	// Amphipod in burrow
	// > is target burrow available, if yes, return that option
	// > if not, return a hallway
	// amphipod in hallway
	// > is target burrow afvailable if yes, return that option
	// lastly, return no possible moves

	// If the amphipod is right where it should be then we
	// don't need to move it anywhere!
	complete, occupancy := w.burrowTypeCompleteAndOccupants(amphipod.Type)
	if amphipod.Position.Row > 1 && amphipod.Position.Column == amphipodDestinationColumnMap[amphipod.Type] {
		allFalseButLast := false
		for i, occu := range occupancy {
			// If there is an amphipod here that doesn't match
			if int8(i+2) > amphipod.Position.Row && occu != amphipod.Type {
				break
			}

			if i == len(occupancy)-1 && occu != "" {
				allFalseButLast = true
				break
			}
		}

		if complete || allFalseButLast {
			return []Position{}
		}
	}

	moves := make([]Position, 0, 12)
	dCol := amphipodDestinationColumnMap[amphipod.Type]

	if !complete {
		var deepestBurrowPoint Position
		foundSpot := false
		for i, p := range []Position{{dCol, 2}, {dCol, 3}, {dCol, 4}, {dCol, 5}, {dCol, 6}, {dCol, 7}, {dCol, 8}} {
			if int8(i) >= w.BurrowHeight {
				break
			}
			if w.canStopHere(amphipod, p) {
				deepestBurrowPoint = p
				foundSpot = true
			}
		}

		if foundSpot {
			return []Position{deepestBurrowPoint}
		}
	}

	// If we got this far, then we didn't find a spot in the burrow
	for _, pos := range []Position{{1, 1}, {2, 1}, {4, 1}, {6, 1}, {8, 1}, {10, 1}, {11, 1}} {
		if w.canStopHere(amphipod, pos) {
			moves = append(moves, pos)
		}
	}

	return moves
}

func (w World) canStopHere(a Amphipod, proposedDestination Position) bool {
	Row := proposedDestination.Row
	Col := proposedDestination.Column

	// Firstly, can't stop on self
	if a.Position.Row == Row && a.Position.Column == Col {
		return false
	}

	for _, b := range w.Amphipods {
		// Can't stop there if another a is there!
		if Row == b.Position.Row && Col == b.Position.Column {
			return false
		}

		// Can't move if ampihpods above
		if a.Position.Column == b.Position.Column && a.Position.Row > b.Position.Row {
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

			if b.Position.Row != 1 {
				continue
			}

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

		// Can't stop here if there are amphipods in the hallway between
		// . current burrow and destination burrow
		for _, b := range w.Amphipods {
			if b.Position.Row != 1 {
				continue
			}

			// If amphipod is moving rightwards
			if Col > a.Position.Column && b.Position.Column > a.Position.Column && b.Position.Column < Col {
				return false
			}

			// If amphipod is moving leftwards
			if Col < a.Position.Column && b.Position.Column < a.Position.Column && b.Position.Column > Col {
				return false
			}
		}

		// Can't stop here if there are amphipods blocking the positions in the
		// . burrow that are above this position
		_, occupancy := w.burrowTypeCompleteAndOccupants(a.Type)
		for y := Row; y > 1; y-- {
			if occupancy[y-2] != "" {
				return false
			}
		}

		// Can't stop here if there are amphipods of the wrong type below proposed
		// . destination spot
		for y := Row + 1; y < w.BurrowHeight+2; y++ {
			for _, b := range w.Amphipods {
				if b.Position.Row == y && b.Position.Column == Col && a.Type != b.Type {
					return false
				}
			}
		}

		// Need to check if burrow is empty
		// . if empty we can only go to the bottom of the burrow
		return true
	}

	return true
}

func (w World) burrowsComplete() bool {
	for _, a := range w.Amphipods {
		if a.Position.Row == 1 || a.Position.Column != amphipodDestinationColumnMap[a.Type] {
			return false
		}
	}
	return true
}

// This function does a simple estimation on getting amphipods
// . to their final burrow place. It's not exact.

var costToCompletionMemoization = make(map[[16]Amphipod]uint16, 100000)

func (w World) costToCompletion() uint16 {
	if v, ok := costToCompletionMemoization[w.Amphipods]; ok {
		return v
	}
	cost := uint16(0)
	for _, a := range w.Amphipods {
		if a.Type == "" {
			continue
		}
		destCol := amphipodDestinationColumnMap[a.Type]
		if a.Position.Column != destCol {
			cost += a.costToMove(Position{destCol, 2})
		}
	}
	costToCompletionMemoization[w.Amphipods] = cost
	return cost
}

func (w *World) Print() string {
	//  0123456789012
	//0 #############
	//1 #...........#
	//2 ###A#B#C#D###
	//3   #A#B#C#D#
	//4   #########

	var output strings.Builder
	output.WriteString("\n#############\n")
	hallway := []string{"#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#", "\n"}
	for _, amphi := range w.Amphipods {
		if amphi.Position.Row == 1 {
			typeChar := string(amphi.Type)
			colorFunc := amphipodColorMap[amphi.Type]
			hallway[amphi.Position.Column] = colorFunc(typeChar)
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
			if amphi.Position.Row != 1 && amphi.Position.Row == int8(2+i) {
				typeChar := string(amphi.Type)
				colorFunc := amphipodColorMap[amphi.Type]
				burrowRow[amphi.Position.Column] = colorFunc(typeChar)
			}
		}
		output.WriteString(strings.Join(burrowRow, ""))
	}

	// Now the last row
	output.WriteString("  #########  \n")
	output.WriteString(fmt.Sprintf("CurrentCost: %d\n", w.CurrentCost))
	output.WriteString(fmt.Sprintf("h(): %d\n", w.h()))
	return output.String()
}

func parseInput(input string) (amphipods [16]Amphipod) {
	lines := strings.Split(input, "\n")
	i := 0

	// Iterate over hallway (for testing purposes)
	for _, x := range []int8{1, 2, 4, 6, 8, 10, 11} {
		char := string(lines[1][x])
		if char != "." && char != "#" {
			amphipods[i] = Amphipod{
				Position: Position{x, 1},
				Type:     amphipodTypeMap[char],
			}
			i++
		}
	}

	// Iterate over burrows
	for x := int8(3); x <= 9; x += 2 {
		for y := int8(2); y < int8(len(lines)-1); y++ {
			char := string(lines[y][x])
			if char != "." && char != "#" {
				amphipods[i] = Amphipod{
					Position: Position{x, y},
					Type:     amphipodTypeMap[char],
				}
				i++
			}
		}
	}

	return amphipods
}

var worldMemoization = make(map[[16]Amphipod][48]*World, 1024)

func (w *World) generateLeafWorlds() [48]*World {
	if worlds, ok := worldMemoization[w.Amphipods]; ok {
		return worlds
	}

	newWorlds := make([]*World, 0, 64)
	for i, a := range w.Amphipods {
		if a.Type == "" {
			continue
		}

		// Identify all the moves this amphipod can make
		//  given this world
		moves := w.possibleMoves(a)

		// Generate the costs of all these moves
		for _, move := range moves {
			// Generate new worlds where this amphipod has moved
			// worlds and update the world cost

			// Copy amphipods from current world,
			// move just the one amphipod.
			// newAmphipods := [8]Amphipod{}
			newAmphipods := w.Amphipods
			newAmphipods[i].Position = move

			cost := a.costToMove(move)
			newW := World{
				Amphipods:      newAmphipods,
				BurrowHeight:   w.BurrowHeight,
				CurrentCost:    w.CurrentCost + cost,
				AdditionalCost: cost,
				AmphipodsLeft:  w.AmphipodsLeft,
			}

			// Reduce amphipods left if moving into a burrow
			// . given that amphipods can only move into a burrow if its
			// . the right one.
			if move.Row > 1 {
				newW.AmphipodsLeft -= 1
			}

			if newW.burrowsComplete() {
				newW.Complete = true
			}
			newWorlds = append(newWorlds, &newW)
		}
	}
	sort.Slice(newWorlds, func(i, j int) bool {
		return newWorlds[i].AdditionalCost < newWorlds[j].AdditionalCost
	})

	newWorldsArray := [48]*World{}
	for i, world := range newWorlds {
		newWorldsArray[i] = world
	}
	worldMemoization[w.Amphipods] = newWorldsArray

	return newWorldsArray
}

func printPath(m map[World]*World, start *World) {
	ok := true
	var wN *World
	w := start
	for ok {
		wN, ok = m[*w]
		log.Printf("%s\nTotal Cost: %d", wN.Print(), wN.CurrentCost)
		w = wN
	}
}

func printPathFrom(m map[World]*World, end *World) {
	history := make([]*World, 0, len(m))
	var wN *World
	w, ok := m[*end]
	history = append([]*World{w, end}, history...)
	for ok {
		wN, ok = m[*w]
		if ok {
			history = append([]*World{wN}, history...)
		}
		w = wN
	}

	for _, w := range history {
		log.Printf("%sTotal Cost: %d\n\n", w.Print(), w.CurrentCost)
	}
}

type LowCostQueue struct {
	worlds   []*World
	contains map[*World]bool
	fScore   map[*World]uint16
}

func (lcq LowCostQueue) Len() int { return len(lcq.worlds) }

func (lcq LowCostQueue) Less(i, j int) bool {
	return lcq.fScore[lcq.worlds[i]] < lcq.fScore[lcq.worlds[j]]
}

func (lcq LowCostQueue) Swap(i, j int) {
	lcq.worlds[i], lcq.worlds[j] = lcq.worlds[j], lcq.worlds[i]
}

func (lcq *LowCostQueue) Push(x any) {
	(*lcq).worlds = append((*lcq).worlds, x.(*World))
	lcq.contains[x.(*World)] = true
}

func (lcq *LowCostQueue) Pop() any {
	old := (*lcq).worlds
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	(*lcq).worlds = old[0 : n-1]
	delete((*lcq).contains, item)
	return item
}

func (lcq LowCostQueue) Contains(w *World) bool {
	return lcq.contains[w]
}

func aStarRestart(start World) uint16 {
	lcq := LowCostQueue{
		worlds:   make([]*World, 0, 256),
		contains: make(map[*World]bool, 256),
		fScore:   make(map[*World]uint16, 256),
	}

	heap.Init(&lcq)
	heap.Push(&lcq, &start)

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.

	// My take on this is that it's a backwards map
	cameFrom := make(map[*World]*World)

	lcq.fScore[&start] = start.h()

	// For node n, gScore[n] is the cost of the cheapest path from start
	// . to n currently known.
	gScore := make(map[*World]uint16)
	gScore[&start] = 0

	// minComplete := World{
	// 	CurrentCost: 0,
	// }
	// var lastAdded World

	for lcq.Len() > 0 {

		current := heap.Pop(&lcq).(*World)
		// log.Printf(current.Print())

		if current.Complete || current.burrowsComplete() {
			return current.CurrentCost
		}

		// Iterate through each leafWorld of current
		leafWorlds := current.generateLeafWorlds()
		for i, _ := range leafWorlds {
			if leafWorlds[i] == nil {
				continue
			}

			// We want to calculate the value for getting to this adjacent world
			tentative_gScore := gScore[current] + leafWorlds[i].AdditionalCost
			if adjWorldGScore, ok := gScore[leafWorlds[i]]; !ok || tentative_gScore < adjWorldGScore {
				cameFrom[leafWorlds[i]] = current
				gScore[leafWorlds[i]] = tentative_gScore
				lcq.fScore[leafWorlds[i]] = tentative_gScore + leafWorlds[i].h()
				if !lcq.Contains(leafWorlds[i]) {
					heap.Push(&lcq, leafWorlds[i])
				}
			}
		}
	}

	return 0
}

func run(input string, burrowHeight int8) uint16 {
	amphipods := parseInput(input)
	initWorld := World{
		Amphipods:      amphipods,
		BurrowHeight:   burrowHeight,
		CurrentCost:    0,
		AdditionalCost: 0,
		AmphipodsLeft:  int8(4 * burrowHeight),
	}

	return aStarRestart(initWorld)
}
