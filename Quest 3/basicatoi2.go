package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, n := range s {
		y := 0
		if n < '0' || n > '9' {
			return 0
		}
		for i := '1'; i <= n; i++ {
			y++
		}
		x = x*10 + y
	}
	return x
}
