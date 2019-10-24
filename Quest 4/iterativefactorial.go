package piscine

func IterativeFactorial(nb int) int {
	if nb > 65535 {
		return 0
	}
	if nb < 0 {
		return (0)
	}
	if nb == 0 || nb == 1 { // !0 = 1 (всегда) математическое правило - нужно запомнить
		return (1)
	}

	result := 1
	for i := 1; i <= nb; i++ { // условия для цикла
		result = result * i // это есть сам цикл 
	}
	return result
}
