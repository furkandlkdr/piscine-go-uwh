package main

func SplitWhiteSpaces(s string) []string {
	var result []string
	word := ""
	if len(s) == 0 {
		return []string{}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		result = append(result, word)
	}

	if len(result) == 0 && len(s) > 0 {
		allSpaces := true
		for i := 0; i < len(s); i++ {
			if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
				allSpaces = false
				break
			}
		}
		if allSpaces {
			return []string{}
		}
	}

	return result
}
