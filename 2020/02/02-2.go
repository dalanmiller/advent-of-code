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
		firstPosition, _ := strconv.Atoi(splitCharRange[0])
		secondPosition, _ := strconv.Atoi(splitCharRange[1])

		first := rune(password[firstPosition-1]) == keyChar[0]
		second := rune(password[secondPosition-1]) == keyChar[0]

		if !(first && second) && (first || second) {
			validPasswords++
		}

	}
	fmt.Printf("Valid Passwords: %d", validPasswords)
}
