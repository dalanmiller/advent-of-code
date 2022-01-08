package main

import (
	"math"
	"strconv"
	"strings"
)

type Player struct {
	Loc   int
	Score int
}

func (p *Player) move(v int) {
	p.Loc = (p.Loc + v) % 10
	if p.Loc == 0 {
		p.Score += 10
	} else {
		p.Score += p.Loc
	}
}

func (p Player) copy() Player {
	return Player{
		Loc:   p.Loc,
		Score: p.Score,
	}
}

type Dice struct {
	RollN int
	V     int
}

func parseInput(input string) (Player, Player) {
	lines := strings.Split(input, "\n")

	var players []Player
	for i := 0; i < 2; i++ {
		vS := strings.Split(lines[i], ": ")
		v, _ := strconv.Atoi(vS[1])

		players = append(players, Player{
			Loc:   v,
			Score: 0,
		})
	}

	return players[0], players[1]
}

func (d *Dice) roll() int {
	d.V += 1
	if d.V == 101 {
		d.V = 1
	}
	d.RollN++
	return d.V
}

func run(input string) int {
	p1, p2 := parseInput(input)

	dice := Dice{
		0, 0,
	}

	winner := false
	for !winner {
		for _, p := range []*Player{&p1, &p2} {

			// v := roll() + roll() + roll()
			v := dice.roll() + dice.roll() + dice.roll()

			p.move(v)

			if p.Score >= 1000 {
				winner = true
				break
			}
		}
	}

	if p1.Score >= 1000 {
		return p2.Score * dice.RollN
	} else {
		return p1.Score * dice.RollN
	}
}

var outcomes = [27]int{3,
	4,
	5,
	4,
	5,
	6,
	5,
	6,
	7,
	4,
	5,
	6,
	5,
	6,
	7,
	6,
	7,
	8,
	5,
	6,
	7,
	6,
	7,
	8,
	7,
	8,
	9}

type State struct {
	a Player
	b Player
	t Turn
}

type Turn int

const (
	ONE Turn = 0
	TWO Turn = 1
)

// Don't use pointers as map keys!
// https://abhinavg.net/posts/pointers-as-map-keys/
var cache = make(map[State][2]int, 25000)

func playGame(state State) (wins [2]int) {
	a := state.a
	b := state.b
	t := state.t

	for _, outcome := range outcomes {
		var source Player
		if t == ONE {
			source = a
		} else {
			source = b
		}
		newP := source.copy()

		newP.move(outcome)
		if newP.Score >= 21 {
			wins[t]++
		} else {
			var state State
			if t == ONE {
				state = State{newP, b, TWO}
			} else {
				state = State{a, newP, ONE}
			}

			result, ok := cache[state]
			if !ok {
				result = playGame(state)
				cache[state] = result
			}
			wins[0] += result[0]
			wins[1] += result[1]
		}
	}
	return wins
}

func run_two(input string) int {
	p1, p2 := parseInput(input)
	wins := playGame(State{p1, p2, ONE})
	return int(math.Max(float64(wins[0]), float64(wins[1])))
}

func main() {
	run_two(
		`Player 1 starting position: 4
Player 2 starting position: 8`,
	)
}
