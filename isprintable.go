package piscine

func IsPrintable(fraza string) bool {
	strbool := true
	strRune := []rune(fraza)
	for _, key := range strRune {
		if key >= 0 && key <= 31 {
			strbool = false
		}
	}
	return strbool
}
