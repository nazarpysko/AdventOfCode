package utils

import "strconv"

func GetSliceInt(s []string) []int {
	arr := make([]int, 0)
	for i := range s {
		n, _ := strconv.Atoi(s[i])
		arr = append(arr, n)
	}

	return arr
}

func Pop(s []any, quantity int) []any {
	popped := s[len(s)-quantity:]
	s = s[:len(s)-quantity]
	return popped
}
