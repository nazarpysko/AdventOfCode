package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strings"
	"time"
)

func rearrangeRucksacks(input []string) int {
	sumPriorities := 0

	for _, rucksack := range input {
		repeatedItems := ""
		firstCompartment := rucksack[:len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		for _, i := range firstCompartment {
			item := string(i)
			if strings.Contains(secondCompartment, item) && !strings.Contains(repeatedItems, item) {
				sumPriorities += getPriorityValue(i)
				repeatedItems = repeatedItems + item
			} else {
			}
		}
	}

	fmt.Println("Total sum priorities:", sumPriorities)
	return sumPriorities
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(03))

	timer := time.Now()
	rearrangeRucksacks(input)
	fmt.Println("Time:", time.Since(timer))
}
