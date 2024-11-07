package main

func hiddenp(hidden, b string) string {
    if hidden == "" {
        return "1"
    }

    j := 0
    for i := 0; i < len(b) && j < len(hidden); i++ {
        if b[i] == hidden[j] {
            j++
        }
    }
    if j == len(hidden) {
        return "1"
    }
  
	return "0"
}