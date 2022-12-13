package main

import (
	"regexp"
	"strconv"
	"strings"
)

type CRP struct {
	spriteCentralPosition int
	drawing               string
}

func newCRP() CRP {
	return CRP{
		spriteCentralPosition: 1,
	}
}

func (c *CRP) render(cycle, x int) {
	cycle--
	cycle %= 40

	if c.spriteCentralPosition-1 <= cycle && cycle <= c.spriteCentralPosition+1 {
		c.drawing += "#"
	} else {
		c.drawing += " "
	}

	if c.spriteCentralPosition != x+1 {
		c.spriteCentralPosition = x + 1
	}

	if cycle == 0 {
		c.drawing += "\n"
	}
}

type program struct {
	x      int
	cycles []int
	crp    CRP
}

func newProgram() program {
	return program{
		x:      1,
		cycles: make([]int, 0),
		crp:    newCRP(),
	}
}

func (p *program) execute(ins string) {
	noop := regexp.MustCompile("^noop")
	addx := regexp.MustCompile("^addx")

	if noop.FindString(ins) != "" {
		p.cycles = append(p.cycles, p.x)
		p.crp.render(len(p.cycles), p.x)
	} else if addx.FindString(ins) != "" {
		n, _ := strconv.Atoi(strings.Split(ins, " ")[1])
		p.cycles = append(p.cycles, p.x, p.x)
		p.crp.render(len(p.cycles)-1, p.x)
		p.crp.render(len(p.cycles), p.x)
		p.x += n
	}
}

func (p *program) getSumSignals() int {
	sum := 0
	for i := 20; i < len(p.cycles); i += 40 {
		sum += i * p.cycles[i-1] // Off By One because first cycle is at index 0 not 1. So i = 20 should point at 19
	}

	return sum
}
