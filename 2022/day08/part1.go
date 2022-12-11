package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	visibleTrees := countVisibleTrees(input)
	fmt.Println("Total visible trees:", visibleTrees)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(8))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
