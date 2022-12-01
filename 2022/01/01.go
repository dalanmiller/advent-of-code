package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type elf struct {
	Calories int
	Items    []int
}

func readFile(input string) []*elf {
	lines := strings.Split(input, "\n")

	elves := []*elf{&elf{0, []int{}}}
	for _, line := range lines {
		if line == "" {
			elves = append(elves, &elf{0, []int{}})
			continue
		}

		e := elves[len(elves)-1]
		v, err := strconv.Atoi(line)

		if err != nil {
			log.Fatalf("Failed to parse int string %s", err)
		}

		e.Calories += v
		e.Items = append(e.Items, v)
	}

	return elves
}

func run(input string) (int, int) {
	elves := readFile(input)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})

	return elves[0].Calories, (elves[0].Calories + elves[1].Calories + elves[2].Calories)
}
