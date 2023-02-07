package main

import (
	"fmt"
	"testing"
)

func TestSNFUtoInteger(t *testing.T) {
	tests := []struct {
		SNFU    string
		decimal int
	}{
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing %s", test.SNFU), func(t *testing.T) {
			got := SNFUtoDecimal(test.SNFU)
			if got != test.decimal {
				t.Errorf("SNFU %s is not %d. Should be %d", test.SNFU, got, test.decimal)
			}
		})
	}
}

func TestDecimalToSNFU(t *testing.T) {
	tests := []struct {
		decimal int
		SNFU    string
	}{
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{7, "12"},
		{8, "2="},
		{9, "2-"},
		{10, "20"},
		{15, "1=0"},
		{20, "1-0"},
		{2022, "1=11-2"},
		{12345, "1-0---0"},
		{314159265, "1121-1110-1=0"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing %d", test.decimal), func(t *testing.T) {
			got := DecimalToSNFU(test.decimal)
			if got != test.SNFU {
				t.Errorf("%d is not SNFU %s. Should be %s", test.decimal, got, test.SNFU)
			}
		})
	}
}
