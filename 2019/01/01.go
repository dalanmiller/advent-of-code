package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func Calculate_fuel_requirements() []int {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fuel_requirements []int
	for scanner.Scan() {
		mass_int, err := strconv.Atoi(scanner.Text())
		mass_float := float64(mass_int)
		if err != nil {
			log.Fatal(err)
		}
		fuel_required := math.Floor(mass_float/3) - 2
		// log.Printf("%.2f => %.2f", mass_float, fuel_required)
		fuel_requirements = append(fuel_requirements, int(fuel_required))
	}

	return fuel_requirements
}

func one_one() int {
	fuel_requirements := Calculate_fuel_requirements()

	result := 0
	for _, req := range fuel_requirements {
		result += req
	}
	log.Printf("one_one result: %d\n", result)

	return result
}
	

func calculate_fuel_requirement(mass float64, cache map[float64]float64) int {
		result := (float64(mass) / 3) - 2
		
		if result <= 0 {
			return 0
		} else if _, ok := cache[result]; ok { 
			result = result + cache[result]
		} else {
			sub_result := float64(calculate_fuel_requirement(result, cache))
			cache[result] = sub_result
			result = result + sub_result
		}
	
		return int(result)
	}
	
func calculate_fuel_fuel_requiements() []int {
	fuel_requirements := Calculate_fuel_requirements()

	results := make([]int, len(fuel_requirements))
	cache := make(map[float64]float64)

	for _, req := range fuel_requirements {
		fuel_required_for_fuel := calculate_fuel_requirement(float64(req), cache)
		results = append(results, req + fuel_required_for_fuel)
		log.Println(req, fuel_required_for_fuel)
	}

	return results
}

func one_two() int {
	fuel_requiements := calculate_fuel_fuel_requiements()

	result := 0
	for _, req := range fuel_requiements {
		result += req
	}

	log.Printf("one_two result: %d", result)
	return result
}

func main() {
	one_one()
	one_two()
}
	