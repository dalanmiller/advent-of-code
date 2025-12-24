package year2025

import (
	"io"
	"strconv"
	"strings"
)

func Day06PartOne(r io.Reader) int {
	all, _ := io.ReadAll(r)
	lines := strings.Split(string(all), "\n")
	lines = lines[:len(lines)]

	rows := make([][]int, len(strings.Fields(lines[0])))

	total := 0
	for i, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Fields(line)

		for j, n := range split {

			if i == len(lines)-1 {

				switch n {
				case "*":
					subprod := 1
					for _, o := range rows[j] {
						subprod *= o
					}
					total += subprod
					continue
				case "+":
					subsum := 0
					for _, o := range rows[j] {
						subsum += o
					}
					total += subsum
				}
				continue
			}

			num, _ := strconv.Atoi(n)
			rows[j] = append(rows[j], num)
		}
	}

	return total
}

func Day06PartTwo(r io.Reader) int {

	// 123 328  51 64
	//  45 64  387 23
	//   6 98  215 314
	// *   +   *   +

	// The rightmost problem is 4 + 431 + 623 = 1058
	// The second problem from the right is 175 * 581 * 32 = 3253600
	// The third problem from the right is 8 + 248 + 369 = 625
	// Finally, the leftmost problem is 356 * 24 * 1 = 8544

	all, _ := io.ReadAll(r)
	lines := strings.Split(string(all), "\n")

	total := 0
	c := len(strings.Fields(lines[0]))
	nums := make([][]int, c)
	for i := range c {
		nums[i] = make([]int, 0)
	}

	sectionCount := 0

	for j, chr := range lines[0] {

		// Reset case where entire column is " "
		if chr == ' ' {
			reset := true
			for k := 1; k < len(lines); k++ {
				if lines[k][j] != ' ' {
					reset = false
					break
				}
			}

			if reset {
				sectionCount++
				continue
			}
		}

		rawNums := []string{}
		for k := 0; k < len(lines)-1; k++ {
			// n, _ := strconv.Atoi(string(chr))
			rawNums = append(rawNums, string(lines[k][j]))
		}
		num, _ := strconv.Atoi(strings.Trim(strings.Join(rawNums, ""), " "))

		nums[sectionCount] = append(nums[sectionCount], num)
	}

	operations := strings.Fields(lines[len(lines)-1])

	for i, op := range operations {
		switch op {
		case "+":
			sum := 0
			for _, n := range nums[i] {
				sum += n
			}
			total += sum
		case "*":
			prod := 1
			for _, n := range nums[i] {
				prod *= n
			}
			total += prod
		}
	}

	return total
}
