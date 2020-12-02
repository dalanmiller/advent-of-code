package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Solution struct {
	x int
	y int
}

func main() {
	file, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Print(err)
	}

	fileString := string(file[:])
	stringList := strings.Split(fileString, "\n")

	var numbers []int
	for _, item := range stringList {

		if string(item) == "" {
			continue
		}

		value, err := strconv.Atoi(string(item))
		if err != nil {
			fmt.Printf("ERROR")
		}
		numbers = append(numbers, value)
	}

	var solutions []Solution
	for _, numberX := range numbers {
		for _, numberY := range numbers {
			if numberX+numberY == 2020 {
				fmt.Printf("x ", fmt.Sprint(numberX), " | y ", fmt.Sprint(numberY))
				solutions = append(solutions, Solution{x: numberX, y: numberY})
			}
		}
	}

	for _, solution := range solutions {
		fmt.Println(fmt.Sprint(solution.x * solution.y))
	}
	fmt.Println()
	fmt.Println("Done")
}
