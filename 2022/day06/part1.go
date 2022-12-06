package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	pos := getPositionMaker(input[0], 4)
	fmt.Println("Position:", pos)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(06))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
