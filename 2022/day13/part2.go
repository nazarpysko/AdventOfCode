package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"sort"
	"time"
)

func solvePart2(input []string) {
	key := 1
	var packets []any

	for i := 0; i < len(input); i += 3 {
		left, right := input[i], input[i+1]
		l, r := unmarshal(left, right)
		packets = append(packets, l, r)
	}

	packets = append(packets, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})

	for i, p := range packets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			key *= i + 1
		}
	}

	fmt.Println("The decoder key is", key)
}

func main() {
	input := utils.ReadInputFile(utils.GenInputFile(13))

	timer := time.Now()
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}
