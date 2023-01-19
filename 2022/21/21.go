package main

import (
	"io"
	"strconv"
	"strings"
)

type operation string

const (
	NULL  operation = ""
	PLUS  operation = "+"
	MINUS operation = "-"
	MULT  operation = "*"
	DIV   operation = "/"
)

var inverseMap = map[operation]operation{
	PLUS:  MINUS,
	MINUS: PLUS,
	MULT:  DIV,
	DIV:   MULT,
}

type opFunction func(a, b int) int

// Typical op map

var opMap = map[operation]opFunction{
	PLUS:  func(a, b int) int { return a + b },
	MINUS: func(a, b int) int { return a - b },
	MULT:  func(a, b int) int { return a * b },
	DIV:   func(a, b int) int { return a / b },
}

// Not sure if there is a more elegant way to
// . encode differences when operands are on either side
// . of the operator.

var opInverseLeftMap = map[operation]opFunction{
	PLUS:  func(a, b int) int { return a - b },
	MINUS: func(a, b int) int { return (a - b) / -1 },
	MULT:  func(a, b int) int { return a / b },
	DIV:   func(a, b int) int { return b / a },
}

var opInverseRightMap = map[operation]opFunction{
	PLUS:  func(a, b int) int { return a - b },
	MINUS: func(a, b int) int { return a + b },
	MULT:  func(a, b int) int { return a / b },
	DIV:   func(a, b int) int { return a * b },
}

// Side the known value is on
type side bool

const (
	LEFT  side = false
	RIGHT side = true
)

type monkey struct {
	Name          string
	Operator      operation
	Value         int
	RawDependents [2]string
	Dependents    [2]*monkey
}

type stage struct {
	Name  string
	Op    operation
	Value int
	Side  side
}

func readInput(input io.Reader) map[string]*monkey {
	bytes, _ := io.ReadAll(input)
	lines := string(bytes)

	monkeys := make([]*monkey, 0, len(lines))
	monkeyMap := make(map[string]*monkey, len(lines))

	for _, line := range strings.Split(lines, "\n") {
		split := strings.Split(string(line), " ")
		m := monkey{
			Name: split[0][:4],
		}
		monkeys = append(monkeys, &m)
		monkeyMap[split[0][:4]] = &m
		switch len(split) {
		case 2:
			n, _ := strconv.Atoi(split[1])
			m.Value = n
		case 4:
			m.RawDependents[0] = split[1]
			m.RawDependents[1] = split[3]
			m.Operator = operation(split[2])
		}
	}

	for _, m := range monkeys {
		if m.Operator == NULL {
			continue
		}

		m.Dependents[0] = monkeyMap[m.RawDependents[0]]
		m.Dependents[1] = monkeyMap[m.RawDependents[1]]
	}

	return monkeyMap
}

func rollUp(m *monkey) {
	if m.Value != 0 {
		// log.Printf("%s yells %d", m.Name, m.Value)
		return
	}

	leftM := m.Dependents[0]
	rightM := m.Dependents[1]

	if leftM.Value == 0 {
		rollUp(leftM)
	}

	if rightM.Value == 0 {
		rollUp(rightM)
	}

	m.Value = opMap[m.Operator](leftM.Value, rightM.Value)

	// log.Printf("%s yells %s %s %s = %d", m.Name, leftM.Name, m.Operator, rightM.Name, m.Value)
}

func rollUpOperations(m *monkey) []stage {
	if m.Name == "humn" {
		return []stage{}
	}
	var left, right []stage
	if m.Dependents[0] != nil {
		left = rollUpOperations(m.Dependents[0])
	}

	if m.Dependents[1] != nil {
		right = rollUpOperations(m.Dependents[1])
	}

	if left != nil {
		return append(left, stage{
			Op:    m.Operator,
			Value: m.Dependents[1].Value,
			Name:  m.Name,
			Side:  RIGHT,
		})
	} else if right != nil {
		return append(right, stage{
			Op:    m.Operator,
			Value: m.Dependents[0].Value,
			Name:  m.Name,
			Side:  LEFT,
		})
	}

	return nil
}

func run(input io.Reader) int {
	mMap := readInput(input)
	rollUp(mMap["root"])
	return mMap["root"].Value
}

func runTwo(input io.Reader) int {
	mMap := readInput(input)

	leftMonkey := mMap["root"].Dependents[0]
	rightMonkey := mMap["root"].Dependents[1]

	rollUp(leftMonkey)
	rollUp(rightMonkey)

	leftStages := rollUpOperations(leftMonkey)
	rightStages := rollUpOperations(rightMonkey)

	var stages []stage
	var target int

	if leftStages != nil {
		stages = leftStages
		target = rightMonkey.Value
	} else {
		stages = rightStages
		target = leftMonkey.Value
	}

	// Go from outside in || last to first, on the stages
	for i := len(stages) - 1; i >= 0; i-- {
		stage := stages[i]

		switch stage.Side {
		case LEFT:
			target = opInverseLeftMap[stage.Op](target, stage.Value)
		case RIGHT:
			target = opInverseRightMap[stage.Op](target, stage.Value)
		}
	}

	return target
}
