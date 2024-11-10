package main

func PrintComb() string {
	result := ""
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				result += string(i) + string(j) + string(k)
				if i == '7' {
					result += "\n"
				} else {
					result += ", "
				}
			}
		}
	}

	return result
}
