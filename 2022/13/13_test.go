package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestExampleThirteen(t *testing.T) {
	test := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

	result := run(strings.NewReader(test))
	if result != 13 {
		log.Fatalf("Result %d != expected 13", result)
	}

}

func TestThirteenOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		// 3650 too low
		// 4442 too low
		// 6251 too high
		{reader, 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestExamplesThirteenTwo(t *testing.T) {
// 	tests := []struct {
// 		test     *strings.Reader
// 		expected int
// 	}{
// 		{strings.NewReader(EXAMPLE), 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if test.expected != result {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }

// func TestThirteenTwo(t *testing.T) {
// 	file, _ := os.Open("./input")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)

// 	tests := []struct {
// 		test     *bufio.Reader
// 		expected int
// 	}{
// 		{reader, 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if result != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }
