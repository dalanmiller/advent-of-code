package main

import (
	"io"
	"strconv"
	"strings"

	"container/ring"
)

// [ A, B, C ]
// Index {A: 0, B: 1, C: 2}
// 1. A moves 1
// [ B, A, C ]
// Index {A: 1, B: 0, C: 2}
// 2. B moves 2
// [ A, C, B ]
// Index {A: 0, B: 2, C: 1}
// 3. C moves -1
// [ C, A, B ]
// Index {A: 1, B: 2, C: 0}

// [ 2, 3, 2, 1]

type value int

const (
	CURRENT_INDEX value = 0
	NUMBER        value = 1
)

type k struct {
	index int
	value int
}

func readInput(input io.Reader, dk int) ([]int, map[k]*ring.Ring) {
	bytes, _ := io.ReadAll(input)
	lines := strings.Split(string(bytes), "\n")

	positions := make(map[k]*ring.Ring, len(lines))
	ogList := make([]int, len(lines))
	r := ring.New(len(lines))

	for i, n := range lines {
		value, _ := strconv.Atoi(n)
		value *= dk

		ogList[i] = value
		positions[k{i, value}] = r
		r.Value = value
		r = r.Next()
	}

	return ogList, positions
}

func run(input io.Reader) int {
	ogList, ringMap := readInput(input, 1)

	var zeroKey *ring.Ring
	for index, number := range ogList {

		// We have to stpe back one as unlinking
		// . will create a new Ring starting with the
		// . the first subsequent link in the ring.
		r := ringMap[k{index, number}].Prev()

		toMove := r.Unlink(1)

		if toMove.Value == 0 {
			zeroKey = toMove
		}

		r.Move(number).Link(toMove)
	}

	s := 0
	for _, n := range []int{1000, 2000, 3000} {
		s += zeroKey.Move(n).Value.(int)
	}

	return s
}

func runPartTwo(input io.Reader, dk int) int {
	ogList, ringMap := readInput(input, dk)

	length := len(ogList) - 1
	halflen := length >> 1

	var zeroKey *ring.Ring
	for i := 1; i <= 10; i++ {
		for index, number := range ogList {

			// We have to stpe back one as unlinking
			// . will create a new Ring starting with the
			// . the first subsequent link in the ring.
			r := ringMap[k{index, number}].Prev()

			toMove := r.Unlink(1)

			if (number > halflen) || (number < -halflen) {
				number %= length
				switch {
				case number > halflen:
					number -= length
				case number < -halflen:
					number += length
				}
			}

			if toMove.Value == 0 {
				zeroKey = toMove
			}

			r.Move(number).Link(toMove)
		}
	}

	s := 0
	for _, n := range []int{1000, 2000, 3000} {
		s += zeroKey.Move(n).Value.(int)
	}

	return s
}
