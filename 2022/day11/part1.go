package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	barrel := newBarrel()
	for i := 0; i < len(input); i += 7 {
		barrel = append(barrel, parseMonkey(input[i+1:i+6]))
	}

	for i := 0; i < 20; i++ {
		barrel = barrel.doRound(3)
	}

	fmt.Println("Monkey business:", barrel.getMonkeyBusiness())
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(11))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
