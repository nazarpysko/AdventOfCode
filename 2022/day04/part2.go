package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(04))

	timer := time.Now()
	campCleanUp(input)
	fmt.Println("Time:", time.Since(timer))
}
