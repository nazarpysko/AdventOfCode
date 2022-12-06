package main

import (
	"regexp"
	"strconv"
	"unicode"
)

const columns = 9
const columnWidth = 4
const rowLength = columnWidth * columns

type instruction struct {
	quantity int
	from     int
	to       int
}

// Indexed from 0 to columns
type crates map[int][]string

func (c crates) pop(from, quantity int) []string {
	popped := make([]string, 0)
	popped = append(popped, c[from][:quantity]...)
	c[from] = c[from][quantity:]

	return popped
}

func (c crates) popReversed(from, quantity int) []string {
	popped := c.pop(from, quantity)
	reversed := make([]string, 0)

	for _, item := range popped {
		reversed = append([]string{item}, reversed...)
	}

	return reversed
}

func (c crates) moveCrates(inst instruction, pop func(from, quantity int) []string) {
	popped := pop(inst.from, inst.quantity)
	c[inst.to] = append(popped, c[inst.to]...)
}

func (c crates) getTopCrates() string {
	topCrates := ""
	for i := 0; i < columns; i++ {
		topCrates += c[i][0]
	}

	return topCrates
}

// parseCrane returns a map of strings slices and the index where instructions to move crates begin
func parseCrane(input []string) (crates, int) {
	parsed := make(crates, 0)
	instructionsStart := 0

	for i, cratesInLine := range input {
		if cratesInLine == "" {
			instructionsStart = i + 1
			break
		}

		for i := 0; i < rowLength; i += columnWidth {
			characterIndex := i + 1
			if i == 0 && unicode.IsDigit(rune(cratesInLine[characterIndex])) {
				break
			}

			if string(cratesInLine[characterIndex]) == " " {
				continue
			}

			column := i / columnWidth
			parsed[column] = append(parsed[column], string(cratesInLine[characterIndex]))
		}
	}

	return parsed, instructionsStart
}

func parseInstruction(inst string) instruction {
	r := regexp.MustCompile(`move (\d+) from (\d) to (\d)`)
	matches := r.FindStringSubmatch(inst)

	quantity, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	return instruction{
		quantity: quantity,
		from:     from - 1, // Off by one to address correct columns in crates
		to:       to - 1,
	}
}
