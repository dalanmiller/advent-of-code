package main

import (
	"container/heap"
	"log"
	"math"
	"sort"
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
	Position Position
	Type     Type
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
	Amphipods    [8]Amphipod
	BurrowHeight int8
	// LeafWorlds     []*World
	CurrentCost    int
	AdditionalCost int
	Complete       bool
}

type WorldList []*World

func (w WorldList) Len() int {
	return len(w)
}

func (w WorldList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WorldList) Less(i, j int) bool {
	return w[i].CurrentCost < w[j].CurrentCost
}

func (w World) burrowTypeCompleteAndOccupants(aType Type) (bool, []bool) {
	col := amphipodDestinationColumnMap[aType]

	// Okay, we need to check if all the spaces in the column are
	//  filled with the right Amphipod type
	occupancy := make([]bool, w.BurrowHeight)
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
				occupancy[p.Row-2] = true

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
		if !occu {
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
	if amphipod.Position.Row > 1 && amphipod.Position.Column == amphipodDestinationColumnMap[amphipod.Type] {
		complete, occupancy := w.burrowTypeCompleteAndOccupants(amphipod.Type)
		allFalseButLast := false
		for i, occu := range occupancy {
			if i == len(occupancy)-1 && occu {
				allFalseButLast = true
				break
			}
		}

		if complete || allFalseButLast {
			return []Position{}
		}
	}

	moves := make([]Position, 0, 12)
	destColumn := amphipodDestinationColumnMap[amphipod.Type]

	complete, _ := w.burrowTypeCompleteAndOccupants(amphipod.Type)

	if !complete {
		var deepestBurrowPoint Position
		foundSpot := false
		for row := int8(2); row < w.BurrowHeight+2; row++ {
			p := Position{destColumn, row}
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
	for _, col := range []int8{1, 2, 4, 6, 8, 10, 11} {
		p := Position{col, 1}

		if w.canStopHere(amphipod, p) {
			moves = append(moves, p)
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

		// Can't stop here if there are amphipods blocking the entrance
		// . above this position
		_, occupancy := w.burrowTypeCompleteAndOccupants(a.Type)
		for y := Row; y > 1; y-- {
			if occupancy[y-w.BurrowHeight] {
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

func (w *World) Print() string {

	//0 #############
	//1 #...........#
	//2 ###A#B#C#D###
	//3   #A#B#C#D#
	//4   #########

	var output strings.Builder
	output.WriteString("#############\n")
	hallway := []string{"#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#", "\n"}
	for _, amphi := range w.Amphipods {
		if amphi.Position.Row == 1 {
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
			if amphi.Position.Row != 1 && amphi.Position.Row == int8(2+i) {
				burrowRow[amphi.Position.Column] = string(amphi.Type)
			}
		}
		output.WriteString(strings.Join(burrowRow, ""))
	}

	// Now the last row
	output.WriteString("  #########  \n")
	return output.String()
}

func parseInput(input string) [8]Amphipod {
	lines := strings.Split(input, "\n")
	var amphipods [8]Amphipod
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

// type WorldQueue struct {
// 	Worlds []*World
// }

// func (wq *WorldQueue) Pop() *World {
// 	world := wq.Worlds[0]
// 	wq.Worlds = wq.Worlds[1:]
// 	return world
// }

// func (wq *WorldQueue) Push(w *World) {
// 	wq.Worlds = append(wq.Worlds, w)
// }

// func (wq *WorldQueue) Len() int {
// 	return len(wq.Worlds)
// }

func (w *World) generateLeafWorlds() []*World {
	newWorlds := make([]*World, 0)
	for i, a := range w.Amphipods {
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
			}

			if newW.burrowsComplete() {
				newW.Complete = true
			}
			newWorlds = append(newWorlds, &newW)
		}
	}
	sort.Sort(WorldList(newWorlds))

	return newWorlds
}

// func generateTree(w *World, completed *[]*World) {

// 	worldsToGenerate := WorldQueue{
// 		Worlds: make([]*World, 0, 100000),
// 	}
// 	worldsToGenerate.Push(w)

// 	for worldsToGenerate.Len() > 0 {

// 		currentWorld := worldsToGenerate.Pop()

// 		newWorlds := make([]*World, 0)
// 		for i, a := range currentWorld.Amphipods {
// 			// Identify all the moves this amphipod can make
// 			//  given this world
// 			moves := currentWorld.possibleMoves(a)

// 			// Generate the costs of all these moves
// 			for _, move := range moves {
// 				// Generate new worlds where this amphipod has moved
// 				// worlds and update the world cost

// 				// Copy amphipods from current world,
// 				// move just the one amphipod.
// 				newAmphipods := make([]Amphipod, len(w.Amphipods))
// 				copy(newAmphipods, currentWorld.Amphipods)
// 				newAmphipods[i].Position = move

// 				cost := a.costToMove(move)
// 				newW := World{
// 					Amphipods:      newAmphipods,
// 					BurrowHeight:   w.BurrowHeight,
// 					CurrentCost:    currentWorld.CurrentCost + cost,
// 					AdditionalCost: cost,
// 				}
// 				newW.print()

// 				newWorlds = append(newWorlds, &newW)
// 			}
// 		}

// 		for _, world := range newWorlds {
// 			// Check if the world is complete,
// 			// if not then generate that world.
// 			if !world.burrowsComplete() {
// 				worldsToGenerate.Push(world)
// 			} else {
// 				world.Complete = true
// 				*completed = append(*completed, world)
// 			}
// 		}

// 		sort.Sort(WorldList(newWorlds))

// 		// Set children of current world
// 		currentWorld.LeafWorlds = newWorlds
// 	}
// }

func printPath(m map[World]*World, start *World) {
	ok := true
	var wN *World
	w := start
	for ok {
		wN, ok = m[*w]
		log.Print(wN.Print())
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
		log.Printf("%s\n", w.Print())
	}
}

type LowCostQueue struct {
	worlds   []*World
	contains map[World]bool
}

func (lrq LowCostQueue) Len() int {
	return len(lrq.worlds)
}

func (lrq LowCostQueue) Less(i, j int) bool {
	return lrq.worlds[i].CurrentCost < lrq.worlds[j].CurrentCost
}

func (lrq LowCostQueue) Swap(i, j int) {
	lrq.worlds[i], lrq.worlds[j] = lrq.worlds[j], lrq.worlds[i]
}

func (lrq *LowCostQueue) Push(x interface{}) {
	world := x.(*World)
	(*lrq).worlds = append((*lrq).worlds, world)
	lrq.contains[*world] = true
}

func (lrq *LowCostQueue) Pop() interface{} {
	old := (*lrq).worlds
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	(*lrq).worlds = old[0 : n-1]
	delete(lrq.contains, *item)
	return item
}

func (lrq *LowCostQueue) Contains(w World) bool {
	return lrq.contains[w]
}

func aStarRestart(start *World) int {
	lrq := LowCostQueue{
		worlds:   []*World{},
		contains: map[World]bool{},
	}

	heap.Init(&lrq)
	heap.Push(&lrq, start)

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.

	// My take on this is that it's a backwards map
	cameFrom := make(map[World]*World)

	//
	// pathTo := make(map[World]*World)

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := make(map[World]int)
	gScore[*start] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how short a path from start to finish can be if it goes through n.

	fScore := make(map[World]int)
	fScore[*start] = 0

	for lrq.Len() > 0 {

		current := lrq.Pop().(*World)

		if current.Complete {
			printPathFrom(cameFrom, current)
			return current.CurrentCost
		}

		leafWorlds := current.generateLeafWorlds()

		for _, adjWorld := range leafWorlds {

			tentative_gScore := gScore[*current] + adjWorld.AdditionalCost
			adjWorldGScore, ok := gScore[*adjWorld]
			if tentative_gScore < adjWorldGScore || !ok {
				cameFrom[*adjWorld] = current
				gScore[*adjWorld] = tentative_gScore
				fScore[*adjWorld] = tentative_gScore + adjWorld.AdditionalCost
				if !lrq.Contains(*adjWorld) {
					lrq.Push(adjWorld)
				}
			}
		}
	}

	return -1
}

func aStar(start *World) int {

	// completedWorlds := make([]*World, 0)
	lrq := LowCostQueue{}
	heap.Init(&lrq)
	heap.Push(&lrq, start)

	path := make(map[World]*World)

	gScore := map[World]int{
		*start: 0,
	}

	for lrq.Len() > 0 {

		// Find min node in the open list
		current := heap.Pop(&lrq).(*World)

		// If the current world has no leaves then we need to
		// . generate some as we haven't inspected this World before
		leafWorlds := current.generateLeafWorlds()

		for _, adjWorld := range leafWorlds {
			// log.Print(adjWorld.Screen)

			// If we find a world that is complete, then we halt
			// . but it may be premature.
			if adjWorld.Complete {
				path[*current] = adjWorld
				printPath(path, start)
				v := gScore[*current]
				return v + adjWorld.AdditionalCost
			}

			// Looking before we leap, we take the cost to get to current node
			// . plus the adjacent node and it's cost.
			tentativeCost := gScore[*current] + adjWorld.AdditionalCost
			_, ok := gScore[*adjWorld]
			// !ok maps to the default value of nil and thus
			// the tentative score is better
			if !ok || tentativeCost < gScore[*adjWorld] {
				path[*current] = adjWorld
				gScore[*adjWorld] = tentativeCost

				heap.Push(&lrq, adjWorld)
			}
		}
	}

	// JIC
	return -1
}

func run(input string, burrowHeight int8) int {
	amphipods := parseInput(input)
	initWorld := World{
		Amphipods:    amphipods,
		BurrowHeight: burrowHeight,
		CurrentCost:  0,
	}

	// completedWorlds := make([]*World, 0)
	// generateTree(&initWorld, &completedWorlds)

	// sort.Sort(WorldList(completedWorlds))

	return aStarRestart(&initWorld)
}
