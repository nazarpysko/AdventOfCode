package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func cargoCranePart1(input []string) {
	crates, instructionsStart := parseCrane(input)
	fmt.Println(crates[0], instructionsStart)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(05))

	timer := time.Now()
	cargoCranePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
