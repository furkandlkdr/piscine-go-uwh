package main

func Any(f func(string) bool, a []string) bool {
	for _, str := range a {
		if f(str) {
			return true
		}
	}
	return false
}
