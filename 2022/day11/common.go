package main

import (
	"github.com/nazarpysko/AoC/2022/utils"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type nextMonkey struct {
	divisibleBy int
	trueTarget  int
	falseTarget int
}

type monkey struct {
	items []int
	op    func(int) int
	nextMonkey
	inspectedItems int
}

func newMonkey() monkey {
	return monkey{}
}

func (m *monkey) incrementInspectedItems() {
	m.inspectedItems++
}

// A barrel is called a group of monkeys
type barrel []monkey

func newBarrel() barrel {
	return make([]monkey, 0)
}

func parseMonkey(schema []string) monkey {
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

	m.nextMonkey.divisibleBy, _ = strconv.Atoi(strings.Split(schema[2], "by ")[1])
	m.nextMonkey.trueTarget, _ = strconv.Atoi(strings.Split(schema[3], "monkey ")[1])
	m.nextMonkey.falseTarget, _ = strconv.Atoi(strings.Split(schema[4], "monkey ")[1])

	return m
}

func (b barrel) throwToMonkey(monkey, item int) {
	b[monkey].items = append(b[monkey].items, item)
}

func (b barrel) doRound(part int) barrel {
	commonMultiple := b.getCommonMultiple()
	for monkeyIndex, m := range b {
		for i := 0; i < len(m.items); i++ {
			var monkeyTarget, item int

			item, b[monkeyIndex].items = utils.Pop(b[monkeyIndex].items)
			item = int(math.Floor(float64(m.op(item))))

			if part == 1 {
				item /= 3
			} else {
				item %= commonMultiple
			}

			if item%m.divisibleBy == 0 {
				monkeyTarget = m.trueTarget
			} else {
				monkeyTarget = m.falseTarget
			}
			b[monkeyIndex].inspectedItems++
			b.throwToMonkey(monkeyTarget, item)
		}
	}

	return b
}

func (b barrel) getCommonMultiple() int {
	common_multiple := 1
	for _, m := range b {
		common_multiple *= m.divisibleBy
	}

	return common_multiple
}

func (b barrel) getMonkeyBusiness() int {
	mostActiveMonkeys := []int{0, 0}
	for _, m := range b {
		minValue, minIndex := utils.Min(mostActiveMonkeys)
		if m.inspectedItems > minValue {
			mostActiveMonkeys[minIndex] = m.inspectedItems
		}
	}

	return mostActiveMonkeys[0] * mostActiveMonkeys[1]
}
