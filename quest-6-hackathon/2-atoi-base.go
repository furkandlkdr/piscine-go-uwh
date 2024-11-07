package main

func isBaseValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			return false
		}
	}
	return true
}

func isBaseContainsNumber(nb, base string) bool {
	myFlag := false
	for i := 0; i < len(nb); i++ {
		myFlag = false
		for j := 0; j < len(base); j++ {
			if nb[i] == base[j] {
				myFlag = true
				break
			}
		}
		if myFlag == false {
			return false
		}
	}
	return true
}

func convertDigitToInt(digit byte, base string) int {
	for i := 0; i < len(base); i++ {
		if digit == base[i] {
			return i
		}
	}
	return 0
}

func AtoiBase(s string, t string) int {
	if !(isBaseContainsNumber(s, t) && isBaseValid(t)) {
		return 0
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = (sum * len(t)) + convertDigitToInt(s[i], t)
	}
	return sum

}
