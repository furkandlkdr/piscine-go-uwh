package main

import "fmt"

func isNumberTwoDigit(num int) bool {
	if num > 9 {
		return true
	}
	return false
}

func ConvertNumberToTwoDigitString(num int) string {
	if isNumberTwoDigit(num) {
		return fmt.Sprint(num)
	}
	return ("0" + fmt.Sprint(num))
}

func PrintComb2() string {
	result := ""

	for first := 0; first <= 99; first++ {
		for second := first + 1; second <= 99; second++ {
			result += ConvertNumberToTwoDigitString(first) + " " + ConvertNumberToTwoDigitString(second)
			if first != 98 {
				result += ", "
			}
		}
	}
	result += "\n"
	return result
}
