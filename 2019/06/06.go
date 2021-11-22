package main

import (
	"strings"
)

type Planet struct {
	Name     string
	Distance int
}

type Orbit struct {
	Planet *Planet
	Orbits []*Planet
}

func descend(depth int, planets []*Planet, orbitalMap map[*Planet]Orbit) {
	for _, planet := range planets {
		planet.Distance = depth

		if orbiters, ok := orbitalMap[planet]; ok {
			descend(depth+1, orbiters.Orbits, orbitalMap)
		}
	}
}

func run(input string) int {

	orbitalMap := make(map[*Planet]Orbit)
	planetMap := make(map[string]*Planet)

	orbits := strings.Split(input, "\n")

	for _, orbit := range orbits {
		split := strings.Split(orbit, ")")
		orbitee := split[0]
		orbiter := split[1]

		// Create the two planets if they don't exist
		// . otherwise pull them from the map
		var p1, p2 *Planet
		if _, ok := planetMap[orbitee]; !ok {
			p := Planet{
				Name: orbitee,
			}
			planetMap[orbitee] = &p
			p1 = &p
		} else {
			p1 = planetMap[orbitee]
		}

		if _, ok := planetMap[orbiter]; !ok {
			p := Planet{
				Name: orbiter,
			}
			planetMap[orbiter] = &p
			p2 = &p
		} else {
			p2 = planetMap[orbiter]
		}

		// Create the new orbit, appending if the orbit
		// . already exists
		if o1, ok := orbitalMap[p1]; ok {
			o1.Orbits = append(o1.Orbits, p2)
		} else {
			orbitalMap[p1] = Orbit{
				Planet: p1,
				Orbits: []*Planet{p2},
			}
		}
	}

	com := planetMap["COM"]
	com_orbits := orbitalMap[com]
	descend(1, com_orbits.Orbits, orbitalMap)

	sum := 0
	for _, v := range planetMap {
		sum += v.Distance
	}

	return sum
}
