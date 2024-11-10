package main

func StrLen(s string) int {
	runeStr := []rune(s)
	return len(runeStr)
}
