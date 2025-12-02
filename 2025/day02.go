package year2025

import (
	"io"
	"log"
	"strconv"
	"strings"
)

func Day02PartOne(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	line := strings.TrimSpace(string(b))
	ranges := strings.Split(line, ",")
	// log.Printf("Ranges: %s", ranges)

	sum := 0

	for _, r := range ranges {
		left_right := strings.Split(r, "-")
		// log.Println(left_right)
		start, _ := strconv.Atoi(left_right[0])
		end, _ := strconv.Atoi(left_right[1])
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 == 0 {
				mid := len(str) / 2
				left := str[:mid]
				right := str[mid:]

				if left == right {

					// log.Printf("%s <> %s : %v", left, right, left == right)
					sum += i
				}
			}
		}
	}

	return sum
}

func splitEvery(s string, n int) []string {
	if n <= 0 {
		return nil
	}
	var out []string
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) { // Not desirable but oh well
			end = len(s)
		}
		out = append(out, s[i:end])
	}
	// log.Printf("%s", out)
	return out
}

func allEqual(ss []string) bool {
	if len(ss) == 0 {
		return true
	}

	first := ss[0]
	for _, s := range ss[1:] {
		if s != first {
			return false
		}
	}
	return true
}

func Day02PartTwo(r io.Reader) int {

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	line := strings.TrimSpace(string(b))
	ranges := strings.Split(line, ",")
	log.Printf("Ranges: %s", ranges)

	sum := 0

	for _, r := range ranges {
		left_right := strings.Split(r, "-")
		// log.Println(left_right)
		start, _ := strconv.Atoi(left_right[0])
		end, _ := strconv.Atoi(left_right[1])
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			for j := 1; j <= (len(str) / 2); j++ {
				if allEqual(splitEvery(str, j)) {
					log.Printf("Winner!: %d", i)
					sum += i
					break
				}
			}

		}
	}
	return sum
}
