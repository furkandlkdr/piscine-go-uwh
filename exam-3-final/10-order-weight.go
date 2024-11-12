package main

func strCmp(s1, s2 string) int {
	minLen := len(s1)
	if len(s1) > len(s2) {
		minLen = len(s2)
	}
	for i := 0; i < minLen; i++ {
		if s1[i] > s2[i] {
			return 1
		} else if s1[i] < s2[i] {
			return -1
		}
	}
	if len(s1) > len(s2) {
		return 1
	} else if len(s2) > len(s1) {
		return -1
	}
	return 0
}

func calculateWeight(numb string) int {
	sum := 0
	for i := 0; i < len(numb); i++ {
		sum += int(numb[i] - '0')
	}
	return sum
}

func parseString(str string) []string {
	startFrom := 0
	var strs []string
	i := 0
	for ; i < len(str); i++ {
		if str[i] == ' ' {
			strs = append(strs, str[startFrom:i])
			startFrom = i + 1
		}
	}
	strs = append(strs, str[startFrom:i])
	return strs
}

func OrderWeight(str string) string {
	result := ""
	var lenghts []int
	strs := parseString(str)
	for i := 0; i < len(strs); i++ {
		lenghts = append(lenghts, calculateWeight(strs[i]))
	}
	for i := 0; i < len(lenghts)-1; i++ {
		for j := i + 1; j < len(lenghts); j++ {
			if (lenghts[i] > lenghts[j]) ||
				(lenghts[i] == lenghts[j] && strCmp(strs[i], strs[j]) == 1) {
				lenghts[i], lenghts[j] = lenghts[j], lenghts[i]

				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}
	for i := 0; i < len(strs); i++ {
		result += strs[i]
		if i != len(strs)-1 {
			result += " "
		}
	}

	return result
}
