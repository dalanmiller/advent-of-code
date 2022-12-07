package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestExamplesSevenOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 95437},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSevenOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 1583951},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesSevenTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 24933642},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSevenTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 214171},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
