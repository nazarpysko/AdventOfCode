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

	sandGrains := 0

	// Pour the sand until source is sand
	for {
		if !cave.pourSand() {
			break
		}

		sandGrains += 1
	}

	fmt.Println("Total units of sand until reach the abyss:", sandGrains)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(14))

	timer := time.Now()
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
