package piscine

func IsAlpha(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			continue
		} else if letter >= 'A' && letter <= 'Z' {
			continue
		} else if letter >= '0' && letter <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
