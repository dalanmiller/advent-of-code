package main

import (
	"strconv"
	"strings"
)

type Board struct {
	square [][]int
	marked map[int]bool
	won    bool
}

func (b *Board) score(last_call int) int {
	sum := 0
	for n, marked := range b.marked {
		if !marked {
			sum += n
		}
	}
	return sum * last_call
}

func (b *Board) check_winner() bool {
	// Check rows
	for _, row := range b.square {
		row_marked := true
		for _, l := range row {
			if !b.marked[l] {
				row_marked = false
				break
			}
		}

		if row_marked {
			return true
		}
	}
	// Check columns
	for i := 0; i < len(b.square); i++ {
		column_marked := true
		for j := 0; j < len(b.square); j++ {
			if !b.marked[b.square[j][i]] {
				column_marked = false
				break
			}
		}

		if column_marked {
			return true
		}
	}

	// Check diagonol
	// i, j, y, z := 0, len(b.square)-1, 0, 0
	// top_left_diagonol _marked := true
	// bottom_left_diagonol_marked := true
	// for i < len(b.square) {
	// 	if !b.marked[b.square[i][j]] {
	// 		top_left_diagonol _marked = false
	// 	}

	// 	if !b.marked[b.square[y][z]] {
	// 		bottom_left_diagonol _marked = false
	// 	}

	// 	i++
	// 	j--
	// 	y++
	// 	z++
	// }

	// if top_left_diagonol _marked || bottom_left_diagonol _marked {
	// 	return true
	// }

	return false
}

func parse_square_input(input []string) *Board {
	square := make([][]int, len(input))
	marked := make(map[int]bool)
	for i, row := range input {
		square[i] = make([]int, len(input))
		string_numbers := strings.Fields(row)
		for j, number := range string_numbers {
			n, _ := strconv.Atoi(string(number))
			square[i][j] = n
			marked[n] = false
		}
	}

	return &Board{
		square: square,
		marked: marked,
	}
}

func run(input string) (int, int) {

	lines := strings.Split(input, "\n")
	ns := strings.Split(lines[0], ",")
	marked_numbers := make([]int, len(ns))
	for i, char := range ns {
		n, _ := strconv.Atoi(string(char))
		marked_numbers[i] = n
	}

	boards := []*Board{}
	current_square := []string{}
	for i, line := range lines[2:] {
		if line == "" {
			boards = append(boards, parse_square_input(current_square))
			current_square = []string{}
		} else {
			treated_line := strings.TrimRight(string(line), " ")
			current_square = append(current_square, treated_line)
		}

		if len(lines[2:])-1 == i {
			boards = append(boards, parse_square_input(current_square))
		}
	}

	// Now we check the numbers
	boards_without_wins := len(boards)
	var first, last int
	for _, n := range marked_numbers {
		for _, board := range boards {
			if board.won {
				continue
			}
			board.marked[n] = true
			if board.check_winner() {

				if boards_without_wins == 1 {
					last = board.score(n)
					return first, last
				}

				// Set board to won
				board.won = true

				// If we haven't found a first yet
				if first == 0 {

					first = board.score(n)
					boards_without_wins--
				} else {
					boards_without_wins--
				}
			}
		}
	}

	return first, last
}
