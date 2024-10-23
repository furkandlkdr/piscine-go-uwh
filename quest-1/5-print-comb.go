package main

func PrintComb() string {
	result := ""
	for first := '0'; first <= '9'; first++ {
		for second := first + 1; second <= '9'; second++ {
			for third := second + 1; third <= '9'; third++ {
				result += string(first) + string(second) + string(third)
				if first != '7' {
					result += ", "
				} else {
					result += "\n"
				}
			}
		}
	}

	return result
}
