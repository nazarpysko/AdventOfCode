package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strings"
	"time"
)

func findOverlaps(first, second string) int {
	firstSection := getArrayInt(strings.Split(first, "-"))
	secondSection := getArrayInt(strings.Split(second, "-"))

	if firstSection[0] <= secondSection[0] && firstSection[1] >= secondSection[1] ||
		secondSection[0] <= firstSection[0] && secondSection[1] >= firstSection[1] {
		return 1
	}
	return 0
}

func campCleanUp(input []string) int {
	overlaps := 0
	for _, line := range input {
		sections := strings.Split(line, ",")
		overlaps += findOverlaps(sections[0], sections[1])
	}

	fmt.Println("Overlaps found:", overlaps)
	return overlaps
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(04))

	timer := time.Now()
	campCleanUp(input)
	fmt.Println("Time:", time.Since(timer))
}
