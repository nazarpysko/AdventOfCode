package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"strings"
	"time"
)

func rearrangeRucksacks(input []string) {
	sumPriorities := 0

	for _, rucksack := range input {
		repeatedItems := ""
		firstCompartment := rucksack[:len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		for _, i := range firstCompartment {
			item := string(i)
			fmt.Printf("Searching %c in %s\n", i, secondCompartment)
			if strings.Contains(secondCompartment, item) && !strings.Contains(repeatedItems, item) {
				sumPriorities += getPriorityValue(i)
				repeatedItems = repeatedItems + item
				fmt.Println("Found. Item priority value is", getPriorityValue(i))
			} else {
				fmt.Println("Not found")
			}
		}
	}

	fmt.Println("Total sum priorities:", sumPriorities)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(03))
	//input := []string{
	//	"vJrwpWtwJgWrhcsFMMfFFhFp",
	//	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	//	"PmmdzqPrVvPwwTWBwg",
	//	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	//	"ttgJtRGJQctTZtZT",
	//	"CrZsJsPPZsGzwwsLwLmpwMDw",
	//}
	timer := time.Now()
	rearrangeRucksacks(input)
	fmt.Println("Time:", time.Since(timer))
}
