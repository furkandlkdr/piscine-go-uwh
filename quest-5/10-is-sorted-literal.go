package main

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	// No idea what is that for or why is this working :P
	return true
}
