package year2025

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func Day01PartOne(r io.Reader) int {
	s := bufio.NewScanner(r)
	position := 50
	count := 0
	for s.Scan() {
		line := s.Text()
		direction := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Oops")
		}

		switch direction {
		case 'L':
			position = mod(position-clicks, 100)
		case 'R':
			position = mod(position+clicks, 100)
		}

		// log.Printf("%v:%d -> %d", string(direction), clicks, position)
		if position == 0 {
			// log.Printf("Count: %d", count)
			count += 1
		}

	}
	return count
}

func Day01PartTwo(r io.Reader) int {
	s := bufio.NewScanner(r)
	position := 50
	count := 0
	for s.Scan() {
		line := s.Text()
		direction := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Oops")
		}

		hits := 0
		switch direction {
		case 'L':
			hits += ((100-position)%100 + clicks) / 100
			old_position := position
			position = mod(position-clicks, 100)
			log.Printf("%d : -%d : %d", old_position, clicks, position)
		case 'R':
			hits = (position + clicks) / 100
			old_position := position
			position = mod(position+clicks, 100)
			log.Printf("%d : +%d : %d", old_position, clicks, position)
		}

		count += hits
	}
	return count
}
