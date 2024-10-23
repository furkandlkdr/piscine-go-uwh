package main

func Index(s string, toFind string) int {
	strLen := len(s)
	findLen := len(toFind)
	for i := 0; i <= strLen-findLen; i++ {
		if s[i:i+findLen] == toFind {
			return i
		}
	}
	return -1
}
