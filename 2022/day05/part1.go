package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func cargoCranePart1(input []string) {
	crates, instructionsStart := parseCrane(input)
	for _, i := range input[instructionsStart:] {
		inst := parseInstruction(i)
		crates.moveCrates(inst, crates.popReversed)
	}

	fmt.Println("Top crates:", crates.getTopCrates())
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(05))

	timer := time.Now()
	cargoCranePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
