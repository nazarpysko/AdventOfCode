package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart2(input []string) {
	grid := newGrid(9)
	for _, line := range input {
		grid.move(parseMotion(line))
	}

	fmt.Println("Positions visited:", grid.getPositionsVisitedByTails())
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(9))

	timer := time.Now()
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
