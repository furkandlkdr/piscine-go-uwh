package challenge

import (
	"strings"
)

func ApplyGravity(s string) string {
	if s == "" {
		return ""
	}

	lines := strings.Split(s, "\n")
	width := 0

	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	result := make([][]rune, len(lines))

	for i, line := range lines {
		result[i] = make([]rune, width)

		for j, ch := range line {
			result[i][j] = ch
		}
	}

	for i := 0; i < len(result); i++ {
		colChars := []rune{}

		for j := 0; j < len(result); j++ {
			if result[j][i] != ' ' {
				colChars = append(colChars, result[j][i])
			}
		}

		for j := 0; j < len(result); j++ {
			if j < len(colChars) {
				result[j][i] = colChars[len(colChars)-1-j]
			} else {
				result[j][i] = ' '
			}
		}
	}

	rows := []string{}

	for _, row := range result {
		strRow := string(row)
		trimmedRow := strings.TrimRight(strRow, " ")
		if trimmedRow != "" {
			rows = append(rows, trimmedRow)
		}
	}

	return strings.Join(rows, "\n")
}
