package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]   // arguments - massiv и берем с него аргументы/ то есть берем 1 и последующие аргументы
	for i := range arguments { //мы перебираем слова из массива , берем индексы
		for _, j := range arguments[i] { // i = аргумент, и перебираем его буквы . j = letter
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
	}
}
