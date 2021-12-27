package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {

	// 1
	outer := Node{Depth: 0}
	left := Node{Value: 1, Depth: 1, Parent: &outer, Branch: LEFT}
	right := Node{Value: 2, Depth: 1, Parent: &outer, Branch: RIGHT}
	outer.Left = &left
	outer.Right = &right

	// 2
	louter := Node{Depth: 1, Branch: LEFT}
	lleft := Node{Value: 1, Depth: 2, Parent: &louter, Branch: LEFT}
	lright := Node{Value: 2, Depth: 2, Parent: &louter, Branch: RIGHT}
	louter.Left = &lleft
	louter.Right = &lright

	router := Node{Depth: 1, Branch: RIGHT}
	rleft := Node{Value: 1, Depth: 2, Parent: &router, Branch: LEFT}
	rright := Node{Value: 2, Depth: 2, Parent: &router, Branch: RIGHT}
	router.Left = &rleft
	router.Right = &rright

	twoOuter := Node{
		Depth:  0,
		Left:   &louter,
		Right:  &router,
		Parent: nil,
	}
	louter.Parent = &twoOuter
	router.Parent = &twoOuter

	tests := []struct {
		test     string
		expected *Node
	}{
		// {"[1,2]", &outer},
		{"[[1,2],[1,2]]", &twoOuter},
	}

	for _, test := range tests {
		_, result := parseInput(test.test)
		if !reflect.DeepEqual(test.expected, result) {
			t.Fatalf("Result %v != expected %v", result, test.expected)
		}
	}

}

func TestString(t *testing.T) {
	louter := Node{Depth: 1}
	lleft := Node{Value: 1, Depth: 2, Parent: &louter}
	lright := Node{Value: 2, Depth: 2, Parent: &louter}
	louter.Left = &lleft
	louter.Right = &lright

	router := Node{Depth: 1}
	rleft := Node{Value: 1, Depth: 2, Parent: &router}
	rright := Node{Value: 2, Depth: 2, Parent: &router}
	router.Left = &rleft
	router.Right = &rright

	twoOuter := Node{
		Depth:  0,
		Left:   &louter,
		Right:  &router,
		Parent: nil,
	}
	louter.Parent = &twoOuter
	router.Parent = &twoOuter

	if twoOuter.String() != "[[1,2],[1,2]]" {
		t.Fatalf("String generation is FUBAR, expected '[[1,2],[1,2]]', got %s", twoOuter.String())
	}
}

func TestExplode(t *testing.T) {

	lleft := &Node{
		Value: 3,
	}

	lright := &Node{
		Value: 2,
	}

	left := &Node{
		Left:  lleft,
		Right: lright,
	}
	lleft.Parent = left
	lright.Parent = left

	mid := &Node{
		Left: left,
		Right: &Node{
			Value: 2,
		},
	}
	left.Parent = mid

	root := &Node{
		Left: &Node{
			Value: 1,
		},
		Right: mid,
	}
	mid.Parent = root

	// [1, [[3, 2], 2]]

	l := []*Node{}
	inOrderTraversal(root, &l, 0)

	root.Right.Left.Explode(&l)

	if root.Right.Left == nil || root.Right.Left.Value != 0 {
		t.Fatalf("Did not explode properly, Left reference is nil or not Value: 0")
	}

	if root.Right.Right.Value != 4 {
		t.Fatalf("Did not explode and add right properly, expected 4, got %d", root.Right.Right.Value)
	}

	if root.Left.Value != 4 {
		t.Fatalf("did not explode and add left properly, expected 4, got %d", root.Left.Value)
	}

	input := "[[[[[9,8],1],2],3],4]"
	l, root = parseInput(input)

	root.Left.Left.Left.Left.Explode(&l)

	if root.String() != "[[[[0,9],2],3],4]" {
		t.Fatalf("Explosion failure, expected %s, got %s", "[[[[0,9],2],3],4]", root.String())
	}

}

func TestExplodes(t *testing.T) {
	tests := []struct {
		test     string
		expected string
	}{
		{`[[[[1,2],1],1],1]
[1,1]`,
			`[[[[0,3],1],1],[1,1]]`},
		{`[1,[1,[1,[2,1]]]]
[1,1]`,
			`[[1,[1,[3,0]]],[2,1]]`},
		{`[1,1]
[[[[1,2],1],1],1]`,
			`[[1,2],[[[0,3],1],1]]`},
	}

	for _, test := range tests {
		result, _ := run(test.test)

		if result != test.expected {
			t.Fatalf("Explode result %s, did not match expected %s", result, test.expected)
		}
	}
}

func TestSplits(t *testing.T) {
	tests := []struct {
		test     string
		expected string
	}{
		{`[0,10]
[1,1]`,
			`[[0,[5,5]],[1,1]]`},
		{`[0,10]
[1,1]
[2,2]`,
			`[[[0,[5,5]],[1,1]],[2,2]]`},
	}

	for _, test := range tests {
		result, _ := run(test.test)

		if result != test.expected {
			t.Fatalf("Split result %s, did not match expected %s", result, test.expected)
		}
	}
}

func TestExamplesEighteenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected string
	}{
		{`[[[[4,3],4],4],[7,[[8,4],9]]]
		[1,1]`,
			`[[[[0,7],4],[[7,8],[6,0]]],[8,1]]`},
		{`[1,1]
		[2,2]
		[3,3]
		[4,4]`, "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`, "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
				{`[1,1]
		[2,2]
		[3,3]
		[4,4]
		[5,5]
		[6,6]`, "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
{`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]`, "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
{`[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]`, "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"},
			
				{`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
		[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
		[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
		[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
		[7,[5,[[3,8],[1,4]]]]
		[[2,[2,2]],[8,[8,1]]]
		[2,9]
		[1,[[[9,3],9],[[9,0],[0,7]]]]
		[[[5,[7,4]],7],1]
		[[[[4,2],2],6],[8,7]]`, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
				{`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
		[[[5,[2,8]],4],[5,[[9,9],0]]]
		[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
		[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
		[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
		[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
		[[[[5,4],[7,7]],8],[[8,3],8]]
		[[9,3],[[9,9],[6,[4,9]]]]
		[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
		[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result %s != expected %s", result, test.expected)
		}
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
		{`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`, 4140},
	}

	for _, test := range tests {
		_, mag := run(test.test)
		if mag != test.expected {
			t.Fatalf("Magnitude not correct, got %d, expected, %d", mag, test.expected)
		}
	}
}

func TestEighteenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 2501},
	}

	for _, test := range tests {
		_, mag := run(test.test)
		if mag != test.expected {
			t.Fatalf("Result %d != expected %d", mag, test.expected)
		}
	}
}

func TestExamplesEighteenTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
		[[[5,[2,8]],4],[5,[[9,9],0]]]
		[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
		[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
		[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
		[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
		[[[[5,4],[7,7]],8],[[8,3],8]]
		[[9,3],[[9,9],[6,[4,9]]]]
		[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
		[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`, 3993},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if test.expected != result {
			t.Fatalf("Result %d != expected %d", result, test.expected)
		}
	}
}

func TestEighteenTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 4935},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if result != test.expected {
			t.Fatalf("Result %d != expected %d", result, test.expected)
		}
	}
}
