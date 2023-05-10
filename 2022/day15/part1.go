package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	grid := make(Grid)

	for _, line := range input {
		sensor, beacon := getSensorAndBeacon(line)
		grid.search(sensor, beacon)
	}

	fmt.Println("Cannot contain beacon in y=2000000:", len(grid[2000000]))
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(15))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
