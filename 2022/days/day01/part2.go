package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func Part2() {
	f, err := os.Open("C:\\Users\\usuario\\Desktop\\AoC\\2022\\day01\\input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	maxCalories := [3]int{0, 0, 0}
	currentElfCalories := 0
	for sc.Scan() {
		if sc.Text() == "" {
			i := min(maxCalories)
			if currentElfCalories > maxCalories[i] {
				maxCalories[i] = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			calories, _ := strconv.Atoi(sc.Text())
			currentElfCalories += calories
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	fmt.Println("Max calories:", sum(maxCalories))
}
