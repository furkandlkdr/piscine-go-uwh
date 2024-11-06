package main

func powerOfTwo(nb int) int {
	flag := 1
	for i := 1; i < nb; i++ {
		flag *= 2
		if flag == nb {
			return i
		}
	}
	return flag
}

func PackArray(integers []int) int {
	if len(integers) == 1 {
		return integers[0]
	}
	k := powerOfTwo(len(integers))
	newIntegers := integers
	var iterationHolder []int

	for it := 1; it <= k; it++ {
		iterationHolder = []int{}
		arrayLen := len(newIntegers)

		if it%2 == 1 {
			for i := 0; i < arrayLen; i += 2 {
				iterationHolder = append(iterationHolder, newIntegers[i]+newIntegers[i+1])
			}
			newIntegers = iterationHolder
		} else {
			for i := 0; i < arrayLen; i += 2 {
				iterationHolder = append(iterationHolder, newIntegers[i]*newIntegers[i+1])
			}
			newIntegers = iterationHolder
		}
	}

	return newIntegers[len(newIntegers)-1]
}
