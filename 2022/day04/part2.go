package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strings"
	"time"
)

func findOverlapsAtAll(first string, second string) int {
	firstSection := utils.GetSliceInt(strings.Split(first, "-"))
	secondSection := utils.GetSliceInt(strings.Split(second, "-"))

	if firstSection[0] >= secondSection[0] && firstSection[0] <= secondSection[1] ||
		secondSection[0] >= firstSection[0] && secondSection[0] <= firstSection[1] {
		return 1
	}
	return 0
}

func campCleanUpPart2(input []string) int {
	overlaps := 0
	for _, line := range input {
		sections := strings.Split(line, ",")
		overlaps += findOverlapsAtAll(sections[0], sections[1])
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
