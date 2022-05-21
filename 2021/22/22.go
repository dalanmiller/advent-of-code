package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Status bool

const (
	ON  Status = true
	OFF Status = false
)

// type Range struct {
// 	Start int
// 	End   int
// }

type Instruction struct {
	S Status
	C Cube
	// X Range
	// Y Range
	// Z Range
}

type Cube struct {
	Lower Coord
	Upper Coord
}

type CubeSet []Cube

type Coord [3]int

const (
	X = 0
	Y = 1
	Z = 2
)

func (c Cube) Volume(bounds float64) uint64 {
	volume := uint64(1)
	for d := X; d <= Z; d++ {
		volume *= uint64(c.Upper[d] - c.Lower[d] + 1)
	}

	return volume
}

func (c Cube) Overlap(oc Cube) *Cube {

	//     OnCube      OffCube
	// 1: [-------]   (-------) No overlap, other D greater
	// 2: [----(^^]-----------) Other overlaps on left side
	// 3: (---[^^^^^^^^^^^^^]-) Other encompasses cube
	// 4: [----(^^^^^^^^^^^^)-] Cube encompasses other
	// 5: (-------[^^^^)------] Other overlaps on right side
	// 6: (-------)   [-------] No overlap, other D lesser

	// https://github.com/lizthegrey/adventofcode/blob/main/2021/day22.go#L44-L50

	var nc Cube

	// Rejiggering this in March 2022
	// Thinking that I can just shrink the OnCube by the intersection it has
	//  with the provided OffCube

	// My version && not working
	// for d := X; d <= Z; d++ {
	// 	if c.Upper[d] > oc.Lower[d] || c.Lower[d] > oc.Upper[d] {
	// 		// Case 1 && Case 6
	// 		return nil
	// 	} else if c.Upper[d] >= oc.Lower[d] && c.Lower[d] <= oc.Upper[d] {
	// 		// Case 2
	// 		nc.Upper[d] = c.Upper[d]
	// 		nc.Lower[d] = oc.Lower[d]
	// 		// } else if c.Upper[d] <= oc.Upper[d] && c.Lower[d] >= oc.Lower[d] {
	// 	} else if oc.Upper[d] <= c.Upper[d] && oc.Lower[d] >= c.Lower[d] {
	// 		// Case 3 - OnCube is completely immersed in OffCube
	// 		nc.Upper[d] = c.Upper[d]
	// 		nc.Lower[d] = c.Lower[d]
	// 	} else if c.Upper[d] <= oc.Upper[d] && c.Lower[d] >= oc.Lower[d] {
	// 		// Case 4 - OffCube is immersed by OnCube ... wow
	// 		nc.Upper[d] = oc.Upper[d]
	// 		nc.Lower[d] = oc.Lower[d]
	// 	} else if c.Lower[d] <= oc.Lower[d] && c.Upper[d] <= oc.Upper[d] {
	// 		// Case 5
	// 		nc.Upper[d] = c.Lower[d]
	// 		nc.Lower[d] = oc.Upper[d]
	// 	} else {
	// 		panic("Shouldn't arrive here, overlap?")
	// 	}
	// }

	// Lizthegrey version and ~working
	for d := X; d <= Z; d++ {
		if c.Lower[d] > oc.Upper[d] {
			// No overlap, case 1
			return nil
		} else if c.Upper[d] < oc.Lower[d] {
			// No overlap, case 6
			return nil
		} else if c.Upper[d] <= oc.Upper[d] && c.Lower[d] >= oc.Lower[d] {
			// Full overlap, case 4
			nc.Upper[d] = c.Upper[d]
			nc.Lower[d] = c.Lower[d]
		} else if oc.Upper[d] <= c.Upper[d] && oc.Lower[d] >= c.Lower[d] {
			// Full overlap, case 3
			nc.Upper[d] = oc.Upper[d]
			nc.Lower[d] = oc.Lower[d]
		} else if oc.Lower[d] <= c.Lower[d] && oc.Upper[d] <= c.Upper[d] {
			// Partial overlap, case 2
			nc.Upper[d] = oc.Upper[d]
			nc.Lower[d] = c.Lower[d]
		} else if c.Lower[d] <= oc.Lower[d] && c.Upper[d] <= oc.Upper[d] {
			// Partial overlap, case 5
			nc.Upper[d] = c.Upper[d]
			nc.Lower[d] = oc.Lower[d]
		} else {
			// There is a bug.
			panic("Unaccounted overlap case.")
		}
	}

	return &nc
}

