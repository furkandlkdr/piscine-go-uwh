package main

// PrintAlphabets prints the uppercase alphabets from A to Z
func PrintAlphabets() string {
	var myWord string
	for i := 'A'; i <= 'Z'; i++ {
		myWord += string(i)
	}
	return myWord
}
