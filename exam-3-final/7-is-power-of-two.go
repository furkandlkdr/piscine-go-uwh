package main

func getPowerOf2(nb int) int {
	result := 1
	for i := 0; i < nb; i++ {
		result *= 2
	}
	return result
}

func IsPowerOf2(nb int) bool {
	for i := 0; i <= nb/2; i++ {
		if getPowerOf2(i) == nb {
			return true
		}
	}
	return false
}
