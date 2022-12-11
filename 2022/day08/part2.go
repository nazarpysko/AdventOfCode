package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart2(input []string) {
	highestScenicScore := 0
	for column := 1; column < len(input[0])-1; column++ {
		for row := 1; row < len(input)-1; row++ {
			t := tree{
				column: column,
				row:    row,
				height: getHeight(input, column, row),
			}
			if scenicScore := getScenicScore(t, input); scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	fmt.Println("Highest scenic score:", highestScenicScore)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(8))

	timer := time.Now()
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
