package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strconv"
	"strings"
	"time"
)

func getArrayInt(s []string) []int {
	arr := make([]int, 0)
	for i := range s {
		n, _ := strconv.Atoi(s[i])
		arr = append(arr, n)
	}

	return arr
}

func findOverlaps(first, second string) int {
	firstSection := getArrayInt(strings.Split(first, "-"))
	secondSection := getArrayInt(strings.Split(second, "-"))

	fmt.Printf("%s,%s: %s <= %s && %s >= %s ---> %t\n",
		first, second,
		firstSection[0], secondSection[0], firstSection[1], secondSection[1],
		firstSection[0] <= secondSection[0] && firstSection[1] >= secondSection[1])

	if firstSection[0] <= secondSection[0] && firstSection[1] >= secondSection[1] {
		return 1
	}
	return 0
}

func campCleanUp(input []string) int {
	overlaps := 0
	for _, line := range input {
		sections := strings.Split(line, ",")

		if findOverlaps(sections[0], sections[1]) == 1 {
			overlaps++
		} else if findOverlaps(sections[1], sections[0]) == 1 {
			overlaps++
		}
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
