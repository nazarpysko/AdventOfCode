package main

import "unicode"

const columnWidth = 4
const rowLength = columnWidth * 9

// parseCrane returns a map of strings slices and the index where instructions to move crates begin
func parseCrane(input []string) (map[int][]string, int) {
	parsed := make(map[int][]string, 0)
	instructionsStart := 0

	for i, crates := range input {
		if crates == "" {
			instructionsStart = i + 1
			break
		}

		for i := 0; i < rowLength; i += columnWidth {
			if i == 0 && unicode.IsDigit(rune(crates[i+1])) {
				break
			}
			column := i/columnWidth + 1
			parsed[column] = append(parsed[column], string(crates[i+1]))
		}
	}

	return parsed, instructionsStart
}
