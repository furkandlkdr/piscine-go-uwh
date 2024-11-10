package main

// PrintReverseAlphabets prints the uppercase alphabets from Z to A
func PrintReverseAlphabets() string {
	result := ""
	for i := 'Z'; i >= 'A'; i-- {
		result += string(i)
	}
	return result
}
