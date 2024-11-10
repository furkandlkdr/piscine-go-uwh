package main

// PrintAlphabets prints the uppercase alphabets from A to Z
func PrintAlphabets() string {
	result := ""
	for i := 'A'; i <= 'Z'; i++ {
		result += string(i)
	}
	return result
}
