package main

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power < 1 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}
