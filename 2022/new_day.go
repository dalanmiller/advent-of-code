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
	"text/template"

	"github.com/joho/godotenv"
)

const TEST_TEMPLATE = `
package main
	
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
}	
	
`

const MainTemplate = `package main
	
func readInput(input string) {
		
}	

func run(input string) int {
	// things = readInput(input)

}`

var NumberWordMap = map[string]string{
	"1": "One",
	"2": "Two",
	"3": "Three",
	"4": "Four",
	"5": "Five",
	"6": "Six",
	"7": "Seven",
	"8": "Eight",
	"9": "Nine",
}

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

	mt, err := template.New("mainFile").Parse(MainTemplate)
	file, err := os.Create(new_file_path)

	if err != nil {
		log.Fatalf("Error creating new go file %s %s", new_file_path, err)
	}

	err = mt.Execute(file, nil)

	if err != nil {
		log.Fatalf("Error writing template to new go main file %s", err)
	}

	tt, err := template.New("testFile").Parse(TEST_TEMPLATE)
	file, err = os.Create(new_test_file_path)

	if err != nil {
		log.Fatalf("Could not create new test file")
	}

	err = tt.Execute(file, nil)

	if err != nil {
		log.Fatalf("Error writing template to new go test file %s", err)
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
