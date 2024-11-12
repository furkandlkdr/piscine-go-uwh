package main

func ToLower(s string) string {
	runeStr := []rune(s)
	for i := 0; i < len(runeStr); i++ {
		if runeStr[i] >= 'A' && runeStr[i] <= 'Z' {
			runeStr[i] = runeStr[i] + ('a' - 'A')
		}
	}
	return string(runeStr)
}
