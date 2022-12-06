package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func cargoCranePart2(input []string) {
	crates, instructionsStart := parseCrane(input)
	for _, i := range input[instructionsStart:] {
		inst := parseInstruction(i)
		crates.moveCrates(inst, crates.pop)
	}

	fmt.Println("Top crates:", crates.getTopCrates())
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(05))

	timer := time.Now()
	cargoCranePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
