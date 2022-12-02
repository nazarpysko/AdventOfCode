package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strconv"
	"time"
)

// min returns index of the lowest value in an array
func min(arr [3]int) int {
	minValue := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}
	return minIndex
}

func sum(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func part2(input []string) {
	maxCalories := [3]int{0, 0, 0}
	currentElfCalories := 0
	for _, line := range input {
		if line == "" {
			i := min(maxCalories)
			if currentElfCalories > maxCalories[i] {
				maxCalories[i] = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentElfCalories += calories
		}
	}

	fmt.Println("Max calories:", sum(maxCalories))
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(1))

	timer := time.Now()
	part2(input)
	fmt.Println("Time:", time.Since(timer))
}
