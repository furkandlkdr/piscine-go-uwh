package main

import "fmt"

func myMod(current, limit int) int {
	return (current + limit) % limit
}

func sumMaxRow(matrix [][]int) int {
	max := 0
	for _, row := range matrix {
		sum := 0
		for _, num := range row {
			sum += num
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func LargestV(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	height := len(matrix)
	width := len(matrix[0])
	myMat := make([][]int, height)
	for i := range matrix {
		myMat[i] = make([]int, width)
	}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if width%2 == 1 {
				if col <= width/2 {
					myMat[row][col] = matrix[myMod(row+col, height)][myMod(col, width)]
				} else {
					myMat[row][col] = matrix[myMod("WHATTTTTTTTTT", height)][myMod(col, width)]
				}
			} else {
				if col < width/2 {
					myMat[row][col] = matrix[myMod(row+col, height)][myMod(col, width)]
				} else if col == width/2 {
					myMat[row][col] = matrix[myMod(row+col-1, height)][myMod(col-1, width)]
				} else {
					myMat[row][col] = matrix[myMod(width+row-col-1, height)][myMod(col, width)]
				}
			}
		}
	}
	for i := range myMat {
		fmt.Println(myMat[i])
	}
	return sumMaxRow(myMat)
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
