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

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// const EXAMPLE = 
	
func TestExamples{{.Day}}One(t *testing.T) {
	tests := []struct {
		test     *strings.Reader		
		expected int
	}{ 
		{strings.NewReader(EXAMPLE), 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func Test{{.Day}}One(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected  int
	}{
		{reader, 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamples{{.Day}}Two(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func Test{{.Day}}Two(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected  int
	}{
		{reader, 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}	
	
`

const MainTemplate = `package main

import (
	"io"
)
	
func readInput(input io.Reader) {
		
}	

func run(input io.Reader) int {
	readInput(input)
}`

var numberWordMap = map[string]string{
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

	mt, _ := template.New("mainFile").Parse(MainTemplate)
	file, err := os.Create(new_file_path)

	if err != nil {
		log.Fatalf("Error creating new go file %s %s", new_file_path, err)
	}

	dayString := strconv.Itoa(max)
	if len(dayString) > 1 {
		split := strings.Split(dayString, "")
		var s strings.Builder
		for _, sp := range split {
			s.Write([]byte(numberWordMap[sp]))
		}

		dayString = s.String()
	}

	err = mt.Execute(file, struct {
		Day string
	}{dayString})

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
	url, _ := url.Parse(fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", max))
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
