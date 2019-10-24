package piscine

import (
	"github.com/01-edu/z01"
)

func IntToDigits(n int) (digits []int) { // превращаю число в массив цифр
	for n > 0 { // пока число больше 0, захожу в цикл
		if n == 0 { // если число равно 0
			digits = append(digits, 0) // добавляю в массив сразу 0
		} else { // если не число не равно 0
			digits = append(digits, n%10) // добавляю в массив остаток числа от деления на 10
		}
		n /= 10 // после каждого выполненного условия делю число на 10
	}
	return digits // вывожу
}

func sortIntegerTable(table []int) []int { // сортирую в порядке возрастания цифры в массиве
	ln := 0                   // завожу переменную количества цифр в массиве
	for _, i := range table { // прохожусь циклом по массиву,
		if i == i { // и подсчитываю количество цифр с каждым шагом
			ln++
		}
	}
	// for range table { //в пределе количсивенного значения этой таблицы она крутится
	// 	ln++
	// }
	for i := 0; i < ln; i++ { // пузырьковая сортировка
		for j := 0; j < ln; j++ {
			if table[j] > table[i] {
				table[i] = table[i] + table[j]
				table[j] = table[i] - table[j]
				table[i] = table[i] - table[j]
			}
		}
	}
	return table //вывожу отсортированный массив
}

func PrintNbrInOrder(n int) { // вывожу ответ в формате руны, потому так дали по заданию
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for _, c := range sortIntegerTable(IntToDigits(n)) { //вывожу массив в формате рун
			z01.PrintRune(rune(c) + '0') // пишу (+ '0') потому что до этого мы не объявляли переменную ответа (ans := '0' -> ans = ans + 1)
		}
	}
}
