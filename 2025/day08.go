package year2025

import (
	"bufio"
	"cmp"
	"container/heap"
	"io"
	"slices"
	"strconv"
	"strings"
)

type minDistanceHeap [][3]int // [distance, point i, point j]

func (h minDistanceHeap) Len() int           { return len(h) }
func (h minDistanceHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h minDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minDistanceHeap) Push(x any) {
	*h = append(*h, x.([3]int))
}

func (h *minDistanceHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func manhattanDistance(i [3]int, j [3]int) int {
	return toTheSecond(i[0]-j[0]) + toTheSecond(i[1]-j[1]) + toTheSecond(i[2]-j[2])
}

func toTheSecond(x int) int {
	return x * x
}

func recurse(node int, circuits map[int][]int, seen map[int]struct{}) int {
	if _, ok := seen[node]; ok {
		return 0
	}
	seen[node] = struct{}{}
	size := 1
	for _, v := range circuits[node] {
		size += recurse(v, circuits, seen)
	}
	return size
}

func Day08PartOne(r io.Reader, connections int) int {
	s := bufio.NewScanner(r)

	coords := [][3]int{}
	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		coords = append(coords, [3]int{x, y, z})
	}

	mHeap := &minDistanceHeap{}
	heap.Init(mHeap)
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			distance := manhattanDistance(coords[i], coords[j])
			heap.Push(mHeap, [3]int{distance, i, j})
		}
	}

	circuits := make(map[int][]int, len(*mHeap))
	for i := range coords {
		// m := make(map[int]struct{})
		circuits[i] = []int{}
	}

	// Populate a map of junction index => array of connected junctions
	for i := 1; i <= connections; i++ {
		x := heap.Pop(mHeap).([3]int)
		circuits[x[1]] = append(circuits[x[1]], x[2])
		circuits[x[2]] = append(circuits[x[2]], x[1])
	}

	// Classically, in graph problems need to keep track of seen nodes
	seen := make(map[int]struct{})
	circuitSizes := []int{}
	for k := range circuits {
		if _, ok := seen[k]; ok {
			continue
		}
		// Mark as seen before descending into recursive madness.
		circuitSizes = append(circuitSizes, recurse(k, circuits, seen))
	}

	slices.SortFunc(circuitSizes, func(a int, b int) int {
		return cmp.Compare(b, a)
	})

	product := 1
	for _, n := range circuitSizes[:3] {
		product *= n
	}
	return product
}

func Day08PartTwo(r io.Reader) int {
	s := bufio.NewScanner(r)

	coords := [][3]int{}
	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		coords = append(coords, [3]int{x, y, z})
	}

	mHeap := &minDistanceHeap{}
	heap.Init(mHeap)
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			distance := manhattanDistance(coords[i], coords[j])
			heap.Push(mHeap, [3]int{distance, i, j})
		}
	}

	circuits := make(map[int][]int, len(*mHeap))
	for i := range coords {
		// m := make(map[int]struct{})
		circuits[i] = []int{}
	}

	// x := heap.Pop(mHeap).([3]int)
	// return coords[x[1]][0] * coords[x[2]][0]

	// Populate a map of junction index => array of connected junctions
	var seen map[int]struct{}
	for mHeap.Len() > 0 {
		x := heap.Pop(mHeap).([3]int)
		circuits[x[1]] = append(circuits[x[1]], x[2])
		circuits[x[2]] = append(circuits[x[2]], x[1])

		seen = make(map[int]struct{})
		if recurse(x[1], circuits, seen) == len(coords) {
			return coords[x[1]][0] * coords[x[2]][0]
		}
	}

	// Classically, in graph problems need to keep track of seen nodes
	//
	// circuitSizes := []int{}
	// for k := range circuits {
	// 	if _, ok := seen[k]; ok {
	// 		continue
	// 	}
	// 	// Mark as seen before descending into recursive madness.
	// 	circuitSizes = append(circuitSizes, recurse(k, circuits, seen))
	// }

	// slices.SortFunc(circuitSizes, func(a int, b int) int {
	// 	return cmp.Compare(b, a)
	// })

	// product := 1
	// for _, n := range circuitSizes[:3] {
	// 	product *= n
	// }
	// return product
	return -1
}
