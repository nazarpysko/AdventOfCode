package utils

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Sign(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		return 1
	} else {
		return -1
	}
}
