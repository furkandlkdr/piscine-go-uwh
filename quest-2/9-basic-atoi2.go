package main

func containsIn0to9(ch rune) bool {
	if ch < '0' || ch > '9' {
		return false
	}
	return true
}

func strLen(s string) int {
	runeStr := []rune(s)
	return len(runeStr)
}

func BasicAtoi2(s string) int {
	if strLen(s) == 0 {
		return 0
	}
	result, startFromIndex := 0, 0
	runeStr := []rune(s)
	for i := 0; i < strLen(s) && runeStr[i] == '0'; i++ {
		startFromIndex++
	}
	for i := startFromIndex; i < strLen(s); i++ {
		if !containsIn0to9(runeStr[i]) {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}

	return result
}
