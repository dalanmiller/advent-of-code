package main

import (
	"log"
	"strconv"
	"strings"
)

func ADD(curpos int, array []int){
	array[array[curpos + 3]] = array[array[curpos + 1]] + array[array[curpos + 2]]
} 

func MULT(curpos int, array []int){
	array[array[curpos + 3]] = array[array[curpos + 1]] * array[array[curpos + 2]]
}

func run(input []byte) []int {		
		line := string(input[:])
		numbers := strings.Split(line, ",")

		var opcodes []int
		for _, b := range numbers {
			val, err := strconv.Atoi(string(b))
			if err != nil {
				log.Fatalf("Could not convert string to int %s", err)
			}
			opcodes = append(opcodes, int(val))
		}

		for i, char := range opcodes {
			
			if i % 4 != 0{
				continue
			}
				
			switch char {
			case 1: 
				ADD(i, opcodes)	
			case 2:
				MULT(i, opcodes)		
			}
		}
		return opcodes
}

func simulate(opcodes []int) int {
	for i, char := range opcodes {
		
		if i % 4 != 0{
			continue
		}
		
		// if i+3 > len(opcodes) || opcodes[i + 3] > len(opcodes) {
		// 	return -1
		// }
		
		switch char {
		case 1: 
			ADD(i, opcodes)	
		case 2:
			MULT(i, opcodes)		
		}
	}

	return opcodes[0]
}

func run_two(input []byte, expected int) (int, int) {
	line := string(input[:])
	numbers := strings.Split(line, ",")

	var opcodes []int
	for _, b := range numbers {
		val, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatalf("Could not convert string to int %s", err)
		}
		opcodes = append(opcodes, int(val))
	}

	var final_noun int
	var final_verb int
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			
			simulation_codes := make([]int, len(opcodes))
			copy(simulation_codes, opcodes)

			simulation_codes[1] = noun
			simulation_codes[2] = verb 

			result := simulate(simulation_codes)			
			
			if result == expected {
				final_noun = noun
				final_verb = verb
				break
			}
		}		
	}
	return final_noun, final_verb
}
