package challange

import (
	"strings"
)

func ApplyGravity(s string) string {
	if s == "" {
		return ""
	}

	lines := strings.Split(s, "\n")
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	height := len(lines)
	grid := make([][]rune, height)

	for i := 0; i < height; i++ {
		grid[i] = make([]rune, maxWidth)
		line := lines[i]
		for j := 0; j < maxWidth; j++ {
			if j < len(line) {
				grid[i][j] = rune(line[j])
			} else {
				grid[i][j] = ' '
			}
		}
	}

	for col := 0; col < maxWidth; col++ {
		chars := []rune{}
		for row := 0; row < height; row++ {
			if grid[row][col] != ' ' {
				chars = append(chars, grid[row][col])
			}
		}
		row := height - 1
		for i := len(chars) - 1; i >= 0; i-- {
			grid[row][col] = chars[i]
			row--
		}
		for ; row >= 0; row-- {
			grid[row][col] = ' '
		}
	}

	resultLines := make([]string, height)
	for i := 0; i < height; i++ {
		resultLines[i] = strings.TrimRight(string(grid[i]), " ")
	}

	return strings.Join(resultLines, "\n")
}
