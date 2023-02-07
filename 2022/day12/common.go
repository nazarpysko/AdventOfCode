package main

type coordinates struct {
	line int
	char int
}

func (c coordinates) addOffset(offset coordinates) coordinates {
	return coordinates{
		char: c.char + offset.char,
		line: c.line + offset.line,
	}
}

// findCoordinates returns the coordinates from the target string in the input
func findCoordinates(input []string, target string) coordinates {
	var c coordinates

	for row, rowString := range input {
		for column, char := range rowString {
			if string(char) == target {
				c = coordinates{row, column}
				break
			}
		}
	}

	return c
}

func isValidOffset(input *[]string, c coordinates) bool {
	return c.char < len((*input)[0]) && c.line < len(*input)
}

func getNeighbors(input *[]string, currentCoors coordinates) []coordinates {
	neighbors := make([]coordinates, 0)
	offsets := []coordinates{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, offset := range offsets {
		neighbor := currentCoors.addOffset(offset)
		if isValidOffset(input, neighbor) && isValidNeighbor(input, currentCoors, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func isValidNeighbor(input *[]string, currentNode, neighbor coordinates) bool {
	current := (*input)[currentNode.line][currentNode.char]
	return string(current) == "E" || current+1 == (*input)[neighbor.line][neighbor.char] || string(current) == "S" &&
}

func bfs(input []string, root coordinates, target string) int {
	queue := []coordinates{root}
	visited := make(map[coordinates]bool)

	distance := make(map[coordinates]int)
	distance[root] = 0

	resultDistance := 0

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if string(input[currentNode.line][currentNode.char]) == target {
			resultDistance = distance[currentNode]
			break
		}

		for _, neighbour := range getNeighbors(&input, currentNode) {
			if _, ok := visited[neighbour]; ok {
				continue
			}

			visited[neighbour] = true
			distance[neighbour] = distance[currentNode] + 1
			queue = append(queue, neighbour)
		}
	}

	return resultDistance
}
