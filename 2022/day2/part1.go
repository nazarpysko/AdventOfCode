package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solveFuncName(input []string) {

}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(2))

	timer := time.Now()
	solveFuncName(input)
	fmt.Println("Time:", time.Since(timer))
}
