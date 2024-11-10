package main

func Compare(a, b string) int {
	minLen := len(a)
	if len(a) > len(b) {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] > b[i] {
			return 1
		} else if b[i] > a[i] {
			return -1
		}
	}

	if len(a) > len(b) {
		return 1
	} else if len(b) > len(a) {
		return -1
	}
	return 0
}
