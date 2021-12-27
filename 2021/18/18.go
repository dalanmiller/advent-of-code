package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	Left   *Node
	Right  *Node
	Value  int
	Depth  int
	Parent *Node
	Branch Branch
}

type Stack struct {
	Stack []*Node
}

func (s *Stack) Push(n *Node) {
	s.Stack = append(s.Stack, n)
}

func (s *Stack) Pop() *Node {
	last := len(s.Stack) - 1
	r := s.Stack[last]
	s.Stack = s.Stack[:last]
	return r
}

func (s *Stack) HasNext() bool {
	return len(s.Stack) > 0
}

func (n Node) String() string {
	if n.IsLeaf() {
		return fmt.Sprintf("%d", n.Value)
	}
	return fmt.Sprintf("[%s,%s]", n.Left.String(), n.Right.String())
}

func (n *Node) Reduce(l []*Node) bool {
	s := Stack{Stack: []*Node{}}
	s.Push(n)
	cur := n
	exploded := false 
	for s.HasNext() {

		// Conditions for a node that represents a 'Number' and ensure that both Left and Right
		// are present and leaf nodes.
		if cur.Depth == 4 && cur.Left != nil && cur.Left.IsLeaf() && cur.Right != nil && cur.Right.IsLeaf() {
			cur.Explode(&l)
			l = nil 
			inOrderTraversal(n, &l, 0)
			exploded = true 
			// No need to return here, cool to keep exploding
		}

		// Need to ensure that if an explosion has happened that a split can happen after that in this 'cycle'
		if !exploded && cur.Value >= 10 {
			cur.Split()
			l = nil 
			inOrderTraversal(n, &l, 0)

			// Immediately stop and re-reduce if we split because of precedence
			return true
		}

		// Thus boys and girls, we traverse
		if cur.Left != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			cur = s.Pop().Right
		}
	}

	return false || exploded
}

func (n Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) Explode(l *[]*Node) {
	// A list of in order nodes in the tree is the least
	// headachey way to do this. Find the node, find the one to it's left,
	// That's the one to increment by the values you have in your current pair.
	// This is far simpler and less complex than reverse and non-reversed in order traversal
	// originating from the node that is exploding for fuck's sake.

	for i, node := range *l {
		if node == n.Left {
			if i-1 >= 0 {
				(*l)[i-1].Value += n.Left.Value
			}

			if i+2 < len(*l) {
				(*l)[i+2].Value += n.Right.Value
			}
			break
		}
	}

	// Set pointer of parent of this node
	// to a new node which is just value 0
	// Set to zero / nil
	n.Left = nil
	n.Right = nil
	n.Value = 0
	n.Depth = 0
}

func (n *Node) Split() {
	v := float64(n.Value) / 2.0
	left := math.Floor(v)
	right := math.Ceil(v)

	// Node no longer has value
	n.Value = 0

	// Create new nodes on left and on right with values
	n.Left = &Node{
		Branch: LEFT,
		Depth:  n.Depth + 1,
		Value:  int(left),
		Parent: n,
	}
	n.Right = &Node{
		Branch: RIGHT,
		Depth:  n.Depth + 1,
		Value:  int(right),
		Parent: n,
	}
}

type Token rune

const (
	OPEN  Token = '['
	CLOSE Token = ']'
	COMMA Token = ','
)

type Branch rune

const (
	LEFT  Branch = 'L'
	RIGHT Branch = 'R'
)

func parseInput(input string) ([]*Node, *Node) {
	cursor := 0
	depth := 0
	root := Node{
		Depth:  depth,
		Parent: nil,
	}

	currentNode := &root
	nodeList := []*Node{}

	input = strings.TrimLeft(input, "\t ")
	for cursor < len(input) {
		switch Token(rune(input[cursor])) {
		case OPEN:
			depth++
			left := &Node{
				Depth:  depth,
				Parent: currentNode,
				Branch: LEFT,
			}

			currentNode.Left = left
			currentNode = left
		case CLOSE:
			depth--
			currentNode = currentNode.Parent

		case COMMA:
			right := &Node{
				Depth:  depth,
				Parent: currentNode.Parent,
				Branch: RIGHT,
			}

			currentNode.Parent.Right = right
			currentNode = right

		default:
			// Parse single to multidigit numbers
			end := cursor + 1
			for end < len(input) && unicode.IsDigit(rune(input[end])) {
				end++
			}
			v, _ := strconv.Atoi(string(input[cursor:end]))
			currentNode.Value = v
			nodeList = append(nodeList, currentNode)
			if end > cursor+1 {
				cursor = end - 1
			}
		}

		cursor++
	}

	return nodeList, &root
}

func inOrderTraversal(n *Node, l *[]*Node, depth int) {
	if n.IsLeaf() {
		*l = append(*l, n)
	}
	n.Depth = depth

	if n.Left != nil {
		inOrderTraversal(n.Left, l, depth+1)
	}

	if n.Right != nil {
		inOrderTraversal(n.Right, l, depth+1)
	}
}

func magnitude(r *Node) int {
	if r.IsLeaf() {
		return r.Value
	}

	return 3*magnitude(r.Left) + 2*magnitude(r.Right)
}

func run(input string) (string, int) {

	numbers := strings.Split(input, "\n")

	first := numbers[0]
	numbers = numbers[1:]

	leftNodes, leftRoot := parseInput(first)
	for i, number := range numbers {
		if i > 0 {
			leftNodes = make([]*Node, 0, len(leftNodes))
			inOrderTraversal(leftRoot, &leftNodes, 0)
		}
		rightNodes, rightRoot := parseInput(number)
		fullNodes := make([]*Node, 0, len(leftNodes)+len(rightNodes))

		root := &Node{
			Left:  leftRoot,
			Right: rightRoot,
			Depth: 0,
		}
		leftRoot.Parent = root
		leftRoot.Branch = LEFT
		rightRoot.Parent = root
		rightRoot.Branch = RIGHT
		inOrderTraversal(root, &fullNodes, 0)

		for root.Reduce(fullNodes) {
			fullNodes = nil

			// Re-traverse the tree and recreate the fullNodes
			// helper slice.
			inOrderTraversal(root, &fullNodes, 0)
		}
		leftRoot = root
	}

	return leftRoot.String(), magnitude(leftRoot)
}

func run_two(input string) int {
	numbers := strings.Split(input, "\n")
	
	max := 0 
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			_, mag := run(strings.Join([]string{n1,n2}, "\n"))
			if mag > max {
				max = mag 
			}
		}
	}

	return max
}
