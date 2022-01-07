package main

import (
	"strings"
)

type Pixel struct {
	X int
	Y int
}

type Image struct {
	Known    map[Pixel]bool
	Min      int
	Max      int
	InfPixel bool
}

func (i Image) countLight() int {
	sum := 0
	for _, v := range i.Known {
		if v {
			sum++
		}
	}
	return sum
}

// func (i Image) String() string {

// 	grid := make([][]string, height+1)

// 	// Create empty grid of strings
// 	for j := 0; j <= height; j++ {
// 		grid[j] = make([]string, width+1)
// 		for i := range grid[j] {
// 			grid[j][i] = "."
// 		}
// 	}

// 	// Need to mark light cells and normalize
// 	var adjX, adjY int
// 	if lX > 0 {
// 		adjX = -lX
// 	} else {
// 		adjX = int(math.Abs(float64(lX)))

// 	}

// 	if lY > 0 {
// 		adjY = -lY
// 	} else {
// 		adjY = int(math.Abs(float64(lY)))
// 	}
// 	for p := range i.Known {
// 		grid[p.Y+adjY][p.X+adjX] = "#"
// 	}

// 	// Combine all the strings into a mega-grid string
// 	var s strings.Builder
// 	for _, row := range grid {
// 		s.WriteString(strings.Join(row, "") + "\n")
// 	}

// 	return s.String()
// }

func (i Image) enhance(algo [512]bool) Image {
	newMin := i.Min - 1
	newMax := i.Max + 1

	// Make new List and new Known, starting out with
	// allocation sized at minimum of their predecessors
	newKnown := make(map[Pixel]bool, len(i.Known))

	// -1 and +1 because we want to be bounded one further in every
	// direction of the current min/max X and min/max Y.
	for y := newMin; y < newMax; y++ {
		for x := newMin; x < newMax; x++ {

			// Create new pixel
			newPixel := Pixel{x, y}

			// Get decimal value from binary string
			v := i.adjacent(newPixel)

			// If algo at this index is true, add
			// pixel to the lookup
			newKnown[newPixel] = algo[v]
		}
	}

	var newInfPixel bool
	if i.InfPixel {
		newInfPixel = false
	} else {
		newInfPixel = true
	}

	return Image{
		Known:    newKnown,
		Min:      newMin,
		Max:      newMax,
		InfPixel: newInfPixel,
	}
}

func (i Image) adjacent(p Pixel) int {
	var value int
	for _, y := range []int{-1, 0, 1} {
		for _, x := range []int{-1, 0, 1} {
			value = value << 1
			// Look up in map if this Pixel exists
			if val, ok := i.Known[Pixel{x + p.X, y + p.Y}]; ok && val {
				value++

				// Otherwise, it's maybe an infinite pixel or an edge-ish pixel
			} else if !ok && i.InfPixel {
				value++
			}
		}
	}

	return value
}

func parseInput(input string) ([512]bool, Image) {
	split := strings.Split(input, "\n\n")
	algoRaw := strings.TrimLeft(split[0], "\t ")

	algo := [512]bool{}
	for i, chr := range algoRaw {
		switch chr {
		case '#':
			algo[i] = true
		case '.':
			algo[i] = false
		}
	}

	imageRaw := split[1]
	imageLines := strings.Split(imageRaw, "\n")
	known := make(map[Pixel]bool, len(imageLines)*4)
	image := Image{
		Known:    known,
		Min:      0,
		Max:      100,
		InfPixel: false,
	}

	for i, line := range imageLines {
		for j, chr := range line {
			nP := Pixel{j, i}
			if chr == '#' {

				image.Known[nP] = true
			} else {
				image.Known[nP] = false
			}
		}
	}

	return algo, image
}

func run(input string, iterations int) int {
	algo, image := parseInput(input)

	// log.Printf("\n%s", image.String())
	for i := 0; i < iterations; i++ {
		image = image.enhance(algo)
		// log.Printf("\n%s", image.String())
	}

	// log.Printf("\n%s", image.String())
	return image.countLight()
}
