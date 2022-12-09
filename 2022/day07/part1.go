package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

func solvePart1(input []string) {
	fs := fileSystem{directories: map[string]int{}}
	for _, line := range input {
		if isCommand(line) && line[2:4] == "cd" {
			fs.cd(line[5:])
		} else if isFile(line) {
			fs.add(getBytesFromLine(line))
		}
	}

	fmt.Println("Solution:", getSolution(fs.directories))
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(07))

	timer := time.Now()
	solvePart1(input)
	fmt.Println("Time:", time.Since(timer))
}
