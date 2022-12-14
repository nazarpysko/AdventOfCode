package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	monkeys := newBarrel()
	for i := 0; i < len(input); i += 8 {
		monkeys.addMonkey(input[i+1 : i+6])
	}

	for i := 0; i < 20; i++ {
		monkeys.doRound()
	}

	fmt.Println("Monkey business:", monkeys.getMonkeyBusiness())
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(11))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
