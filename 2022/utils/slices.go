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

// Pop the first element of the slice
func Pop(s []int) (int, []int) {
	return s[0], s[1:]
}

// Min returns the lowest value of an array and its index
func Min(s []int) (int, int) {
	minValue := s[0]
	minIndex := 0

	for i, v := range s {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	return minValue, minIndex
}
