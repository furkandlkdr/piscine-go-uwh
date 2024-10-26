package main

import "fmt"

func WriteLine(lc rune, mc rune, rc rune, lenght int) {
	fmt.Printf("%c", lc)

	for i := 0; i < lenght-2; i++ {
		fmt.Printf("%c", mc)
	}
	if lenght > 1 {
		fmt.Printf("%c\n", rc)
	} else {
		fmt.Printf("\n")
	}
}
