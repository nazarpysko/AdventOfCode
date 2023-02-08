package main

import "testing"

func TestIsOrderRight(t *testing.T) {
	tests := []struct {
		name  string
		left  string
		right string
		valid bool
	}{
		{"Left side is smaller ",
			"[1,1,3,1,1]", "[1,1,5,1,1]", true},
		{"Mixed types and left is smaller",
			"[[1],[2,3,4]]", "[[1],4]", true},
		{"Mixed types and right is smaller",
			"[9]", "[[8,7,6]]", false},
		{"Left side ran out of items",
			"[[4,4],4,4]", "[[4,4],4,4,4]", true},
		{"Right side ran out of items",
			"[7,7,7,7]", "[7,7,7]", false},
		{"Left side ran out of items",
			"[]", "[3]", true},
		{"Right side ran out of items",
			"[[[]]]", "[[]]", false},
		{"Right side is smaller",
			"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isOrderRight(test.left, test.right)
			if got != test.valid {
				t.Errorf("Got %t and should be %t", got, test.valid)
			}
		})
	}
}
