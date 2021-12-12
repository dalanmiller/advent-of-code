package main

import (
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		test     string
		expected Cave
	}{
		{`start-end`,
			Cave{
				Name: "start",
				Size: 'S',
			},
		},
	}

	for _, test := range tests {
		caveMap := make(map[string]*Cave)
		result := parseInput(test.test, caveMap)
		if result.Name != "start" {
			t.Fatal("Wrong name")
		}

		if len(result.Connections) != 1 {
			t.Fatal("Wrong number connections")
		}

		if result.Size == 'L' {
			t.Fatal("Incorrect size")
		}
	}
}

func TestSpelunk(t *testing.T) {

	start := Cave{
		Name: "start",
		Size: 'S',
	}

	TEST1 := Cave{
		Name: "TEST1",
		Size: 'L',
	}

	mid := Cave{
		Name: "mid",
		Size: 'S',
	}

	TEST2 := Cave{
		Name: "TEST2",
		Size: 'L',
	}

	end := Cave{
		Name: "end",
		Size: 'S',
	}

	start.Connections = append(start.Connections, &TEST1, &TEST2)
	TEST1.Connections = append(TEST1.Connections, &start, &mid, &end)
	TEST2.Connections = append(TEST2.Connections, &start, &mid, &end)
	mid.Connections = append(mid.Connections, &TEST1, &TEST2)
	end.Connections = append(end.Connections, &TEST2, &TEST1)

	var validPaths []Path
	spelunk(start, Path{}, &validPaths, 1)

	if len(validPaths) != 6 {
		t.Fatalf("Len of paths expected %d, got %d", 6, len(validPaths))
	}

	for _, path := range validPaths {
		pl := len(path)
		name := path[pl-1].Name
		if name != "end" {
			t.Fatalf("Last node in path is not 'end', got %s", name)
		}
	}

}

func TestExamplesTwelveOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`start-TEST1
start-TEST2
TEST1-mid
TEST1-end
TEST2-mid
TEST2-end`, 6},
		{`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`, 19},
		{`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`, 226},
	}

	for _, test := range tests {
		result := run(test.test, 1)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwelveOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 4167},
	}

	for _, test := range tests {
		result := run(test.test, 1)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesTwelveTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`start-A
start-b
A-c
A-b
b-d
A-end
b-end`, 36},
		{`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`, 103},
		{`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`, 3509},
	}

	for _, test := range tests {
		result := run(test.test, 2)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwelveTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 98441},
	}

	for _, test := range tests {
		result := run(test.test, 2)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
