package main

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var lineRegexp = regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnel[s]? lead[s]? to valve[s]? ([A-Z,\s]*)`)

func readInput(input io.Reader) (*map[string]int, *map[string]int, *map[string]map[string]int) {
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB

	s := bufio.NewScanner(input)

	tunnelMap := map[string][]string{}
	rateMap := map[string]int{}
	bitfieldMap := map[string]int{}
	shortestPathMap := map[string]map[string]int{}

	for s.Scan() {
		line := s.Text()
		matches := lineRegexp.FindAllStringSubmatch(line, -1)

		destTunnels := strings.Split(matches[0][3], ", ")
		rate, _ := strconv.Atoi(matches[0][2])

		if rate > 0 {
			rateMap[matches[0][1]] = rate
		}
		tunnelMap[matches[0][1]] = destTunnels
	}

	// Assemble bitfield map
	i := 1
	for valve := range rateMap {
		bitfieldMap[valve] = 1 << i
		i++
	}

	contains := func(s string, sl []string) bool {
		for _, i := range sl {
			if s == i {
				return true
			}
		}

		return false
	}

	// Assemble shortestPathMap via Henry=Warshall
	for valveX, destTunnelsX := range tunnelMap {
		for valveY := range tunnelMap {
			if _, ok := shortestPathMap[valveX]; !ok {
				shortestPathMap[valveX] = map[string]int{}
			}

			if contains(valveY, destTunnelsX) {
				shortestPathMap[valveX][valveY] = 1
			} else {
				shortestPathMap[valveX][valveY] = math.MaxInt
			}
		}
	}

	// Iterate through each of the valves * 3 and then every permutation will be
	//  ran through to determine the shortest path from i node to j node in the graph
	for x := range shortestPathMap {
		for y := range shortestPathMap {
			for z := range shortestPathMap {
				shortestPathMap[y][z] = int(math.Min(float64(shortestPathMap[y][z]), float64(shortestPathMap[y][x])+float64(shortestPathMap[x][z])))
			}
		}
	}

	return &rateMap, &bitfieldMap, &shortestPathMap
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func iterate(
	currentValve string,
	minutesLeft int,
	valveState int,
	flowPressure int,
	valvePressureMap *map[string]int,
	bitFieldMap *map[string]int,
	shortestPathMap *map[string]map[string]int,
	bestResults *map[int]int,
) {

	(*bestResults)[valveState] = max((*bestResults)[valveState], flowPressure)

	for valve := range *valvePressureMap {
		newMinutesLeft := minutesLeft - (*shortestPathMap)[currentValve][valve] - 1

		if (*bitFieldMap)[valve]&valveState > 0 || (newMinutesLeft <= 0) {
			continue
		}

		iterate(
			valve,
			newMinutesLeft,
			valveState|(*bitFieldMap)[valve],
			flowPressure+(newMinutesLeft*(*valvePressureMap)[valve]),
			valvePressureMap,
			bitFieldMap,
			shortestPathMap,
			bestResults,
		)
	}
}

func run(input io.Reader) (int, int) {
	// For every possible action, should recurse to next step
	// . The "30 minute limit" is a clear helper here as it can only
	// . move so deep into the stack.

	// Actions can be:
	// . 1. turn on valve
	// . 2. move to new tunnel

	// Therefore the total possible actions are
	// * Turn on valve if it's not on already
	// * Move to N tunnels

	// Should have a struct representing 'state' that is
	// . forked when either 1 or 2 happens.

	// MANY HOURS LATER

	// I ultimately found a solution on /r/AdventOfCode by /u/juanplopes
	// https://github.com/juanplopes/advent-of-code-2022/blob/main/day16.py

	// It's so great and concise on so many levels. I've translated it to golang here and
	// included a commented up version in this folder. I can't claim any of this solution as my own.

	// The two major learnings here for me are:
	// 1. Don't stray from Henry-Warshall, I wanted to believe I didn't need it but I did
	// 2. Bitfields are amazingly efficient both memorywise and computationally.
	// . and in this case made for an excellent state key

	valvePressureMap, bitfieldMap, shortestPathMap := readInput(input)

	bestResultsPartOne := map[int]int{}

	iterate("AA",
		30,
		0,
		0,
		valvePressureMap,
		bitfieldMap,
		shortestPathMap,
		&bestResultsPartOne,
	)

	bestResultsPartTwo := map[int]int{}

	iterate(
		"AA",
		26,
		0,
		0,
		valvePressureMap,
		bitfieldMap,
		shortestPathMap,
		&bestResultsPartTwo,
	)

	maxPressurePartOne := 0
	maxPressurePartTwo := 0

	for _, v := range bestResultsPartOne {
		if v > maxPressurePartOne {
			maxPressurePartOne = v
		}
	}

	for sX, vX := range bestResultsPartTwo {
		for sY, vY := range bestResultsPartTwo {
			if sX&sY == 0 && vY+vX > maxPressurePartTwo {
				maxPressurePartTwo = vY + vX
			}
		}
	}

	return maxPressurePartOne, maxPressurePartTwo
}
