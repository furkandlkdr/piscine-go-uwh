package main

func jump13digit(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		if ch >= 'n' {
			return ch - 13
		}
		return ch + 13
	} else if ch >= 'A' && ch <= 'Z' {
		if ch >= 'N' {
			return ch - 13
		}
		return ch + 13
	}
	return ch
}

func rot13(s string) string {
	result := ""
	runeStr := []rune(s)
	for i := 0; i < len(runeStr); i++ {
		result += string(jump13digit(runeStr[i]))
	}
	return result
}
