package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	stringT := []rune(str)

	for _, word := range stringT {

		z01.PrintRune(rune(word))

	}

}
