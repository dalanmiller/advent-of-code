package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Pixel struct {
	X int
	Y int
	Z int
}

type Image [][]int

func (i Image) width() int {
	return len(i[0])
}

func (i Image) height() int {
	return len(i)
}

func (i Image) countLight() (c int) {
	for a, row := range i {
		for b := range row {
			if i[a][b] == 1 {
				c++
			}
		}
	}

	return c
}
func (i Image) String() (grid string) {
	for _, row := range i {
		o := make([]string, len(row))
		for i, chr := range row {
			o[i] = strconv.Itoa(chr)
		}
		grid += fmt.Sprintf("%s\n", strings.Join(o, ""))
	}

	return grid
}

func (i *Image) grow() {
	newWidth := i.width() + 2

	// Create new start row and make it the new width
	startBlankRow := [][]int{}
	startBlankRow = append(startBlankRow, make([]int, newWidth))

	// Add new start row 0s
	*i = append(startBlankRow, (*i)...)

	// Add 0 mid-row 0
	for y := 1; y < len(*i); y++ {

		// Prepend
		(*i)[y] = append([]int{0}, (*i)[y]...)
		// Add to end
		(*i)[y] = append((*i)[y], 0)
	}

	// Add new end row of 0s
	endBlankRow := make([]int, newWidth)
	*i = append(*i, endBlankRow)
}

func (i Image) enhance(algo [512]bool) Image {
	newImage := Image{}

	// Expand existing image
	newImage = append(newImage, make([]int, i.width()+2))
	for _, row := range i {
		// newRow := append([]int{0}, row...)
		// newRow = append(newRow, 0)
		newImage = append(newImage, make([]int, len(row)+2))
	}
	newImage = append(newImage, make([]int, i.width()+2))

	// Grow old image
	i.grow()

	// Read through adjacent cells from old Image to determine
	// new pixel on newImage
	for b, row := range newImage {
		for a := range row {
			if algo[i.adjacent(a, b)] {
				newImage[b][a] = 1
			} else {
				newImage[b][a] = 0
			}
		}
	}

	return newImage
}

func (i Image) adjacent(a, b int) int {
	pixels := []int{}
	for _, y := range []int{-1, 0, 1} {
		for _, x := range []int{-1, 0, 1} {
			if 0 <= x+a && x+a < i.width() && 0 <= y+b && y+b < i.height() {
				pixels = append(pixels, i[y+b][x+a])
			} else {
				pixels = append(pixels, 0)
			}
		}
	}

	v := make([]string, len(pixels))
	for i, p := range pixels {
		v[i] = strconv.Itoa(p)
	}
	value, _ := strconv.ParseInt(strings.Join(v, ""), 2, 32)
	return int(value)
}

func parseInput(input string) ([512]bool, []Pixel) {
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
	image := make(Image, 0, len(imageLines))
	for _, line := range imageLines {
		row := make([]int, len(line))
		for i, chr := range line {
			if chr == '#' {
				row[i] = 1
			} else if chr == '.' {
				row[i] = 0
			}
		}
		image = append(image, row)
	}
	return algo, image
}

func run(input string, n int) int {
	algo, image := parseInput(input)

	for i := 0; i < n; i++ {
		image = image.enhance(algo)
		log.Printf("\n%s", image.String())

	}

	return image.countLight()
}
