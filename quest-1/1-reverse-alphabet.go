package main

// PrintReverseAlphabets prints the uppercase alphabets from Z to A
func PrintReverseAlphabets() string {
	var myWord string
	for i := 'Z'; i >= 'A'; i-- {
		myWord += string(i)
	}
	return myWord
}
