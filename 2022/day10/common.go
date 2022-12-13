package main

import (
	"regexp"
	"strconv"
	"strings"
)

type program struct {
	x      int
	cycles []int
}

func newProgram() program {
	return program{
		x:      1,
		cycles: make([]int, 0),
	}
}

func (p *program) execute(ins string) {
	noop := regexp.MustCompile("^noop")
	addx := regexp.MustCompile("^addx")

	if noop.FindString(ins) != "" {
		p.cycles = append(p.cycles, p.x)
	} else if addx.FindString(ins) != "" {
		n, _ := strconv.Atoi(strings.Split(ins, " ")[1])
		p.cycles = append(p.cycles, p.x, p.x)
		p.x += n
	}
}

func (p *program) getSolution() int {
	sum := 0
	for i := 20; i < len(p.cycles); i += 40 {
		sum += i * p.cycles[i-1]
	}

	return sum
}
