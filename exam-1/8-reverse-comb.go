package main

func ReverseComb() string {
	result := ""
	for i := '9'; i >= '2'; i-- {
		for j := i - 1; j >= '1'; j-- {
			for k := j - 1; k >= '0'; k-- {
				result += string(i) + string(j) + string(k)
				if i == '2' {
					result += "\n"
				} else {
					result += ", "
				}
			}
		}
	}
	return result
}
