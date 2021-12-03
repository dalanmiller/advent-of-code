package main

import (
	"log"
	"strconv"
	"strings"
)

func parse_input(input string) [][]int {
	lines := strings.Split(input, "\n")
	numbers := make([][]int, len(lines))
	for i, line := range lines {
		digits := strings.Split(line, "")
		for _, digit := range digits {
			n, _ := strconv.Atoi(digit)
			numbers[i] = append(numbers[i], n)
		}
	}

	return numbers
}

func run(input string) int {

	numbers := parse_input(input)

	count := make(map[int]map[int]int)
	for _, number := range numbers {
		for j, n := range number {
			_, ok := count[j]
			if !ok {
				count[j] = make(map[int]int)
			}
			count[j][n]++
		}
	}

	positions := make([]int, len(numbers[0]))
	for i := range positions {
		positions[i] = i
	}

	gamma_array := make([]int, len(positions))
	epsilon_array := make([]int, len(positions))
	for i, pos := range positions {
		digit_counts := count[pos]

		if digit_counts[0] > digit_counts[1] {

			gamma_array[i], epsilon_array[i] = 0, 1
		} else {
			gamma_array[i], epsilon_array[i] = 1, 0
		}
	}

	gamma_string_array := make([]string, len(gamma_array))
	epsilon_string_array := make([]string, len(gamma_array))
	for i := range gamma_array {
		gamma_string_array = append(gamma_string_array, strconv.Itoa(gamma_array[i]))
		epsilon_string_array = append(epsilon_string_array, strconv.Itoa(epsilon_array[i]))
	}

	gamma_string := strings.Join(gamma_string_array, "")
	epsilon_string := strings.Join(epsilon_string_array, "")

	gamma, err := strconv.ParseInt(gamma_string, 2, 32)
	if err != nil {
		log.Fatalf("Unable to parse gamma, %s", err)
	}
	epsilon, err := strconv.ParseInt(epsilon_string, 2, 32)
	if err != nil {
		log.Fatalf("Unable to parse epsilon, %s", err)
	}

	return int(gamma * epsilon)
}

func counts(pos int, lines [][]int) (int, int) {
	var ones, zeroes int
	for _, line := range lines {
		if line[pos] == 1 {
			ones++
		} else {
			zeroes++
		}
	}

	if ones == zeroes {
		return 1, 0
	} else if ones > zeroes {
		return 1, 0
	} else {
		return 0, 1
	}
}

func filter_numbers(pos int, digit int, rows [][]int) [][]int {
	var temp [][]int
	for _, number := range rows {
		if number[pos] == digit {
			temp = append(temp, number)
		}
	}
	return temp
}

func run_two(input string) int {

	numbers := parse_input(input)

	possible_oxygen_numbers := make([][]int, len(numbers))
	possible_co2_numbers := make([][]int, len(numbers))
	copy(possible_oxygen_numbers, numbers)
	copy(possible_co2_numbers, numbers)

	var n, max, min int
	for len(possible_co2_numbers) != 1 || len(possible_oxygen_numbers) != 1 {
		max, _ = counts(n, possible_oxygen_numbers)
		if len(possible_oxygen_numbers) > 1 {
			possible_oxygen_numbers = filter_numbers(n, max, possible_oxygen_numbers)
		}

		_, min = counts(n, possible_co2_numbers)
		if len(possible_co2_numbers) > 1 {
			possible_co2_numbers = filter_numbers(n, min, possible_co2_numbers)
		}

		n++
	}

	oxygen_string_array := make([]string, len(possible_oxygen_numbers[0]))
	co2_string_array := make([]string, len(possible_co2_numbers[0]))

	for i, n := range possible_oxygen_numbers[0] {
		oxygen_string_array[i] = strconv.Itoa(n)
		co2_string_array[i] = strconv.Itoa(possible_co2_numbers[0][i])
	}

	oxygen_generator_rating, _ := strconv.ParseInt(strings.Join(oxygen_string_array, ""), 2, 32)
	co2_scrubber_rating, _ := strconv.ParseInt(strings.Join(co2_string_array, ""), 2, 32)

	return int(co2_scrubber_rating * oxygen_generator_rating)
}
