package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	q     int
	start int
	end   int
}

type crane struct {
	stacks       [][]string
	instructions []instruction
}

func reverseStack(stack []string) []string {
	new := make([]string, len(stack))

	k := 0
	for j := len(stack) - 1; j >= 0; j-- {
		new[k] = stack[j]
		k++
	}

	return new
}

func (c *crane) move(i instruction, craneVersion int) {
	//`   [D]
	//[N] [C]
	//[Z] [M] [P]`

	delta := len(c.stacks[i.start]) - i.q
	movingStack := c.stacks[i.start][delta:len(c.stacks[i.start])]

	c.stacks[i.start] = c.stacks[i.start][0:delta]
	switch craneVersion {
	case 9000:
		c.stacks[i.end] = append(c.stacks[i.end], reverseStack(movingStack)...)
	case 9001:
		c.stacks[i.end] = append(c.stacks[i.end], movingStack...)
	}
}

var crateRegex = regexp.MustCompile(`\[(\w)\]`)
var instructionRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func readInput(input io.Reader) crane {
	scanner := bufio.NewScanner(input)

	s := crane{
		instructions: []instruction{},
		stacks:       nil,
	}

	columns := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Init stacks because we haven't seen a line until now
		if columns == 0 {
			columns = (len(line) + 1) / 4
			s.stacks = make([][]string, columns)
			for i := range s.stacks {
				s.stacks[i] = []string{}
			}
		}

		// Only care about crate lines and
		//  instruction lines
		if strings.Contains(line, "[") {
			matches := crateRegex.FindAllIndex([]byte(line), -1)

			for _, match := range matches {
				index := (match[0] + 1) / 4
				s.stacks[index] = append(s.stacks[index], line[match[0]+1:match[0]+2])
			}

		} else if len(line) > 0 && line[0] == 'm' {

			match := instructionRegex.FindStringSubmatch(line)

			q, _ := strconv.Atoi(match[1])
			start, _ := strconv.Atoi(match[2])
			end, _ := strconv.Atoi(match[3])

			s.instructions = append(s.instructions, instruction{
				q,
				start - 1, // 1 index to 0 index
				end - 1,   // 1 index to 0 index
			})

		}
	}
	for i := range s.stacks {
		s.stacks[i] = reverseStack(s.stacks[i])
	}

	return s
}

func run(input io.Reader, craneVersion int) string {
	crane := readInput(input)

	for _, inst := range crane.instructions {
		crane.move(inst, craneVersion)
	}

	b := strings.Builder{}
	for i := range crane.stacks {
		b.WriteString(crane.stacks[i][len(crane.stacks[i])-1])
	}

	return b.String()
}
