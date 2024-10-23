package main

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, str := range tab {
		if f(str) {
			count++
		}
	}
	return count
}
