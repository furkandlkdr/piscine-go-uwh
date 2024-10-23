package main

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	// ASCII'den alinan degeler
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}
