package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart2(input []string) {
	program := newProgram()
	for _, ins := range input {
		program.execute(ins)
	}

	fmt.Println(program.crp.drawing)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(10))

	timer := time.Now()
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
