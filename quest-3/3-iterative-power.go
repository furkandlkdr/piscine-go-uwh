package main

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}