package main

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	n     int
	items []int
	mod   int
	op    func(int) int
	test  func(int) int
}

func (m *monkey) runOpPartOne(i int) {
	m.items[0] = m.op(i) / 3
}

func (m *monkey) runOpPartTwo(i int, mod int) {
	m.items[0] = m.op(i) % int(mod)
}

func (m *monkey) runTest(i int) int {
	return m.test(i)
}

func (m *monkey) throw(om *monkey) {

	// Pop off front of current monkey
	item := m.items[0]
	m.items = m.items[1:]

	// Append to other monkey
	om.items = append(om.items, item)
}

// Monkey 0:
//   Starting items: 79, 98
//   Operation: new = old * 19
//   Test: divisible by 23
//     If true: throw to monkey 2
//     If false: throw to monkey 3

var oldPlus = func(i int) int {
	return i + i
}

var oldMult = func(i int) int {
	return i * i
}

func readInput(input io.Reader) []*monkey {
	monkeys := []*monkey{}
	s := bufio.NewScanner(input)

	// Scans for next "monkey chunk"
	for s.Scan() {
		// Line one
		lineOne := s.Text()
		if lineOne == "" {
			continue
		}
		n, _ := strconv.Atoi(string(strings.Split(lineOne, " ")[1][0]))

		// Line two
		s.Scan()
		lineTwo := s.Text()
		ltSplit := strings.Split(lineTwo, " ")
		items := []int{}
		for _, sp := range ltSplit[4:] {
			itemValue, _ := strconv.ParseInt(strings.TrimRight(sp, ","), 10, 64)
			items = append(items, int(itemValue))
		}

		// Line three
		s.Scan()
		lineThree := s.Text()
		split := strings.Split(lineThree, " ")[6:]
		var op func(int) int
		if split[0] == "+" {
			if split[1] == "old" {
				op = oldPlus
			} else {
				v, _ := strconv.Atoi(split[1])
				op = func(i int) int {
					return i + int(v)
				}
			}
		} else if split[0] == "*" {
			if split[1] == "old" {
				op = oldMult
			} else {
				v, _ := strconv.Atoi(split[1])
				op = func(i int) int {
					return i * int(v)
				}
			}
		}

		// Line 4, 5, 6
		s.Scan()
		lineFour := s.Text()
		s.Scan()
		lineFive := s.Text()
		s.Scan()
		lineSix := s.Text()

		lfSplit := strings.Split(lineFive, " ")
		trueMonkey, _ := strconv.Atoi(lfSplit[len(lfSplit)-1])
		lsSplit := strings.Split(lineSix, " ")
		falseMonkey, _ := strconv.Atoi(lsSplit[len(lsSplit)-1])

		tN, _ := strconv.Atoi(lineFour[21:])
		test := func(i int) int {
			if i%int(tN) == 0 {
				return trueMonkey
			} else {
				return falseMonkey
			}
		}

		monkeys = append(monkeys, &monkey{
			n:     n,
			items: items,
			mod:   tN,
			op:    op,
			test:  test,
		})
	}

	return monkeys
}

func run(input io.Reader, rounds int, part int) int {
	monkeys := readInput(input)
	monkeyCount := map[int]int{}

	mod := 1
	for _, monkey := range monkeys {
		mod *= monkey.mod
	}

	for i := 1; i <= rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeyCount[monkey.n]++

				if part == 1 {
					monkey.runOpPartOne(item)
				} else {
					monkey.runOpPartTwo(item, mod)
				}

				targetMonkey := monkey.runTest(monkey.items[0])

				monkey.throw(monkeys[targetMonkey])
			}
		}
		// if i%1000 == 0 || i == 1 || i == 20 {
		// 	log.Printf("%d | %v\n", i, monkeyCount)
		// }
	}

	counts := []int{}
	for _, v := range monkeyCount {
		counts = append(counts, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	return int(counts[0]) * int(counts[1])
}
