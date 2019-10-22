package piscine

func IsUpper(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
