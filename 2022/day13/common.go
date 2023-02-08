package main

import (
	"encoding/json"
)

func unmarshal(left, right string) (any, any) {
	var l, r any

	json.Unmarshal([]byte(left), &l)
	json.Unmarshal([]byte(right), &r)

	return l, r
}

func isOrderRight(left, right string) bool {
	var l, r = unmarshal(left, right)

	if cmp(l, r) <= 0 {
		return true
	}

	return false
}

func cmp(left, right any) int {
	l, lok := left.([]any)
	r, rok := right.([]any)

	switch {
	case !lok && !rok:
		return int(left.(float64) - right.(float64))
	case !lok:
		l = []any{left}
	case !rok:
		r = []any{right}
	}

	for i := 0; i < len(l) && i < len(r); i++ {
		if res := cmp(l[i], r[i]); res != 0 {
			return res
		}
	}

	return len(l) - len(r)
}
