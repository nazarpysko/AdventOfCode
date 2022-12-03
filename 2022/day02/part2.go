package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

var playToWin = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var playToLose = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var playToDraw = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func rockPaperScissorsSlave(input []string) {
	totalScore := 0
	for _, play := range input {

		switch string(play[2]) {
		case "X":
			play = play[:2] + playToLose[string(play[0])]
			break
		case "Y":
			play = play[:2] + playToDraw[string(play[0])]
			break
		case "Z":
			play = play[:2] + playToWin[string(play[0])]
		}

		totalScore += scorePlay[play]
	}

	fmt.Println("Total score:", totalScore)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(2))

	timer := time.Now()
	rockPaperScissorsSlave(input)
	fmt.Println("Time:", time.Since(timer))
}
