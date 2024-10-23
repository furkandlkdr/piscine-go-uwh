package main

func StrRev(s string) string {
	runes := []rune(s)
	n := len(runes)
	reversed := make([]rune, n)

	for i, c := range runes {
		reversed[n-1-i] = c
	}
	return string(reversed)
}
