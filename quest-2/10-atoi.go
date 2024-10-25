package main

// strLen returns the length of a string.
func strLen(s string) int {
	runeStr := []rune(s)
	return len(runeStr)
}

// containsIn0to9 checks if a character is a digit between '0' and '9'.
func containsIn0to9(ch rune) bool {
	if ch < '0' || ch > '9' {
		return false
	}
	return true
}

// charToInt converts a character '0' to '9' to its integer value.
func charToInt(ch rune) int {
	return int(ch - '0')
}

// Atoi converts a string to an integer with optional sign handling.
func Atoi(s string) int {
	if strLen(s) == 0 {
		return 0
	}
	sign, result, startFromIndex := 1, 0, 0
	runeStr := []rune(s)
	if runeStr[0] == '+' || runeStr[0] == '-' {
		startFromIndex++
		if runeStr[0] == '-' {
			sign = -1
		}
	}
	for i := startFromIndex; i < strLen(s) && runeStr[i] == '0'; i++ {
		startFromIndex++
	}
	for i := startFromIndex; i < strLen(s); i++ {
		if !containsIn0to9(runeStr[i]) {
			return 0
		}
		result = result*10 + charToInt(runeStr[i])
	}
	return sign * result
}
