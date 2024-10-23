package main

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	current := ""

	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, current)
			current = ""
			i += sepLen - 1
		} else {
			current += string(s[i])
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
