package piscine

func AlphaCount(str string) int {
	count := 0              // количесвто букв
	for _, a := range str { // пробегаюсь не по индексам, а рунам на этих индексах
		if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') { // проверяю, если рун - буква
			count++ // елси буква, то увеличиваю количество
		}
	}
	return count // вывожу количество
}
