package year2025

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Year is the Advent of Code season this package targets.
const Year = 2025

var baseDir string

func init() {
	// Make all input helpers robust to the current working directory by
	// anchoring paths off of this source file.
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller failed: unable to locate source file path")
	}
	baseDir = filepath.Dir(filename)
}

// InputPath returns the absolute path for a day's input.
func InputPath(day int) string {
	return filepath.Join(baseDir, "inputs", fmt.Sprintf("day%02d.txt", day))
}

// MustReadInput reads a day's input file or panics with a helpful error.
func MustReadInput(day int) string {
	raw, err := os.ReadFile(InputPath(day))
	if err != nil {
		panic(fmt.Sprintf("read input for day %02d: %v", day, err))
	}
	return string(raw)
}

// ReaderForInput creates a reader over the day's input.
func ReaderForInput(day int) io.Reader {
	return bytes.NewReader([]byte(MustReadInput(day)))
}

// TrimmedLines splits the input on newlines and trims trailing whitespace.
func TrimmedLines(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	if len(lines) == 1 && lines[0] == "" {
		return []string{}
	}
	return lines
}

// AssertEqual is a tiny test helper to compare solver results.
func AssertEqual(t *testing.T, solver func(io.Reader) int, input io.Reader, expected int) {
	t.Helper()
	if got := solver(input); got != expected {
		t.Fatalf("got %d, expected %d", got, expected)
	}
}

func AssertEqual8(t *testing.T, solver func(io.Reader, int) int, input io.Reader, connections int, expected int) {
	t.Helper()
	if got := solver(input, connections); got != expected {
		t.Fatalf("got %d, expected %d", got, expected)
	}
}

func AssertEqual64(t *testing.T, solver func(io.Reader) int64, input io.Reader, expected int64) {
	t.Helper()
	if got := solver(input); got != expected {
		t.Fatalf("got %d, expected %d", got, expected)
	}
}
