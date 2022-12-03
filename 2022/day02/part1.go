package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

const (
	win      = 6
	draw     = 3
	lose     = 0
	rock     = 1
	paper    = 2
	scissors = 3
)

var scorePlay = map[string]int{
	"A X": draw + rock,
	"A Y": win + paper,
	"A Z": lose + scissors,
	"B X": lose + rock,
	"B Y": draw + paper,
	"B Z": win + scissors,
	"C X": win + rock,
	"C Y": lose + paper,
	"C Z": draw + scissors,
}

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
