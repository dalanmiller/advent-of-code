package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Scanner struct {
	N       int
	Beacons []Beacon
	// Coordinates relative to Scanner 0
	X int
	Y int
	Z int
	// overlaps []*Scanner
}

type Rotation int

const (
	XAxis Rotation = iota
	YAxis
	ZAxis
)

// func (s Scanner) rotate(r Rotation) Scanner {
// 	newScanner := s
// 	newScanner.Beacons = make([]Beacon, len(s.Beacons))
// 	copy(newScanner.Beacons, s.Beacons)

// 	for i, b := range newScanner.Beacons {
// 		newScanner.Beacons[i] = b.rotate(r)
// 	}

// 	return newScanner
// }

// type BeaconPair struct {
// 	A Beacon
// 	B Beacon
// }

// type DistanceSet struct {
// 	storage     []float64
// 	contentHash map[float64]bool
// }

// func (ds *DistanceSet) Add(i float64) {
// 	_, ok := ds.contentHash[i]
// 	if !ok {
// 		x := sort.Search(len(ds.storage), func(j int) bool { return ds.storage[j] >= i })

// 		if x == len(ds.storage) {
// 			ds.storage = append(ds.storage, i)
// 		} else {
// 			ds.storage = append(ds.storage[:x+1], ds.storage[x:]...)
// 			ds.storage[x] = i
// 		}

// 		ds.contentHash[i] = true
// 	}
// }

// func (ds *DistanceSet) Intersection(b *DistanceSet) int {
// 	count := 0
// 	for i, _ := range ds.contentHash {
// 		if _, ok := b.contentHash[i]; ok {
// 			count++
// 		}
// 	}
// 	return count
// }

type Beacon struct {
	X int
	Y int
	Z int
	// distances DistanceSet
}

// func (b Beacon) rotate(r Rotation) Beacon {

// 	newBeacon := Beacon{
// 		X: b.X,
// 		Y: b.Y,
// 		Z: b.Z,
// 	}

// 	switch r {
// 	case XAxis:
// 		// 1,1,1 => 1,1,-1
// 		// 1,1,-1 => -1,1,-1
// 		// -1,1,-1 => -1,1,1
// 		// -1,1,1 => 1,1,1
// 		if b.Y > 0 && b.Z > 0 || b.Y < 0 && b.Z < 0 {
// 			newBeacon.Y = -newBeacon.Y
// 		} else if b.Y > 0 && b.Z < 0 || b.Y < 0 && b.Z > 0 {
// 			newBeacon.Z = -newBeacon.Z
// 		}
// 	case YAxis:
// 		// 1,1,1 => 1,1,-1
// 		// 1,1,-1 => -1,1,-1
// 		// -1,1,-1 => -1,1,1
// 		// -1,1,1 => 1,1,1
// 		if b.X > 0 && b.Z < 0 || b.X < 0 && b.Z > 0 {
// 			newBeacon.X = -newBeacon.X
// 		} else if b.X < 0 && b.Z < 0 || b.X > 0 && b.Z > 0 {
// 			newBeacon.Z = -newBeacon.Z
// 		}
// 	case ZAxis:
// 		// 1,1,1 => 1,-1,1
// 		// 1,-1,1 => -1,-1,1
// 		// -1,-1,1 => -1,1,1
// 		// -1,1,1 => 1,1,1
// 		if b.Y > 0 && b.X > 0 || b.Y < 0 && b.X < 0 {
// 			newBeacon.Y = -newBeacon.Y
// 		} else if b.Y > 0 && b.X < 0 || b.Y < 0 && b.X > 0 {
// 			newBeacon.X = -newBeacon.X
// 		}
// 	}

// 	return newBeacon
// }

// func (s *Scanner) generateOrientations() []Scanner {
// 	// 24 different orientations
// 	orientations := make([]Scanner, 0, 24)

// 	// Put a smiley face on a d6.
// 	// You can make the smiley face point in six different directions, and each direction you can rotate the smiley face four different ways.
// 	// 6 * 4 = 24.
// 	// https://www.reddit.com/r/adventofcode/comments/rjs8rd/comment/hp5jopz/

// 	cur := *s
// 	cur.Beacons = make([]Beacon, len((*s).Beacons))
// 	copy(cur.Beacons, (*s).Beacons)
// 	orientations = append(orientations, cur)
// 	for i := 0; i < 3; i++ {
// 		cur = cur.rotate(XAxis)
// 		orientations = append(orientations, cur)
// 	}

