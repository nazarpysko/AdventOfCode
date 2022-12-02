package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strconv"
	"time"
)

func part1(input []string) {
	maxCalories := 0
	currentElfCalories := 0
	for _, line := range input {
		if line == "" {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentElfCalories += calories
		}
	}

	fmt.Println("Max calories:", maxCalories)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(1))

	timer := time.Now()
	part1(input)
	fmt.Println("Time:", time.Since(timer))
}
