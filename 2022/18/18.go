package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type axis int

const (
	X axis = iota
	Y
	Z
)

type cube [3]int

func (c *cube) neighbors() []cube {
	cubes := make([]cube, 0, 6)

	for _, dx := range []int{-1, 1} {
		cubes = append(cubes, cube{c[X] + dx, c[Y], c[Z]})
	}

	for _, dy := range []int{-1, 1} {
		cubes = append(cubes, cube{c[X], c[Y] + dy, c[Z]})
	}

	for _, dz := range []int{-1, 1} {
		cubes = append(cubes, cube{c[X], c[Y], c[Z] + dz})
	}

	return cubes
}

func readInput(input io.Reader) (cubes []cube) {

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[X])
		y, _ := strconv.Atoi(split[Y])
		z, _ := strconv.Atoi(split[Z])

		cubes = append(cubes, cube{x, y, z})
	}

	return cubes

}

func run(input io.Reader) (int, int) {
	cubes := readInput(input)

	cubeMap := make(map[cube]bool, len(cubes))

	min, max := cube{100, 100, 100}, cube{-100, -100, -100}
	for _, cube := range cubes {
		cubeMap[cube] = true

		for _, D := range []axis{X, Y, Z} {
			if cube[D] < min[D] {
				min[D] = cube[D]
			}

			if cube[D] > max[D] {
				max[D] = cube[D]
			}
		}
	}

	for _, D := range []axis{X, Y, Z} {
		min[D] -= 1
		max[D] += 1
	}

	c1 := 0
	for _, cube := range cubes {
		for _, n := range cube.neighbors() {
			if _, ok := cubeMap[n]; !ok {
				c1++
			}
		}
	}

	queue := []cube{min}
	visited := map[cube]bool{}
	for k, v := range cubeMap {
		visited[k] = v
	}
	c2 := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, n := range cur.neighbors() {
			if !(n[X] >= min[X] && n[X] <= max[X] && n[Y] >= min[Y] && n[Y] <= max[Y] && n[Z] >= min[Z] && n[Z] <= max[Z]) {
				continue
			}

			if _, ok := cubeMap[n]; ok {
				c2++
			}

			if _, visited := visited[n]; visited {
				continue
			}

			visited[n] = true
			queue = append(queue, n)
		}
	}

	return c1, c2
}
