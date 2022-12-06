package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func countGs(c crates) int {
	count := 0
	for _, V := range c {
		for _, i := range V {
			if string(i) == "G" {
				count += 1
			}
		}
	}

	return count
}

func cargoCranePart2(input []string) {
	crates, instructionsStart := parseCrane(input)
	for j, i := range input[instructionsStart:] {
		inst := parseInstruction(i)
		crates.moveCrates(inst, crates.pop)
		gs := countGs(crates)
		if gs > 4 {
			fmt.Println("Duplication detected in instruction %d (%d G's counted)", j, gs)
		}
	}

	fmt.Println("Top crates:", crates.getTopCrates())
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(05))

	timer := time.Now()
	cargoCranePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
