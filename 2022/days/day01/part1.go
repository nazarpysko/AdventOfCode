package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Part1() {
	f, err := os.Open("C:\\Users\\usuario\\Desktop\\AoC\\2022\\day01\\input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	maxCalories := 0
	currentElfCalories := 0
	for sc.Scan() {
		if sc.Text() == "" {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
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

	fmt.Println("Max calories:", maxCalories)
}
