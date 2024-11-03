package main

func QuadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}

	for i := 0; i < y; i++ {
		if i == 0 {
			WriteLine('o', '-', 'o', x)
		} else if i == y-1 {
			WriteLine('o', '-', 'o', x)
		} else {
			WriteLine('|', ' ', '|', x)
		}
	}
}
