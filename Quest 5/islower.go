package piscine

func IsLower(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
