package main

func Compact(ptr *[]string) int {
	var result []string
	myStr := *ptr
	counter := 0
	for i := 0; i < len(myStr); i++ {
		if !(myStr[i] == "" || myStr[i] == "0" || myStr[i] == "false") {
			result = append(result, myStr[i])
			counter += 1
		}
	}
	if len(myStr) == 0 || counter == 0 {
		myNil := make([]string, 0)
		*ptr = myNil
		return 0
	}
	*ptr = result

	return counter
}
