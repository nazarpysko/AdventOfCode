package main

import (
	"fmt"
	"github.com/nazarpysko/AoC/2022/utils"
	"math"
)

const baseSNFU = 5

func SNFUDigitToDecimal(r rune) int {
	integer := 0
	if r == '=' {
		integer = -2
	} else if r == '-' {
		integer = -1
	} else {
		integer = int(r - '0')
	}

	return integer
}

func SNFUtoDecimal(snfu string) int {
	snfu = utils.Reverse(snfu)
	decimal := 0
	for i, c := range snfu {
		decimal += int(math.Pow(float64(baseSNFU), float64(i))) * SNFUDigitToDecimal(c)
	}

	return decimal
}

func DecimalToSNFU(decimal int) string {
	if decimal == 0 {
		return ""
	}

	remainder := decimal % baseSNFU
	decimal /= baseSNFU

	switch remainder {
	case 0, 1, 2:
		return DecimalToSNFU(decimal) + fmt.Sprintf("%d", remainder)
	case 3:
		return DecimalToSNFU(1+decimal) + "="
	case 4:
		return DecimalToSNFU(1+decimal) + "-"
	}

	return ""
}
