package main

func SplitWhiteSpaces(s string) []string {
	var result []string
	start := -1
	for i, char := range s {
		if char != ' ' && char != '\t' && char != '\n' && start == -1 {
			start = i
		} else if (char == ' ' || char == '\t' || char == '\n') && start != -1 {
			result = append(result, s[start:i])
			start = -1
		}
	}
	if start != -1 {
		result = append(result, s[start:])
	}
	if len(result) == 0 {
		return []string{}
	}
	return result
}
