package main

func ConcatParams(args []string) string {
	result := ""
	for i := 0; i < len(args); i++ {
		if i > 0 {
			result += "\n"
		}
		result += args[i]
	}
	return result
}
