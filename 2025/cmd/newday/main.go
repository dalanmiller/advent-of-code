package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

const (
	year        = 2025
	packageName = "year2025"
)

var moduleRoot = resolveModuleRoot()

var (
	goFileTemplate = template.Must(template.New("goFile").Parse(`package {{.Package}}

import (
	"bufio"
	"io"
)

func Day{{.DayPadded}}PartOne(r io.Reader) int {
	s := bufio.NewScanner(r)
	_ = s
	return 0
}

func Day{{.DayPadded}}PartTwo(r io.Reader) int {
	s := bufio.NewScanner(r)
	_ = s
	return 0
}
`))

	testFileTemplate = template.Must(template.New("testFile").Parse(`package {{.Package}}

import (
	"strings"
	"testing"
)

const day{{.DayPadded}}Example = ""

func TestDay{{.DayPadded}}PartOneExample(t *testing.T) {
	AssertEqual(t, Day{{.DayPadded}}PartOne, strings.NewReader(day{{.DayPadded}}Example), 0)
}

func TestDay{{.DayPadded}}PartOneInput(t *testing.T) {
	AssertEqual(t, Day{{.DayPadded}}PartOne, ReaderForInput({{.Day}}), 0)
}

func TestDay{{.DayPadded}}PartTwoExample(t *testing.T) {
	AssertEqual(t, Day{{.DayPadded}}PartTwo, strings.NewReader(day{{.DayPadded}}Example), 0)
}

func TestDay{{.DayPadded}}PartTwoInput(t *testing.T) {
	AssertEqual(t, Day{{.DayPadded}}PartTwo, ReaderForInput({{.Day}}), 0)
}
`))
)

type templateData struct {
	Day       int
	DayPadded string
	Package   string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("info: .env not loaded: %v", err)
	}

	day := detectNextDay()
	if len(os.Args) > 1 {
		override, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("day override must be an integer: %v", err)
		}
		day = override
	}

	if day < 1 || day > 25 {
		log.Fatalf("day must be between 1 and 25; got %d", day)
	}

	td := templateData{
		Day:       day,
		DayPadded: fmt.Sprintf("%02d", day),
		Package:   packageName,
	}

	writeFileOrExit(filepath.Join(moduleRoot, fmt.Sprintf("day%02d.go", day)), goFileTemplate, td)
	writeFileOrExit(filepath.Join(moduleRoot, fmt.Sprintf("day%02d_test.go", day)), testFileTemplate, td)

	inputPath := filepath.Join(moduleRoot, "inputs", fmt.Sprintf("day%02d.txt", day))
	if _, err := os.Stat(inputPath); err == nil {
		log.Printf("input already exists at %s; skipping download", inputPath)
		return
	}

	if err := os.MkdirAll(filepath.Dir(inputPath), 0o755); err != nil {
		log.Fatalf("failed to create inputs directory: %v", err)
	}

	if err := downloadInput(day, inputPath); err != nil {
		log.Printf("warning: could not download input: %v", err)
		log.Printf("created empty file instead; fill %s manually", inputPath)
		_ = os.WriteFile(inputPath, []byte{}, 0o644)
	}
}

func detectNextDay() int {
	files, err := filepath.Glob(filepath.Join(moduleRoot, "day??.go"))
	if err != nil {
		log.Fatalf("unable to scan for existing days: %v", err)
	}

	re := regexp.MustCompile(`day(\d{2})\.go$`)
	highest := 0
	for _, f := range files {
		matches := re.FindStringSubmatch(f)
		if len(matches) != 2 {
			continue
		}
		n, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}
		if n > highest {
			highest = n
		}
	}
	return highest + 1
}

func writeFileOrExit(path string, tmpl *template.Template, td templateData) {
	if _, err := os.Stat(path); err == nil {
		log.Fatalf("file %s already exists", path)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("unable to create %s: %v", path, err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, td); err != nil {
		log.Fatalf("unable to render %s: %v", path, err)
	}
}

func resolveModuleRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("runtime caller lookup failed")
	}
	// ../.. from cmd/newday/main.go back to module root
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
}

func downloadInput(day int, dest string) error {
	session := strings.TrimSpace(os.Getenv("AOC_SESSION"))
	if session == "" {
		session = strings.TrimSpace(os.Getenv("session"))
	}
	if session == "" {
		return fmt.Errorf("missing AOC_SESSION; skipping download")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("bad response (%d): %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading body: %w", err)
	}

	if err := os.WriteFile(dest, body, 0o644); err != nil {
		return fmt.Errorf("writing input file: %w", err)
	}
	return nil
}
