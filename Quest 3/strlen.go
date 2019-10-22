package piscine

func StrLen(str string) int {

	stringT := []rune(str)
	i := 0
	for _, word := range stringT {

		word++
		i++
	}

	return i

}
