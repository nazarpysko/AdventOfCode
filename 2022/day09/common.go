package main

import (
	"github.com/nazarpysko/AoC/2022/utils"
	"strconv"
)

type motion struct {
	direction string
	distance  int
}

func parseMotion(m string) motion {
	direction := string(m[0])
	distance, _ := strconv.Atoi(string(m[2]))
	return motion{
		direction: direction,
		distance:  distance,
	}
}

type coordinate struct {
	x int
	y int
}

type grid struct {
	visitedTailPositions map[coordinate]bool
	currentHead          coordinate
	currentTail          coordinate
}

func newGrid() grid {
	grid := grid{
		visitedTailPositions: make(map[coordinate]bool, 0),
		currentHead:          coordinate{0, 0},
		currentTail:          coordinate{0, 0},
	}

	grid.visitedTailPositions[grid.currentTail] = true
	return grid
}

func (g *grid) moveHead(x, y int) {
	g.currentHead.x += x
	g.currentHead.y += y
}

func (g *grid) moveTail(x, y int) {
	g.currentTail.x += x
	g.currentTail.y += y
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

		dx, dy := g.currentHead.x-g.currentTail.x, g.currentHead.y-g.currentTail.y

		if utils.Abs(dx) <= 1 && utils.Abs(dy) <= 1 {
			continue
		}

		g.moveTail(utils.Sign(dx), utils.Sign(dy))

		g.visitedTailPositions[g.currentTail] = true
	}
}

func (g *grid) getPositionsVisitedByTails() int {
	return len(g.visitedTailPositions)
}
