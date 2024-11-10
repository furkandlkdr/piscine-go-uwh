package main

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb < 2 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
