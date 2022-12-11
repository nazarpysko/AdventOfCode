package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"time"
)

const (
	totalSpaceAvailable = 70000000
	spaceNeededToUpdate = 30000000
)

func solvePart2(input []string) {
	fs := fileSystem{directories: map[string]int{}}
	for _, line := range input {
		if isCommand(line) && line[2:4] == "cd" {
			fs.cd(line[5:])
		} else if isFile(line) {
			fs.add(getBytesFromLine(line))
		}
	}

	outmostDirectorySpace := fs.directories["/"]
	freeSpace := totalSpaceAvailable - outmostDirectorySpace
	spaceToBeDeleted := spaceNeededToUpdate - freeSpace
	smallestCandidate := 9223372036854775807 // Infinity as an int64
	for _, size := range fs.directories {
		if size >= spaceToBeDeleted && size < smallestCandidate {
			smallestCandidate = size
		}
	}

	fmt.Println("Size to be deleted:", smallestCandidate)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(07))

	timer := time.Now()
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