func (c Cube) Subtract(o Cube) (newCubes []Cube) {
	if c == o {
		// Full subtraction
		return newCubes
	}

	// We need to construct the prisms after cutting out the overlap.
	// There are 6 potential pieces we need to construct:
	// The slab below (ZBottom)
	// Where the Z coordinates overlap, we have (top-down, X and Y only):
	// [-----YUpper---------]
	// [Left] Overlap [Right]
	// [-----YLower---------]
	// The slab above (ZAbove)

	if o.Lower[Z] != c.Lower[Z] {
		// Construct the bottom slab.
		lower := c.Lower
		upper := c.Upper
		upper[Z] = o.Lower[Z] - 1
		newCubes = append(newCubes, Cube{lower, upper})
	}

	// Construct the Y-lower slab, bounded at top and bottom Z by overlap Zs
	if o.Lower[Y] != c.Lower[Y] {
		lower := c.Lower
		upper := c.Upper
		lower[Z] = o.Lower[Z]
		upper[Z] = o.Upper[Z]
		upper[Y] = o.Lower[Y] - 1
		newCubes = append(newCubes, Cube{lower, upper})
	}

	// Construct the left and right slabs.
	if o.Lower[X] != c.Lower[X] {
		lower := o.Lower
		upper := o.Upper
		lower[X] = c.Lower[X]
		upper[X] = o.Lower[X] - 1
		newCubes = append(newCubes, Cube{lower, upper})
	}

	if o.Upper[X] != c.Upper[X] {
		lower := o.Lower
		upper := o.Upper
		lower[X] = o.Upper[X] + 1
		upper[X] = c.Upper[X]
		newCubes = append(newCubes, Cube{lower, upper})
	}

	// Construct the Y-upper slab, bounded at top and bottom Z by overlap Zs
	if o.Upper[Y] != c.Upper[Y] {
		lower := c.Lower
		upper := c.Upper
		lower[Z] = o.Lower[Z]
		upper[Z] = o.Upper[Z]
		lower[Y] = o.Upper[Y] + 1
		newCubes = append(newCubes, Cube{lower, upper})
	}
	if o.Upper[Z] != c.Upper[Z] {
		// Construct the ZAbove slab.
		lower := c.Lower
		upper := c.Upper
		lower[Z] = o.Upper[Z] + 1
		newCubes = append(newCubes, Cube{lower, upper})
	}
	return newCubes
}

var instructionRegex = regexp.MustCompile(`(\w+) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)

func parseInput(input string, bounds int) ([]Instruction, []Cube, []Cube) {
	lines := strings.Split(input, "\n")

	instructions := make([]Instruction, 0, len(lines))
	onCubes := []Cube{}
	offCubes := []Cube{}
	for _, line := range lines {

		values := instructionRegex.FindStringSubmatch(line)

		var s Status
		if values[1] == "on" {
			s = ON
		} else {
			s = OFF
		}

		var xStart, xEnd, yStart, yEnd, zStart, zEnd int
		xStart, _ = strconv.Atoi(values[2])
		xEnd, _ = strconv.Atoi(values[3])
		yStart, _ = strconv.Atoi(values[4])
		yEnd, _ = strconv.Atoi(values[5])
		zStart, _ = strconv.Atoi(values[6])
		zEnd, _ = strconv.Atoi(values[7])

		skip := false
		for _, v := range []int{xStart, yStart, zStart} {
			if v < -bounds {
				skip = true
			}
		}

		for _, v := range []int{xEnd, yEnd, zEnd} {
			if v > bounds {
				skip = true
			}
		}
		if skip {
			continue
		}

		cube := Cube{
			Lower: Coord{xStart, yStart, zStart},
			Upper: Coord{xEnd, yEnd, zEnd},
		}

		instructions = append(instructions, Instruction{
			S: s,
			C: cube,
		})

		if s == ON {
			onCubes = append(onCubes, cube)
		} else {
			offCubes = append(offCubes, cube)
		}
	}

	return instructions, onCubes, offCubes
}

func run(input string, bounds int) uint64 {
	instructions, _, _ := parseInput(input, bounds)

	cubesOn := make(map[Cube]bool)
	for _, inst := range instructions {
		var cubesToAdd []Cube
		if inst.S {
			cubesToAdd = append(cubesToAdd, inst.C)
		}

		for on := range cubesOn {
			overlap := inst.C.Overlap(on)
			if overlap == nil {
				continue
			}

			// Removing from our collection of cubesOn because we've found an overlap
			// This means that one cube is going to split into possibly a few after
			// removing the overlap
			delete(cubesOn, on)

			// We then subtract the overlap and append each resulting into the
			// cubesToAdd list of this iteration.
			cubesToAdd = append(cubesToAdd, on.Subtract(*overlap)...)
		}

		for _, c := range cubesToAdd {
			cubesOn[c] = true
		}
	}

	var volume uint64
	for cube := range cubesOn {
		volume += cube.Volume(50)
	}
	return volume
}
