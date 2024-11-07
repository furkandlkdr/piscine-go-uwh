package main

func jumpDigits(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		if ch > 'l' {
			return ch - 12
		}
		return ch + 14
	} else if ch >= 'A' && ch <= 'Z' {
		if ch > 'L' {
			return ch - 12
		}
		return ch + 14
	}
	return ch
}

func Rot14(s string) string {
	var rot14 string
	runeStr := []rune(s)
	for i := 0; i < len(runeStr); i++ {
		rot14 += string(jumpDigits(runeStr[i]))
	}
	return rot14
}