// 	cur = cur.rotate(ZAxis)
// 	orientations = append(orientations, cur)
// 	for i := 0; i < 3; i++ {
// 		cur = cur.rotate(YAxis)
// 		orientations = append(orientations, cur)
// 	}

// 	cur = cur.rotate(ZAxis)
// 	orientations = append(orientations, cur)
// 	for i := 0; i < 3; i++ {
// 		cur = cur.rotate(XAxis)
// 		orientations = append(orientations, cur)
// 	}

// 	cur = cur.rotate(ZAxis)
// 	orientations = append(orientations, cur)
// 	for i := 0; i < 3; i++ {
// 		cur = cur.rotate(YAxis)
// 		orientations = append(orientations, cur)
// 	}

// 	cur = cur.rotate(XAxis)
// 	orientations = append(orientations, cur)
// 	for i := 0; i < 3; i++ {
// 		cur = cur.rotate(ZAxis)
// 		orientations = append(orientations, cur)
// 	}

// 	cur = cur.rotate(XAxis).rotate(XAxis)
// 	orientations = append(orientations, cur)
// 	for i := 0; i < 3; i++ {
// 		cur = cur.rotate(ZAxis)
// 		orientations = append(orientations, cur)
// 	}

// 	return orientations
// }

// func intersectBeacons(a Scanner, b Scanner) int {
// 	seenA := make(map[Beacon]bool, len(a.Beacons))
// 	for _, b := range a.Beacons {
// 		seenA[b] = true
// 	}

// 	seenAB := make(map[Beacon]bool, len(seenA))
// 	for _, b := range b.Beacons {
// 		_, ok := seenA[b]
// 		if ok {
// 			seenAB[b] = true
// 		}
// 	}

// 	return len(seenAB)
// }

var digits = regexp.MustCompile(`\d+`)

