package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	for index, char := range s1 {
		if index == n-1 {
			return char
		}
	}
	return 0
}
