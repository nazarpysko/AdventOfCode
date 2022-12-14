package main

import (
	"github.com/nazarpysko/AoC/2022/utils"
	"strconv"
	"strings"
	"unicode"
)

type nextMonkeyFunc struct {
	divisibleBy int
	trueTarget  int
	falseTarget int
}

type monkey struct {
	items []int
	op    func(int) int
	nextMonkeyFunc
}

func newMonkey() monkey {
	return monkey{}
}

// A barrel is called a group of monkeys
type barrel struct {
	monkeys []monkey
}

func newBarrel() barrel {
	return barrel{
		monkeys: make([]monkey, 0),
	}
}

func (b *barrel) addMonkey(schema []string) {
	m := newMonkey()
	m.items = utils.GetSliceInt(strings.Split(strings.Split(schema[0], ": ")[1], ", "))
	m.op = func(old int) int {
		new := old
		var operand int

		aux := strings.Split(schema[1], "new = old ")[1]
		op := aux[0]
		operandString := aux[2:]

		if unicode.IsDigit(rune(operandString[0])) {
			operand, _ = strconv.Atoi(aux[2:])
		} else {
			operand = old
		}

		switch string(op) {
		case "+":
			new += operand
		case "-":
			new -= operand
		case "*":
			new *= operand
		case "/":
			new /= operand
		}

		return new
	}
}
