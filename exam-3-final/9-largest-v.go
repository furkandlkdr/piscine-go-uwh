package main

func rowSum(row []int) int {
	sum := 0
	for i := 0; i < len(row); i++ {
		sum += row[i]
	}
	return sum
}

func modd(n1, n2 int) int {
	return (n1 + n2 + n2) % n2
}

func LargestV(matrix [][]int) int {
	height := len(matrix)
	if height == 0 {
		return 0
	}
	width := len(matrix[0])
	myMt := make([][]int, height)

	for row := 0; row < height; row++ {
		myMt[row] = make([]int, width)
		for col := 0; col < width; col++ {
			half := width / 2
			myCol := col + row
			if width%2 == 0 {
				if col == half {
					myCol = row + col - 1
				} else if col > half {
					myCol = width - 1 - col + row
				}
			} else {
				if col > half {
					myCol = width - 1 - col + row
				}
			}
			myMt[row][col] = matrix[modd(myCol, height)][col]
		}
	}
	// calculate result
	max := rowSum(myMt[0])
	for i := 1; i < len(myMt); i++ {
		if max < rowSum(myMt[i]) {
			max = rowSum(myMt[i])
		}
	}
	return max
}
