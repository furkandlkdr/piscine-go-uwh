package main

import "fmt"

func CountDigits(nb int) int {
	counter := 0
	for nb > 0 {
		nb = nb / 10
		counter++
	}
	return counter
}

func myPower(nb, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	result := ""
	if n < 0 {
		sign = "-"
		n = n * -1
	}
	myLen := CountDigits(n)
	for i := 0; i < myLen; i++ {
		result += fmt.Sprint(n / myPower(10, CountDigits(n)-1))
		n = n % myPower(10, CountDigits(n)-1)
	}

	return sign + result
}
