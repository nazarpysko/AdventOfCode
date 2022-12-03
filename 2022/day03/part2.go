package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strings"
	"time"
)

func getBadge(lines []string) int32 {
	var badge int32
	for _, i := range lines[0] {
		item := string(i)
		if strings.Contains(lines[1], item) && strings.Contains(lines[2], item) {
			badge = i
			break
		}
	}

	return badge
}

func rearrangeRucksacksPart2(input []string) {
	prioritySum := 0
	for i := 0; i < len(input); i += 3 {
		badge := getBadge(input[i : i+3])
		prioritySum += getPriorityValue(badge)
	}

	fmt.Println("Priority sum:", prioritySum)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(3))

	timer := time.Now()
	rearrangeRucksacksPart2(input)
	fmt.Println("Time:", time.Since(timer))
}
