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

func descend(depth int, planets []*Planet, orbitalMap map[*Planet]*Orbit) {
	for _, planet := range planets {
		planet.Distance = depth

		if orbiters, ok := orbitalMap[planet]; ok {
			descend(depth+1, orbiters.Orbits, orbitalMap)
		}
	}
}

func find(search string, root_planet *Planet, orbitalMap map[*Planet]*Orbit) []*Planet {
	planets := orbitalMap[root_planet].Orbits
	for _, planet := range planets {
		if forward_path, ok := path(search, planet, orbitalMap); ok {
			return append([]*Planet{root_planet}, forward_path...)
		}
	}
	return nil
}

func path(search string, current_planet *Planet, orbitalMap map[*Planet]*Orbit) ([]*Planet, bool) {
	planets := orbitalMap[current_planet].Orbits
	for _, planet := range planets {

		// if planet is the one we are searching for,
		// . return immediately and start constructing path
		if planet.Name == search {
			return []*Planet{current_planet, planet}, true
		}

		// If planet has no orbiters, skip
		_, ok := orbitalMap[planet]
		if !ok {
			continue
		}

		if forward_path, ok := path(search, planet, orbitalMap); ok {
			return append([]*Planet{current_planet}, forward_path...), true
		}
	}

	return nil, false
}

func run(input string) (int, int) {

	orbitalMap := make(map[*Planet]*Orbit)
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
			orbitalMap[p1] = &Orbit{
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

	delta_distance := 0
	if _, ok := planetMap["YOU"]; ok {
		you_path := find("YOU", com, orbitalMap)
		santa_path := find("SAN", com, orbitalMap)

		if you_path == nil || santa_path == nil {
			return sum, 0
		}

		// Traverse fewer nodes
		seen := make(map[*Planet]int)
		var shorter, longer []*Planet
		if len(you_path) < len(santa_path) {
			shorter = you_path
			longer = santa_path
		} else {
			shorter = you_path
			longer = santa_path
		}

		// Add seen planets to seen set
		for i, planet := range shorter {
			seen[planet] = i
		}

		// Determine farthest intersecting planet in paths
		max := -1
		var max_planet *Planet
		for _, planet := range longer {
			if j, ok := seen[planet]; ok && j > max {
				max = j
				max_planet = planet
			}
		}

		you_distance := planetMap["YOU"].Distance - 1
		santa_distance := planetMap["SAN"].Distance - 1
		max_distance := max_planet.Distance
		delta_distance = you_distance - max_distance + santa_distance - max_distance
	}

	return sum, delta_distance
}
