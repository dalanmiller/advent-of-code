package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var EXAMPLES = []struct {
	stringPair string
	outcome    bool
}{
	{`[1,1,3,1,1]
[1,1,5,1,1]`, true},
	{`[[1],[2,3,4]]
[[1],4]`, true},
	{`[9]
[[8,7,6]]`, false},
	{`[[4,4],4,4]
[[4,4],4,4,4]`, true},
	{`[7,7,7,7]
[7,7,7]`, false},
	{`[]
[3]`, true},
	{`[[[]]]
[[]]`, false},
	{`[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`, false}}

func TestExamplesThirteenOne(t *testing.T) {
	for i, example := range EXAMPLES {
		split := strings.Split(example.stringPair, "\n")
		pair := pair{
			Left:  parsePacketString(split[0]),
			Right: parsePacketString(split[1]),
			// Raw:   fmt.Sprintf("%s | %s", split[0], split[1]),
		}
		result := pair.Compare()
		if result != example.outcome {
			log.Fatalf("%d - Result %v != expected %v", i, result, example.outcome)
		}
	}
}

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

func TestParsePacketString(t *testing.T) {
	tests := []struct {
		test     string
		expected *packetData
	}{
		{"[]", &packetData{
			DataType: LIST,
			List:     []*packetData{},
		}},
		{"[3]", &packetData{
			DataType: LIST,
			List:     []*packetData{{DataType: INTEGER, Integer: 3}},
		}},

		{"[[[]]]", &packetData{
			DataType: LIST,
			List: []*packetData{
				{DataType: LIST, List: []*packetData{{DataType: LIST, List: []*packetData{}}}}},
		}},
		{"[[]]", &packetData{
			DataType: LIST,
			List: []*packetData{
				{DataType: LIST, List: []*packetData{}},
			},
		}},
	}

	for i, test := range tests {
		result := parsePacketString(test.test)
		if !cmp.Equal(result, test.expected) {
			t.Fatalf("%d | Result %v != expected %v", i, result, test.expected)
		}
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
