package main

import (
	"strings"
)

func validMarker(s string) bool {
	for _, char := range s {
		if strings.Count(s, string(char)) > 1 {
			return false
		}
	}

	return true
}

func getPositionMaker(line string, markerLength int) int {
	position := 0
	for i, _ := range line {
		if i > len(line)-markerLength {
			break
		}

		endMarker := i + markerLength
		if validMarker(line[i:endMarker]) {
			position = endMarker
			break
		}
	}

	return position
}
