package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	for n != 0 {
		ost := n % 10

		i := '0'
		for j := 1; j <= ost; j++ {
			i++
		}
		z01.PrintRune(i)

		n = n / 10
	}
	z01.PrintRune(10)
}

// for i := '0'; i <= '7'; i++ {
// 	for j := '1'; j <= '8'; j++ {
// 		for g := '2'; g <= '9'; g++ {
// 			if i < j && j < g {
// 				z01.PrintRune(i)
// 				z01.PrintRune(j)
// 				z01.PrintRune(g)
// 				z01.PrintRune(32)
// 			}

// 		}

// 	}

// }
// z01.PrintRune(10)

// nb := -5   //создаем переменную чтобы проверить отрицательна или нет
// if nb > 0 { //если переменная больше 0
// 	z01.PrintRune('T') //то выводит Т
// } else { //иначе
// 	z01.PrintRune('F') //выводит F
// }

// 123
