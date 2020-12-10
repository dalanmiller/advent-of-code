package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "../helpers"
)

func makeEmptyPassport(fields []string) map[string]string {
	passport := map[string]string{}
	for _, field := range fields {
		passport[field] = ""
	}
	return passport
}

func part1() {
	rows, _ := ReadInput("./input")

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validPassports := 0
	currentPassport := makeEmptyPassport(fields)

	for _, row := range rows {
		if strings.TrimSpace(row) == "" {
			valid := true
			for _, value := range currentPassport {
				// fmt.Println(value)
				if value == "" {
					valid = false
					break
				}
			}
			if valid {
				validPassports++
			}

			currentPassport = makeEmptyPassport(fields)
		}

		cleanRow := strings.TrimSpace(row)
		for _, item := range strings.Split(cleanRow, " ") {
			fieldItems := strings.Split(item, ":")
			fieldName := fieldItems[0]

			if _, ok := currentPassport[fieldName]; ok {
				currentPassport[fieldName] = fieldItems[1]
			}
		}
	}
	fmt.Println("Part 1 | valid passports: ", validPassports)

}

func validateBYR(byrValue string) bool {
	val, _ := strconv.Atoi(byrValue)
	return len(byrValue) == 4 && 1920 <= val && val <= 2002
}

func validateIYR(iyrValue string) bool {
	val, _ := strconv.Atoi(iyrValue)
	return len(iyrValue) == 4 && 2010 <= val && val <= 2020
}

func validateEYR(eyrValue string) bool {
	val, _ := strconv.Atoi(eyrValue)
	return len(eyrValue) == 4 && 2020 <= val && val <= 2030
}

func validateHGT(hgtValue string) bool {
	if len(hgtValue) < 3 {
		return false
	}
	ending := hgtValue[len(hgtValue)-2:]
	re := regexp.MustCompile(`^(\d+)`)

	number, _ := strconv.Atoi(string(re.Find([]byte(hgtValue))))
	if ending == "cm" {
		return 150 <= number && number <= 193
	}
	return 59 <= number && number <= 76
}

func validateHCL(hclValue string) bool {
	re := regexp.MustCompile(`#[0-9a-fA-F]{6}`)
	return re.Match([]byte(hclValue))
}

func validateECL(eclValue string) bool {
	possibleEyeColors := map[string]int{"amb": 0, "blu": 0, "brn": 0, "gry": 0, "grn": 0, "hzl": 0, "oth": 0}
	if _, ok := possibleEyeColors[eclValue]; ok {
		return true
	} else {
		return false
	}
}

func validatePID(pidValue string) bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	return re.Match([]byte(pidValue))
}

func part2() {
	rows, _ := ReadInput("./input")

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validPassports := 0
	currentPassport := makeEmptyPassport(fields)

	validationMap := map[string]func(string) bool{
		"byr": validateBYR,
		"iyr": validateIYR,
		"eyr": validateEYR,
		"hgt": validateHGT,
		"hcl": validateHCL,
		"ecl": validateECL,
		"pid": validatePID,
	}

	for _, row := range rows {
		if strings.TrimSpace(row) == "" {
			valid := true
			for key, value := range currentPassport {
				if value == "" || !validationMap[key](value) {
					valid = false
					break
				}
			}
			if valid {
				validPassports++
			}

			currentPassport = makeEmptyPassport(fields)
		}

		cleanRow := strings.TrimSpace(row)
		for _, item := range strings.Split(cleanRow, " ") {
			fieldItems := strings.Split(item, ":")
			fieldName := fieldItems[0]

			if fieldName == "cid" {
				continue
			}

			if _, ok := currentPassport[fieldName]; ok {
				currentPassport[fieldName] = fieldItems[1]
			}
		}
	}
	fmt.Println("Part 2 | valid passports: ", validPassports)
}

func main() {
	part1()
	part2()
}
