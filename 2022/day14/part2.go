package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart2(input []string) {
	// Make the cave
	cave := (&Cave{}).newCave()

	for _, path := range input {
		cave.addPath(path)
	}

	cave.addPath(fmt.Sprintf("0,%d -> 8000,%d", cave.deepestLevel+2, cave.deepestLevel+2))

	sandGrains := 0

	// Pour the sand until source is sand
	for {
		if cave.isSourceSand() || !cave.pourSand() {
			break
		}

		sandGrains += 1
	}

	fmt.Println("Total units of sand until source is sand:", sandGrains)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(14))

	timer := time.Now()
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
