package main

func isAlphanumeric(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func Capitalize(s string) string {
	result := make([]byte, len(s))
	inWord := false

	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			if !inWord {
				if s[i] >= 'a' && s[i] <= 'z' {
					result[i] = s[i] - ('a' - 'A')
				} else {
					result[i] = s[i]
				}
				inWord = true
			} else {
				if s[i] >= 'A' && s[i] <= 'Z' {
					result[i] = s[i] + ('a' - 'A')
				} else {
					result[i] = s[i]
				}
			}
		} else {
			result[i] = s[i]
			inWord = false
		}
	}
	return string(result)
}
