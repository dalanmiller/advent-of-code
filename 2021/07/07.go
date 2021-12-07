package main

import (
	"math"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	crabs := strings.Split(input, ",")

	crabs_values := make([]int, len(crabs))
	var max int
	for i, v := range crabs {
		n, _ := strconv.Atoi(v)
		crabs_values[i] = n

		if n > max {
			max = n
		}
	}

	field := make([]int, max+1)
	for _, cv := range crabs_values {
		field[cv]++
	}

	return field
}

func calculateFuelRequirement(distance int, numCrabs int, memo map[int]int) int {
	if val, ok := memo[distance]; ok {
		return val * numCrabs
	}

	// Sum of all values between 1 -> distance
	v := distance * (distance + 1) / 2

	memo[distance] = v
	return v * numCrabs
}

func run(input string) (int, int) {
	field := parseInput(input)

	// possible_positions
	part_one_fuel_required := make([]int, len(field))
	part_two_fuel_required := make([]int, len(field))
	memo := make(map[int]int)
	for i := range field {
		sum_one, sum_two := 0, 0

		for j := range field {

			// Continue given the crabs in that location don't need to go anywhere
			if i == j {
				continue
			}

			distance := int(math.Abs(float64(i - j)))
			// Distance between target position and current position
			//  times number of crab ships

			// Part 1 calculation
			sum_one += distance * field[j]

			// Part 2 calculation
			// TIL maps are a reference value in Golang and don't need
			//  to have a pointer.
			sum_two += calculateFuelRequirement(distance, field[j], memo)

		}

		part_one_fuel_required[i] = sum_one
		part_two_fuel_required[i] = sum_two
	}

	min_part_one_fuel_cost := math.MaxInt
	min_part_two_fuel_cost := math.MaxInt
	for i := range part_one_fuel_required {
		if part_one_fuel_required[i] < min_part_one_fuel_cost {
			min_part_one_fuel_cost = part_one_fuel_required[i]
		}

		if part_two_fuel_required[i] < min_part_two_fuel_cost {
			min_part_two_fuel_cost = part_two_fuel_required[i]
		}
	}
	return min_part_one_fuel_cost, min_part_two_fuel_cost
}
