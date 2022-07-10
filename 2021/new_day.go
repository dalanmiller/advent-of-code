package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	list, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal("No dirs")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var max int
	for _, dir := range list {
		if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") {
			n, err := strconv.Atoi(dir.Name())
			if err != nil {
				log.Fatalf("Not convertable to int %s", dir.Name())
			}
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
	
	func run(input string) int {

	}`), 0777)
	if err != nil {
		log.Fatalf("Error creating new go file %s %s", new_file_path, err)
	}
	err = os.WriteFile(new_test_file_path, []byte(`package main
	
func TestExamplesOneOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{ 
		{"", 0},
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
		test string
		expected  int
	}{
		{string(file), 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesOneTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{ 
		{"", 0},
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
		test string
		expected  int
	}{
		{string(file), 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result[0] != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}`), 0777)
	if err != nil {
		log.Fatalf("Error creating new go test file %s", new_file_path)
	}

	header := http.Header{}
	header.Set("Cookie", fmt.Sprintf("session=%s", os.Getenv("session")))
	url, _ := url.Parse(fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", max))
	request := &http.Request{
		URL:    url,
		Header: header,
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("Unable to request input file")
	}

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	input_file_path := fmt.Sprintf("%02d/input", max)
	err = os.WriteFile(input_file_path, body, 0777)
	if err != nil {
		log.Fatalf("Unable to write input resp body to input file")
	}
}
