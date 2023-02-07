package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	sum := 0
	for _, snfu := range input {
		sum += SNFUtoDecimal(snfu)
	}

	fmt.Println("Total sum of SNFU is", DecimalToSNFU(sum))
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(25))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
