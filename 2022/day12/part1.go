package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	startingNode := findCoordinates(input, "S")
	fmt.Println("Shortest path:", bfs(input, startingNode, "E"))
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(12))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