func parseInput(input string) []*Scanner {
	lines := strings.Split(input, "\n")

	var scanners []*Scanner
	var currentScanner *Scanner
	for _, line := range lines {
		if line == "\n" || line == "" {
			continue
		}

		if strings.HasPrefix(line, "--") {
			digit := digits.FindString(line)
			v, _ := strconv.Atoi(digit)
			currentScanner = &Scanner{
				N: v,
				// distances: DistanceSet{
				// 	contentHash: map[float64]bool{},
				// 	storage:     []float64{},
				// },
			}
			if v == 0 {
				currentScanner.X = 0
				currentScanner.Y = 0
				currentScanner.Z = 0
			}
			scanners = append(scanners, currentScanner)
		} else {
			coord := strings.Split(line, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			z, _ := strconv.Atoi(coord[2])

			currentScanner.Beacons = append(
				currentScanner.Beacons,
				Beacon{
					X: x, Y: y, Z: z,
				},
			)
		}
	}

	return scanners
}

// func incompleteMapping(m map[int]bool) bool {
// 	for _, mapped := range m {
// 		if !mapped {
// 			return true
// 		}
// 	}

// 	return false
// }

// func calculateBeaconToBeaconDistances(scanners []*Scanner) {
// 	for _, scanner := range scanners {
// 		for i, bA := range scanner.Beacons {
// 			for j, bB := range scanner.Beacons {
// 				if i == j {
// 					continue
// 				}

// 				manhattanDistance := math.Sqrt(math.Pow(float64(bA.X-bB.X), 2) + math.Pow(float64(bA.Y-bB.Y), 2) + math.Pow(float64(bA.Z-bB.Z), 2))
// 				// scanner.distances.Add(manhattanDistance)
// 				bA.distances.Add(manhattanDistance)
// 				bB.distances.Add(manhattanDistance)
// 			}
// 		}
// 	}
// }

// type Match struct {
// 	S1 *Scanner
// 	S2 *Scanner
// 	B1 Beacon
// 	B2 Beacon
// }

// func findMatches(scannerMap map[int]*Scanner) []Match {
// 	matches := []Match{}
// 	for i, s1 := range scannerMap {
// 		for j, s2 := range scannerMap {
// 			if j <= i {
// 				continue
// 			}

// 			for _, b1 := range s1.Beacons {
// 				for _, b2 := range s2.Beacons {
// 					if b1.distances.Intersection(&b2.distances) >= 12 {
// 						matches = append(matches, Match{
// 							S1: s1,
// 							S2: s2,
// 							B1: b1,
// 							B2: b2,
// 						})
// 						if len(matches) == 12 {
// 							return matches
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return matches
// }

type Axis int

const (
	Xx Axis = 0
	Yx      = 1
	Zx      = 2
)

type Sign int

const (
	POS Sign = 1
	NEG Sign = -1
)

type Diff struct {
	Axis  Axis
	Sign  Sign
	Delta int
}

func generateDiffs(s1 *Scanner, scannerMap map[int]*Scanner) (map[int]Diff, map[int]Diff, map[int]Diff) {
	xDiff, yDiff, zDiff := make(map[int]Diff), make(map[int]Diff), make(map[int]Diff)

	for _, s2 := range scannerMap {
		if s1.N == s2.N {
			continue
		}

		for axis := range []Axis{Xx, Yx, Zx} {
			for sign := range []Sign{POS, NEG} {
				counter := make(map[int]int, len(s1.Beacons)*len(s2.Beacons))

				for _, b1 := range s1.Beacons {
					for _, b2 := range s2.Beacons {
						counter[b1.X-b2.X*sign]++
					}
				}

				var n, max int
				for x, y := range counter {
					if y > max {
						max = y
						n = x
					}
				}

				if max >= 12 {
					xDiff[s2.N] = Diff{
						Axis:  Axis(axis),
						Sign:  Sign(sign),
						Delta: n,
					}
				}
			}
		}
	}

	for otherScannerN := range xDiff {
		s2 := scannerMap[otherScannerN]
		for axis := range []Axis{Xx, Yx, Zx} {
			for sign := range []Sign{POS, NEG} {
				yCounter := make(map[int]int, len(s1.Beacons)*len(s2.Beacons))
				zCounter := make(map[int]int, len(s1.Beacons)*len(s2.Beacons))

				for _, b1 := range s1.Beacons {
					for _, b2 := range s2.Beacons {
						yCounter[b1.Y-b2.Y*sign]++
						zCounter[b1.Z-b2.Z*sign]++
					}
				}

				var nY, maxY, nZ, maxZ int
				for x, y := range yCounter {
					if y > maxY {
						maxY = y
						nY = x
					}
				}

				if maxY >= 12 {
					yDiff[s2.N] = Diff{
						Axis:  Axis(axis),
						Sign:  Sign(sign),
						Delta: nY,
					}
				}

				for x, y := range zCounter {
					if y > maxZ {
						maxZ = y
						nZ = x
					}
				}

				if maxZ >= 12 {
					zDiff[s2.N] = Diff{
						Axis:  Axis(axis),
						Sign:  Sign(sign),
						Delta: nZ,
					}
				}
			}
		}
	}

	return xDiff, yDiff, zDiff
}

func run(input string) int {
	scanners := parseInput(input)

	scannerMap := make(map[int]*Scanner)
	for _, scanner := range scanners {
		scannerMap[scanner.N] = scanner
	}

	allBeacons := make(map[Beacon]bool)
	for _, b := range scanners[0].Beacons {
		allBeacons[b] = true
	}

	toScan := []*Scanner{scanners[0]}
	for len(toScan) > 0 {
		current := toScan[0]
		toScan = toScan[1:]

		xD, yD, zD := generateDiffs(current, scannerMap)
		for i := range xD {
			dstX := xD[i].Delta
			dstY := yD[i].Delta
			dstZ := zD[i].Delta

			scannerMap[i].X = dstX
			scannerMap[i].Y = dstY
			scannerMap[i].Z = dstZ

			next := scannerMap[i]
			delete(scannerMap, i)
			newBeacons := []Beacon{}
			for _, b := range next.Beacons {
				var xi, yi, zi int
				switch xD[i].Axis {
				case Xx:
					xi = b.X
				case Yx:
					xi = b.Y
				case Zx:
					xi = b.Z
				}

				switch yD[i].Axis {
				case Xx:
					yi = b.X
				case Yx:
					yi = b.Y
				case Zx:
					yi = b.Z
				}

				switch zD[i].Axis {
				case Xx:
					zi = b.X
				case Yx:
					zi = b.Y
				case Zx:
					zi = b.Z
				}

				newBeacons = append(newBeacons, Beacon{
					X: dstX + int(xD[i].Sign)*xi,
					Y: dstY + int(yD[i].Sign)*yi,
					Z: dstZ + int(zD[i].Sign)*zi,
				})
			}

			next.Beacons = newBeacons
			for _, b := range next.Beacons {
				allBeacons[b] = true
			}
			toScan = append(toScan, next)
		}
	}

	return len(allBeacons)
}
