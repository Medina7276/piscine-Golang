package piscine

func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, letter := range s {
		if letter >= '0' && letter <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
