package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {  //
		return 1
	}

	if power < 0 { //Negative powers will return 0
		return 0
	}
	return nb * RecursivePower(nb, power-1)
}
