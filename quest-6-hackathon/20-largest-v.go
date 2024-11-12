package main

import "fmt"

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

func main() {
	fmt.Println(LargestV([][]int{
		{0, 3, 2, 2, 3},
		{1, 0, 3, 3, 0},
		{2, 1, 0, 0, 1},
		{3, 2, 1, 1, 2},
	}))
	fmt.Println(LargestV([][]int{
		{0, 5, 0},
		{1, 0, 1},
		{2, 1, 2},
		{3, 2, 3},
		{4, 3, 4},
		{5, 4, 5},
	}))
}
