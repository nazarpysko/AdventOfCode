package main

import "strconv"

type tree struct {
	column int
	row    int
	height int
}

func getHeight(input []string, column, row int) int {
	height, _ := strconv.Atoi(string(input[row][column]))
	return height
}

func countVisibleTrees(input []string) int {
	visibleTrees := len(input)*2 + len(input[0])*2 - 4

	for column := 1; column < len(input[0])-1; column++ {
		for row := 1; row < len(input)-1; row++ {
			tree := tree{column: column, row: row, height: getHeight(input, column, row)}
			if isVisible(tree, input) {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func isVisible(t tree, input []string) bool {
	visibleHorizontal := true

	// Check horizontal axis right
	for x := 0; x < t.column; x++ {
		if getHeight(input, x, t.row) >= t.height {
			visibleHorizontal = false
			break
		}
	}

	if visibleHorizontal {
		return true
	}

	visibleHorizontal = true
	// Check horizontal axis left
	for x := t.column + 1; x < len(input[0]); x++ {
		if getHeight(input, x, t.row) >= t.height {
			visibleHorizontal = false
			break
		}
	}

	if visibleHorizontal {
		return true
	}

	visibleVertical := true
	// Check vertical
	for y := 0; y < t.row; y++ {
		if getHeight(input, t.column, y) >= t.height {
			visibleVertical = false
			break
		}
	}

	if visibleVertical {
		return true
	}

	visibleVertical = true
	for y := t.row + 1; y < len(input); y++ {
		if getHeight(input, t.column, y) >= t.height {
			visibleVertical = false
			break
		}
	}

	return visibleVertical
}
