package main

func getPriorityValue(c int32) int {
	if c >= 65 && c <= 90 {
		return int(c - 38)
	} else {
		return int(c - 96)
	}
}
