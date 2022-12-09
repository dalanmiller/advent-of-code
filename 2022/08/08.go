package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type treeGrid [][]int

func (tg *treeGrid) determineScenicity(row, col int) int {

	if row == 3 && col == 2 {
		fmt.Println("hi")
	}

	h := (*tg)[row][col]
	// top
	i, a := row-1, 0
	for i >= 0 && i < len((*tg)) {
		if (*tg)[i][col] < h {
			a++
			i--
			continue
		} else if (*tg)[i][col] == h {
			a++
			i++
		}
		break
	}

	// right
	i, b := col+1, 0
	for i >= 0 && i < len((*tg)) {
		if (*tg)[row][i] < h {
			b++
			i++
			continue
		} else if (*tg)[row][i] >= h {
			b++
			i++
		}
		break
	}

	// bottom
	i, c := row+1, 0
	for i >= 0 && i < len((*tg)) {
		if (*tg)[i][col] < h {
			c++
			i++
			continue
		} else if (*tg)[i][col] == h {
			c++
			i++
		}
		break
	}

	// left
	i, d := col-1, 0
	for i >= 0 && i < len((*tg)) {
		if (*tg)[row][i] < h {
			d++
			i--
			continue
		} else if (*tg)[row][i] == h {
			d++
			i++
		}
		break
	}

	return a * b * c * d
}

func (tg *treeGrid) determineVisibility(row, col int) bool {

	// Detect edges
	// if (row == 0 || row == len((*tg)[0])-1) || (col == 0 || col == len((*tg))-1) {
	// 	return true
	// }

	// Can't be equal to or higher than this one
	h := (*tg)[row][col]

	// Scan rows
	i := 0
	for {
		if i == col || i == len((*tg)[0]) {
			// We made it! This means we didn't skip ahead or break which means
			// the tree is visible
			return true
		}

		// Check if higher
		if (*tg)[row][i] >= h {
			// Check if left or right of tree
			// If left, then we can fast forward to right of tree
			// If right, we can break and move onto column check
			if i < col {
				i = col + 1
			} else {
				break
			}
		} else {
			i++
		}
	}

	i = 0
	for {
		if i == row || i == len((*tg)[0]) {
			// We made it! This means we didn't skip ahead or break which means
			// the tree is visible
			return true
		}

		// Check if higher
		if (*tg)[i][col] >= h {
			// Check if above or below tree
			// If above, then we can fast forward to just below tree
			// If below, we can break and fail
			if i < row {
				i = row + 1
			} else {
				break
			}
		} else {
			i++
		}
	}

	return false
}

func run(input io.Reader) (int, int) {

	s := bufio.NewScanner(input)

	grid := treeGrid{[]int{}}
	grid[0] = []int{}

	row, column := 0, 0

	// Fill grid
	for s.Scan() {
		line := s.Text()

		for _, r := range line {
			n, _ := strconv.Atoi(string(r))
			grid[row] = append(grid[row], n)
			column++
		}
		row++
		column = 0

		grid = append(grid, make([]int, 0, len(line)))
	}

	// Don't know how to tell if I'm scanning the last line /o\
	grid = grid[0 : len(grid)-1]

	// We already know that the outer trees are visible
	n := len(grid)*4 - 4

	// Scenic score max
	scenicScore := 0

	// We can start on the inner square
	for i := 1; i <= len(grid)-2; i++ {
		for j := 1; j <= len(grid[0])-2; j++ {
			if grid.determineVisibility(i, j) {
				n++
			}

			if curS := grid.determineScenicity(i, j); curS > scenicScore {
				scenicScore = curS
			}
		}
	}

	return n, scenicScore
}
