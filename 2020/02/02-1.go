package main

import (
	"fmt"
	"strconv"
	"strings"

	. "../helpers"
)

func main() {
	var lines []string
	var err error
	if lines, err = ReadInput("./input"); err == nil {
		fmt.Printf("something bad", err)
	}

	validPasswords := 0

	for _, line := range lines {
		split_line := strings.Split(line, " ")

		charRange := split_line[0]
		keyChar := []rune(strings.Replace(split_line[1], ":", "", -1))
		password := strings.Trim(split_line[2], " \n")

		splitCharRange := strings.Split(charRange, "-")
		minChars, _ := strconv.Atoi(splitCharRange[0])
		maxChars, _ := strconv.Atoi(splitCharRange[1])

		count := 0
		for _, char := range password {
			if char == keyChar[0] {
				count++
			}
		}

		if count >= minChars && count <= maxChars {
			validPasswords++
		}
	}
	fmt.Printf("Valid Passwords: %d", validPasswords)
}
