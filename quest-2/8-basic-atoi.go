package main

func strLen(s string) int {
	runeStr := []rune(s)
	return len(runeStr)
}

func BasicAtoi(s string) int {
	if strLen(s) == 0 {
		return 0
	}
	result, startFromIndex := 0, 0
	runeStr := []rune(s)
	for i := 0; i < strLen(s) && runeStr[i] == '0'; i++ {
		startFromIndex++
	}
	for i := startFromIndex; i < strLen(s); i++ {
		result = result*10 + int(runeStr[i]-'0')
	}

	return result
}
