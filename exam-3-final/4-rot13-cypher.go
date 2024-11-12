package main

func jump13digit(c rune) rune {
	if c >= 'a' && c <= 'z' {
		if c > 'n' {
			return c - 13
		}
		return c + 13
	} else if c >= 'A' && c <= 'Z' {
		if c > 'N' {
			return c - 13
		}
		return c + 13
	}
	return c
}

func rot13(s string) string {
	runeStr := []rune(s)
	for i := 0; i < len(runeStr); i++ {
		runeStr[i] = jump13digit(runeStr[i])
	}
	return string(runeStr)
}
