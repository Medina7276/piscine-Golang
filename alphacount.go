package piscine

func AlphaCount(str string) int {
	count := 0
	for _, a := range str {
		if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') {
			count++
		}
	}
	return count
}
