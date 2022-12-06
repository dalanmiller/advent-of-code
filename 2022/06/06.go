package main

import (
	"bufio"
	"io"
)

type buffer []byte

func (b *buffer) append(n byte) {
	(*b) = append((*b)[1:len(*b)], n)
}

func (b *buffer) allDifferent() bool {
	m := make(map[byte]bool, len(*b))
	for i := range *b {
		if (*b)[i] == 0 {
			return false
		}
		m[(*b)[i]] = true
	}

	return len(m) == len(*b)
}

func run(input io.Reader, distinct int) int {

	s := bufio.NewScanner(input)
	s.Split(bufio.ScanRunes)
	buffer := buffer(
		make([]byte, distinct),
	)

	n := 1
	for s.Scan() {
		buffer.append(s.Bytes()[0])
		if buffer.allDifferent() {
			return n
		}
		n++
	}

	return -1
}
