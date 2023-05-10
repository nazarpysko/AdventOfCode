package main

import (
	"regexp"
	"strconv"
)

const (
	beacon = "B"
	sensor = "S"
	void   = "#"
)

type Coordinate struct {
	x, y int
}

type x = int
type y = int

type Grid map[y][]x

func (g Grid) search(from, to Coordinate) {
	i := 1
	for {
		locationsDrawn := g.draw(from, i)
		if _, exists := locationsDrawn[to]; exists {
			break
		}
		i++
	}
}

func (g Grid) draw(from Coordinate, distance int) map[Coordinate]bool {
	drawn := make(map[Coordinate]bool)
	for i := 1; i <= distance; i++ {
		g[i] = append(g[i], from.x)
		g[-i] = append(g[i], from.x)
		g[from.x] = append(g[from.x], i, -i)

		drawn[Coordinate{from.x, i}] = true
		drawn[Coordinate{from.x, -i}] = true
		drawn[Coordinate{i, from.x}] = true
		drawn[Coordinate{-i, -from.x}] = true

		for j := i; j < distance; j++ {
			x, y := j, distance-j
			g[y] = append(g[y], x, -x)
			g[-y] = append(g[-y], x, -x)

			drawn[Coordinate{x, y}] = true
			drawn[Coordinate{x, -y}] = true
			drawn[Coordinate{-x, y}] = true
			drawn[Coordinate{-x, -y}] = true
		}
	}

	return drawn
}

func getSensorAndBeacon(line string) (Coordinate, Coordinate) {
	re := regexp.MustCompile(`x=(-?\d+), y=(-?\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	var x, y int

	x, _ = strconv.Atoi(matches[0][1])
	y, _ = strconv.Atoi(matches[0][2])
	sensor := Coordinate{x, y}

	x, _ = strconv.Atoi(matches[1][1])
	y, _ = strconv.Atoi(matches[1][2])
	beacon := Coordinate{x, y}

	return sensor, beacon
}
