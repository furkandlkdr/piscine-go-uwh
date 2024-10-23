package main

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0)
	//If you use make instead of var, you can bypass the nil check
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
	return result
}
