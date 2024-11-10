package main

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func Capitalize(s string) string {
	result := ""
	currentlyInWord := false

	for i := 0; i < len(s); i++ {
		if isAlphaNum(s[i]) {
			if !currentlyInWord {
				if s[i] >= 'a' && s[i] <= 'z' {
					result += string(s[i] - ('a' - 'A'))
				} else {
					result += string(s[i])
				}
				currentlyInWord = true
			} else {
				if s[i] >= 'A' && s[i] <= 'Z' {
					result += string(s[i] + ('a' - 'A'))
				} else {
					result += string(s[i])
				}
			}

		} else {
			result += string(s[i])
			currentlyInWord = false
		}
	}
	return result
}
