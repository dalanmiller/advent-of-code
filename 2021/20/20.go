package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Pixel struct {
	X int
	Y int
}

type Image struct {
	Lookup map[Pixel]bool
	LeastX int
	MaxX   int
	LeastY int
	MaxY   int
}

func (i Image) Edges() (int, int, int, int) {
	return i.LeastX, i.MaxX, i.LeastY, i.MaxY
}

func (i Image) countLight() int {
	return len(i.Lookup)
}

func (i Image) String() string {
	lX, mX, lY, mY := i.Edges()
	height := mY - lY
	width := mX - lX
	grid := make([][]string, mY-lY+1)

	// Create empty grid of strings
	for j := 0; j <= height; j++ {
		grid[j] = make([]string, width+1)
		for i := range grid[j] {
			grid[j][i] = "."
		}
	}

	// Need to mark light cells and normalize
	var adjX, adjY int
	if lX > 0 {
		adjX = -lX
	} else {
		adjX = int(math.Abs(float64(lX)))

	}

	if lY > 0 {
		adjY = -lY
	} else {
		adjY = int(math.Abs(float64(lY)))
	}
	for p, _ := range i.Lookup {
		grid[p.Y+adjY][p.X+adjX] = "#"
	}

	// Combine all the strings into a mega-grid string
	var s strings.Builder
	for _, row := range grid {
		s.WriteString(strings.Join(row, "") + "\n")
	}

	return s.String()
}

func (i Image) enhance(algo [512]bool) Image {
	lX, mX, lY, mY := i.Edges()

	// Make new List and new Lookup, starting out with
	// allocation sized at minimum of their predecessors
	newLookup := make(map[Pixel]bool, len(i.Lookup))

	// Set initial values for new min/max x/y
	nlX, nmX, nlY, nmY := math.MaxInt, 0, math.MaxInt, 0

	// -1 and +1 because we want to be bounded one further in every
	// direction of the current min/max X and min/max Y.
	for y := lY - 1; y <= mY+1; y++ {
		for x := lX - 1; x <= mX+1; x++ {

			// Create new pixel
			newPixel := Pixel{x, y}

			// Get decimal value from binary string
			v := i.adjacent(newPixel)

			// If algo at this index is true, add
			// pixel the new list and lookup
			if algo[v] {
				newLookup[newPixel] = true
			}

			// Determine if we need to set a new min/max
			if newPixel.X < nlX {
				nlX = newPixel.X
			}

			if newPixel.X > nmX {
				nmX = newPixel.X
			}

			if newPixel.Y < nlY {
				nlY = newPixel.Y
			}

			if newPixel.Y > nmY {
				nmY = newPixel.Y
			}
		}
	}

	return Image{
		Lookup: newLookup,
		LeastX: nlX,
		MaxX:   nmX,
		LeastY: nlY,
		MaxY:   nmY,
	}
}

func (i Image) adjacent(p Pixel) int {
	values := make([]int, 0, 8)
	for _, y := range []int{-1, 0, 1} {
		for _, x := range []int{-1, 0, 1} {

			// Look up in map if this Pixel exists
			if _, ok := i.Lookup[Pixel{x + p.X, y + p.Y}]; ok {
				values = append(values, 1)

				// Otherwise, it's a 0
			} else {
				values = append(values, 0)
			}
		}
	}

	// Do configuration to make this into a representation of the binary value
	var v strings.Builder
	for _, val := range values {
		v.WriteString(strconv.Itoa(val))
	}
	value, _ := strconv.ParseInt(v.String(), 2, 32)
	return int(value)
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
	lookup := make(map[Pixel]bool, len(imageLines))
	image := Image{
		Lookup: lookup,
		LeastX: math.MaxInt,
		MaxX:   0,
		LeastY: math.MaxInt,
		MaxY:   0,
	}

	for i, line := range imageLines {
		for j, chr := range line {
			if chr == '#' {
				nP := Pixel{j, i}
				image.Lookup[nP] = true

				if j < image.LeastX {
					image.LeastX = j
				}

				if j > image.MaxX {
					image.MaxX = j
				}

				if i < image.LeastY {
					image.LeastY = i
				}

				if i > image.MaxY {
					image.MaxY = i
				}
			}
		}
	}
	return algo, image
}

func run(input string, iterations int) int {
	algo, image := parseInput(input)

	log.Printf("\n%s", image.String())
	for i := 0; i < iterations; i++ {
		image = image.enhance(algo)

	}

	log.Printf("\n%s", image.String())
	return image.countLight()
}
