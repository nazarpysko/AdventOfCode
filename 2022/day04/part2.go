package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strings"
	"time"
)

func findOverlapsAtAll(first string, second string) bool {
	firstSection := getArrayInt(strings.Split(first, "-"))
	secondSection := getArrayInt(strings.Split(second, "-"))

	if firstSection[0] <= secondSection[0] || firstSection[1] >= secondSection[1] {
		return true
	}
	return false
}

func campCleanUpPart2(input []string) int {
	overlaps := 0
	for _, line := range input {
		sections := strings.Split(line, ",")
		if findOverlapsAtAll(sections[0], sections[1]) {
			overlaps++
		}
	}

	fmt.Println("Overlaps found:", overlaps)
	return overlaps
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(04))

	timer := time.Now()
	campCleanUpPart2(input)
	fmt.Println("Time:", time.Since(timer))
}
