package main

func AppendRange(min, max int) []int {
	if min >= max {
		var s []int
		return s
	}
	myLen := max - min
	var result []int
	// result := make([]int, myLen) // if you want to use make
	for i := 0; i < myLen; i++ {
		// result[i] = min // if you want to use make
		result = append(result, min)
		min++
	}
	return result
}
