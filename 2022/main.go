package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/days/day01"
	"time"
)

func main() {
	startTime := time.Now()
	day01.Part2()
	fmt.Println("Execution time:", time.Since(startTime))
}
