package year2025

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Day03PartOne(r io.Reader) int {
	s := bufio.NewScanner(r)

	joltage := 0
	for s.Scan() {
		line := s.Text()
		batteries := strings.Split(line, "")

		max := 0
		for i := 0; i < len(batteries); i++ {
			for j := len(batteries) - 1; j > i; j-- {
				left, _ := strconv.Atoi(batteries[i])
				right, _ := strconv.Atoi(batteries[j])
				current := (left * 10) + right
				// log.Println(current)
				if current > max {
					max = current
				}
			}
		}
		// log.Printf("Max: %d", max)
		joltage += max
	}

	return joltage
}

func maxAfterRemovingK(line string, k int) int {

	// Make a stack to store largest seen
	stack := make([]byte, 0, len(line))

	// Iterate through indexes of line
	for i := 0; i < len(line); i++ {

		// Keep line at i in d
		d := line[i] // 1, 2

		// While:
		//  * We still have k erasures left
		//  * Our stack isn't empty
		//  * And the element stack at the end is less than what is currenty in d
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] < d {

			// Remove last element in the stack
			stack = stack[:len(stack)-1]
			// Reduce k by one
			k--
		}
		// Append to stack d
		stack = append(stack, d)
	}

	// At the end if there's still k erasures left
	//  let's remove them from the end where the lowest value digits are
	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	jolts := 0
	for _, b := range stack {
		// Nice golang trick to get the integer value of the string numberr
		jolts = jolts*10 + int(b-'0')
	}
	return jolts
}

func Day03PartTwo(r io.Reader) int64 {
	s := bufio.NewScanner(r)

	joltage := int64(0)
	for s.Scan() {
		line := s.Text()
		max := maxAfterRemovingK(line, len(line)-12)
		joltage += int64(max)
	}

	return joltage
}
