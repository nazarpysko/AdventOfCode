package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func test(input []string) {
	totalScore := 0
	for _, play := range input {
		totalScore += scorePlay[play]
	}

	fmt.Println("Total score:", totalScore)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(2))

	timer := time.Now()
	test(input)
	fmt.Println("Time:", time.Since(timer))
}
