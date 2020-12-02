package helpers

import (
	"bufio"
	"os"
)

func ReadInput(path string) ([]string, error) {
	var lines []string
	var err error

	var file *os.File
	if file, err = os.Open(path); err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
