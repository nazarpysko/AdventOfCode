package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func rockScissorsPaper(input []string) {
	totalScore := 0
	for _, play := range input {
		totalScore += scorePlay[play]
	}

	fmt.Println("Total score:", totalScore)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(2))

	timer := time.Now()
	rockScissorsPaper(input)
	fmt.Println("Time:", time.Since(timer))
}
