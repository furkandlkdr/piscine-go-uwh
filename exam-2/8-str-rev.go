package main

func StrRev(s string) string {
	runeStr := []rune(s)
	myLen := len(runeStr)
	result := make([]rune, myLen)
	for i := 0; i < myLen; i++ {
		result[i] = runeStr[myLen-1-i]
	}
	return string(result)
}
