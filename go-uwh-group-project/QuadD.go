package main

func QuadD(x, y int) {
	if x < 1 || y < 1 {
		return
	}

	for i := 0; i < y; i++ {
		if i == 0 {
			WriteLine('A', 'B', 'C', x)
		} else if i == y-1 {
			WriteLine('A', 'B', 'C', x)
		} else {
			WriteLine('B', ' ', 'B', x)
		}
	}
}
