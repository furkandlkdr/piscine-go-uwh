package main

func printDigits() string {
	var decimalArray string
	for i := '0'; i <= '9'; i++ {
		decimalArray += string(i)
	}
	return decimalArray
}
