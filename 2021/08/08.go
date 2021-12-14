package main

import (
	"sort"
	"strconv"
	"strings"
)

type Signal struct {
	Segments int
	Raw      string
	N        string
}

type Row struct {
	Signals []Signal
	Outputs []Signal
}

// Segments => Numbers
// 0 => 0
// 1 => 0
// 2 => 1
// 3 => 7
// 4 => 4
// 5 => 5, 3, 2
// 6 => 6, 0, 9
// 7 => 8

// For each wire "A, B, C, ...", is incorrectly lighting up SEGMENT A, B, C.... etc.
// First layer: deduce based on digits with unique segment counts 1, 7, 4, 8
// Second layer: ???

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func parseSignalsString(signal string) Signal {
	runeSlice := []rune(signal)
	sort.Sort(sortRuneString(runeSlice))
	sortedCleaned := strings.TrimLeft(string(runeSlice), "\t")

	return Signal{
		Segments: len(sortedCleaned),
		Raw:      sortedCleaned,
	}
}

func parseInput(input string) []Row {
	lines := strings.Split(input, "\n")

	var rows []Row
	for _, line := range lines {

		var signals, outputs []Signal
		signals_output := strings.Split(line, " | ")
		raw_signals := signals_output[0]
		raw_outputs := signals_output[1]

		signals_slice := strings.Split(raw_signals, " ")
		outputs_slice := strings.Split(raw_outputs, " ")

		for _, signal := range signals_slice {
			s := parseSignalsString(signal)
			signals = append(signals, s)
		}

		for _, output := range outputs_slice {
			o := parseSignalsString(output)
			outputs = append(outputs, o)
		}

		rows = append(rows, Row{
			Signals: signals,
			Outputs: outputs,
		})
	}

	return rows
}

func intersect(a Signal, b Signal) ([]rune, []rune) {
	seen_a := make(map[rune]bool)
	for _, chr := range a.Raw {
		seen_a[chr] = true
	}

	var intersection, difference []rune
	for _, chr := range b.Raw {
		_, ok := seen_a[chr]

		if ok {
			intersection = append(intersection, chr)
		}

		delete(seen_a, chr)
	}

	keySlice := make([]rune, 0, len(seen_a))
	for k, _ := range seen_a {
		keySlice = append(keySlice, k)
	}
	sort.Slice(keySlice, func(a, b int) bool {
		return int(keySlice[a]) < int(keySlice[b])
	})

	difference = append(difference, keySlice...)

	return intersection, difference
}

func run(input string) int {
	rows := parseInput(input)

	var sum int
	for _, row := range rows {
		for _, output := range row.Outputs {
			lo := output.Segments
			if lo == 2 || lo == 3 || lo == 4 || lo == 7 {
				sum++
			}
		}
	}

	return sum
}

func run_two(input string) int {
	rows := parseInput(input)

	var sum int
	for _, row := range rows {
		// log.Printf("Row - %d", i)
		signals := row.Signals
		outputs := row.Outputs

		var Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine Signal
		var unknown []Signal
		for _, signal := range signals {
			switch signal.Segments {
			case 2:
				One = signal
				One.N = "1"
			case 3:
				Seven = signal
				Seven.N = "7"
			case 4:
				Four = signal
				Four.N = "4"
			case 7:
				Eight = signal
				Eight.N = "8"
			default:
				unknown = append(unknown, signal)
			}
		}

		for _, unk := range unknown {
			if unk.Segments == 5 {
				// 5, 3, or 2 henceforth
				_, d := intersect(unk, One)
				if len(d) == 4 {
					// 2 or 5
					_, d := intersect(unk, Four)
					if len(d) == 3 {
						//2
						Two = unk
						Two.N = "2"
					} else { // 2 difference
						//5
						Five = unk
						Five.N = "5"
					}
				} else {
					Three = unk
					Three.N = "3"
				}

			} else {
				// 0, 6, or 9 henceforth
				_, d := intersect(unk, Seven)
				if len(d) == 3 {
					// 0 or 9 henceforth
					_, d := intersect(unk, Four)
					if len(d) == 2 {
						// 9
						Nine = unk
						Nine.N = "9"
					} else {
						// 5
						Zero = unk
						Zero.N = "0"
					}

				} else {
					// 6
					Six = unk
					Six.N = "6"
				}
			}
		}

		// Need to deduce final Signal based on one not found

		signalMap := make(map[string]Signal)
		for _, signal := range []Signal{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine} {
			signalMap[signal.Raw] = signal
		}

		final_output := make([]string, len(outputs))
		for i, output := range outputs {
			final_output[i] = signalMap[output.Raw].N
		}

		value, _ := strconv.Atoi(strings.Join(final_output, ""))
		sum += value
	}

	return sum
}
