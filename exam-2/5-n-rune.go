package main

func NRune(s string, n int) rune {
	runeStr := []rune(s)
	if n > len(runeStr) || n < 1 {
		return 0
	}
	return runeStr[n-1]
}
