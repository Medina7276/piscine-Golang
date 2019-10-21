package piscine

func IterativeFactorial(nb int) int {

	if nb < 0 || nb > 12 {
		return 0
	} else if nb == 0 {
		return 1
	}

	c := nb
	b := 1
	for nb := 1; nb <= c; nb++ {

		b = b * nb
	}
	return b

}
