package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		left, right := input[i], input[i+1]

		if isOrderRight(left, right) {
			sum += i/3 + 1
		}
	}

	fmt.Println("Sum of indices:", sum)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(13))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
