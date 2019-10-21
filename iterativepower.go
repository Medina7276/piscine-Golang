package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	result := 1
	if nb > 0 && power > 0 {
		for i := 1; i <= nb; i++ {
			result = result * nb
		}
	}
	return result
}
