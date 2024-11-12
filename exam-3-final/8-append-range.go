package main

func AppendRange(min, max int) []int {
	var mySlice []int
	if min >= max {
		return mySlice
	}
	for i := min; i < max; i++ {
		mySlice = append(mySlice, i)
	}
	return mySlice
}
