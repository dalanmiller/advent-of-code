package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func TestExamplesSixteenOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 1651},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixteenOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 1737},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// BenchmarkSixteenOne-10    	       5	 203341133 ns/op	 1398945 B/op	    1265 allocs/op
// BenchmarkSixteenOne-10    	       5	 206798108 ns/op	 1412169 B/op	    1261 allocs/op
// BenchmarkSixteenOne-10    	       2	 552934834 ns/op	 2261488 B/op	    1775 allocs/op

func BenchmarkSixteenOne(b *testing.B) {
	for x := 0; x <= b.N; x++ {
		b.StopTimer()
		file, _ := os.Open("./input")
		defer file.Close()
		reader := bufio.NewReader(file)

		tests := []struct {
			test     *bufio.Reader
			expected int
		}{
			{reader, 1737},
		}
		b.StartTimer()
		result, _ := run(tests[0].test)
		b.StopTimer()
		if result != tests[0].expected {
			b.Fatalf("Result %d != expected %d", result, tests[0].expected)
		}
	}
}

func TestExamplesSixteenTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 1707},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixteenTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 2216},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
