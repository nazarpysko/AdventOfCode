package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solveFuncName(input []string) {

}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(day))

	timer := time.Now()
	solveFuncName(input)
	fmt.Println("Time:", time.Since(timer))
}
