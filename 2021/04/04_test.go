package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestSquareParse(t *testing.T) {

	expected_board := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19}}

	expected_map := make(map[int]bool)
	for i, row := range expected_board {
		for j, _ := range row {
			expected_map[expected_board[i][j]] = false
		}
	}

	tests := []struct {
		test     []string
		expected Board
	}{
		{[]string{
			"22 13 17 11  0",
			"8  2 23  4 24",
			"21  9 14 16  7",
			"6 10  3 18  5",
			"1 12 20 15 19",
		},
			Board{
				square: expected_board,
				marked: expected_map,
			},
		},
	}

	for _, test := range tests {
		result := parse_square_input(test.test)
		if !reflect.DeepEqual(result.square, test.expected.square) {
			t.Fatalf("Result board %v != expected %v", result, test.expected)
		}

		if !reflect.DeepEqual(result.marked, test.expected.marked) {
			t.Fatalf("Result marked map %v != expected %v", result, test.expected)
		}
	}
}

func TestBoardWinner(t *testing.T) {
	expected_board := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19}}

	expected_map := make(map[int]bool)
	for i, row := range expected_board {
		for j, _ := range row {
			expected_map[expected_board[i][j]] = false
		}
	}

	tests := []struct {
		test     Board
		expected bool
	}{
		{
			test: Board{
				square: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19}},
				marked: map[int]bool{
					22: true,
					13: true,
					17: true,
					11: true,
					0:  true,
				},
			},
			expected: true,
		},
		// {
		// 	test: board{
		// 		square: [][]int{
		// 			{22, 13, 17, 11, 0},
		// 			{8, 2, 23, 4, 24},
		// 			{21, 9, 14, 16, 7},
		// 			{6, 10, 3, 18, 5},
		// 			{1, 12, 20, 15, 19},
		// 		},
		// 		marked: map[int]bool{
		// 			1:  true,
		// 			10: true,
		// 			14: true,
		// 			4:  true,
		// 			0:  true,
		// 		},
		// 	},
		// 	expected: true,
		// },
		{
			test: Board{
				square: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				marked: map[int]bool{},
			},
			expected: false,
		},
		{
			test: Board{
				square: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				marked: map[int]bool{
					0:  true,
					24: true,
					7:  true,
					5:  true,
					19: true,
				},
			},
			expected: true,
		},
		{
			test: Board{
				square: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				marked: map[int]bool{
					17: true,
					23: true,
					14: true,
					3:  true,
					20: true,
				},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		result := test.test.check_winner()
		if result != test.expected {
			log.Fatalf("board check failing, expected %v, %v, %v", test.expected, test.test.square, test.test.marked)
		}
	}

}

func TestExamplesFourOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
			4512,
		},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFourOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 32844},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesFourTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
			1924,
		},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFourTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 4920},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
