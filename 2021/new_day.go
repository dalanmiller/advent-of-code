package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {

	list, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal("No dirs")
	}

	var max int
	for _, dir := range list {
		if dir.IsDir() {
			n, err := strconv.Atoi(dir.Name())
			if err != nil {
				log.Fatalf("Not convertable to int %s", dir.Name())
			}
			log.Println(n)
			if n > max {
				max = n
			}
		}
	}
	max = max + 1

	new_dir_path := fmt.Sprintf("%02d", max)
	err = os.Mkdir(new_dir_path, 0777)
	if err != nil {
		log.Fatalf("Error creating directory %s", new_dir_path)
	}
	new_file_path := fmt.Sprintf("%02d/%02d.go", max, max)
	new_test_file_path := fmt.Sprintf("%02d/%02d_test.go", max, max)
	err = os.WriteFile(new_file_path, []byte(`package main
	
	func run() int {

	}`), 0777)
	if err != nil {
		log.Fatalf("Error creating new go file %s %s", new_file_path, err)
	}
	err = os.WriteFile(new_test_file_path, []byte(`package main
	
func TestExamplesOneOne(t *testing.T) {
	tests := []struct {
		test     int
		expected int
	}{ 
		{0, 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestOneOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test int
		expected  int
	}{
		{0, 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result[0] != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesOneTwo(t *testing.T) {
	tests := []struct {
		test     int
		expected int
	}{ 
		{0, 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestOneTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test int
		expected  int
	}{
		{0, 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result[0] != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

		
	`), 0777)
	if err != nil {
		log.Fatalf("Error creating new go test file %s", new_file_path)
	}
}
