package main

func isOrderRight(left, right string) bool {
	for i, _ := range left {
		if left[i] > right[i] {
			return false
		}
	}

	return true
}