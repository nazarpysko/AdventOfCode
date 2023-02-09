package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strconv"
	"strings"
)

const (
	sand = "."
	air  = "o"
	rock = "#"
)

type Coordinate struct {
	x, y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

type Cave struct {
	structure    map[Coordinate]string
	deepestLevel int
}

func (c *Cave) newCave() *Cave {
	return &Cave{
		structure:    make(map[Coordinate]string),
		deepestLevel: 0,
	}
}

func (c *Cave) setDeepestLevel(coordinate Coordinate) {
	if coordinate.y > c.deepestLevel {
		c.deepestLevel = coordinate.y
	}
}

func (c *Cave) updateCave(coordinate Coordinate) {
	(*c).structure[coordinate] = rock
	(*c).setDeepestLevel(coordinate)
}

func (c *Cave) String() string {
	var str string
	for c, value := range (*c).structure {
		str += fmt.Sprintf("(%d, %d) Value: %s\n", c.x, c.y, value)
	}
	return str
}

func (c *Cave) addPath(path string) {
	points := strings.Split(path, " -> ")
	c.updateCave(getCoordinate(points[0]))

	for i, point := range points[1:] {
		currentCoordinate := getCoordinate(point)

		if _, exists := (*c).structure[currentCoordinate]; exists {
			continue
		}

		// As points are sliced to points[1:], i is offset by 1. If it is used, it will be point to previous point
		lastCoordinate := getCoordinate(points[i])

		dx := lastCoordinate.x - currentCoordinate.x
		dy := lastCoordinate.y - currentCoordinate.y

		if dx != 0 {
			for x := currentCoordinate.x; x != lastCoordinate.x; x = x + 1*utils.Sign(dx) {
				c.updateCave(Coordinate{x, currentCoordinate.y})
			}
		} else {
			for y := currentCoordinate.y; y != lastCoordinate.y; y = y + 1*utils.Sign(dy) {
				c.updateCave(Coordinate{currentCoordinate.x, y})
			}
		}
	}
}

func getCoordinate(point string) Coordinate {
	c := strings.Split(point, ",")

	x, _ := strconv.Atoi(c[0])
	y, _ := strconv.Atoi(c[1])

	return Coordinate{x, y}
}

// pourSand returns true if the thrown grain of sand comes to a rest or false if it reaches the abyss
func (c *Cave) pourSand() bool {
	target := Coordinate{500, 0}
	restFound := false

	for target.y < c.deepestLevel {
		target, restFound = c.cameToRest(target)
		if restFound {
			return true
		}
	}

	return false
}

// cameToRest returns a bool if the sand coordinate comes to a rest and a new coordinate if rest is not found
func (c *Cave) cameToRest(sandPoint Coordinate) (Coordinate, bool) {
	startingSandPoint := sandPoint
	sandPoint.y++
	// Can go down
	if _, exists := c.structure[sandPoint]; !exists {
		return sandPoint, false
	}

	sandPoint.x--
	// Can go down left
	if _, exists := c.structure[sandPoint]; !exists {
		return sandPoint, false
	}

	sandPoint.x += 2
	// Can go down right
	if _, exists := c.structure[sandPoint]; !exists {
		return sandPoint, false
	}

	c.structure[startingSandPoint] = sand
	return sandPoint, true
}

func (c *Cave) isSourceSand() bool {
	return c.structure[Coordinate{500, 0}] == sand
}
