package main

import (
	"github.com/nazarpysko/AoC/2022/utils"
	"strconv"
)

type motion struct {
	direction string
	distance  int
}

type coordinate struct {
	x int
	y int
}

type grid struct {
	visitedTailPositions map[coordinate]bool
	head                 coordinate
	tail                 []coordinate
}

func newGrid(tailLenght int) grid {
	grid := grid{
		visitedTailPositions: make(map[coordinate]bool, 0),
		head:                 coordinate{0, 0},
		tail:                 make([]coordinate, tailLenght),
	}

	grid.updateTailVisitedPositions()
	return grid
}

func (g *grid) moveHead(x, y int) {
	g.head.x += x
	g.head.y += y
}

func (g *grid) moveTail(segment, x, y int) {
	g.tail[segment].x += x
	g.tail[segment].y += y
}

func (g *grid) move(m motion) {
	var x, y int

	switch m.direction {
	case "U":
		x, y = 0, 1
	case "D":
		x, y = 0, -1
	case "R":
		x, y = 1, 0
	case "L":
		x, y = -1, 0
	}

	g.moveRope(m.distance, x, y)
}

func (g *grid) moveRope(distance, x, y int) {
	for i := 0; i < utils.Abs(distance); i++ {
		g.moveHead(x, y)

		for tailSegment := 0; tailSegment < len(g.tail); tailSegment++ {
			var dx, dy int

			if tailSegment == 0 {
				dx, dy = g.head.x-g.tail[0].x, g.head.y-g.tail[0].y
			} else {
				dx, dy = g.tail[tailSegment-1].x-g.tail[tailSegment].x, g.tail[tailSegment-1].y-g.tail[tailSegment].y
			}

			if utils.Abs(dx) <= 1 && utils.Abs(dy) <= 1 {
				continue
			}

			g.moveTail(tailSegment, utils.Sign(dx), utils.Sign(dy))

		}

		g.updateTailVisitedPositions()
	}
}

func (g *grid) updateTailVisitedPositions() {
	lastTailSegment := len(g.tail) - 1
	g.visitedTailPositions[g.tail[lastTailSegment]] = true
}

func (g *grid) getPositionsVisitedByTails() int {
	return len(g.visitedTailPositions)
}

func parseMotion(m string) motion {
	direction := string(m[0])
	distance, _ := strconv.Atoi(m[2:])
	return motion{
		direction: direction,
		distance:  distance,
	}
}
