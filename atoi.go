package piscine

func Atoi(s string) int {
	x := 0
	k := 1
	s1 := []rune(s)

	if s != "" {
		for i, n := range s1 {
			y := 0
			if n < '0' || n > '9' {
				if n == '-' || n == '+' {
					if i != 0 {
						return 0
					}
					if n == '-' {
						k = -1
					}
					if n == '+' {
						k = 1
					}
				} else {
					return 0
				}
			}
			for i := '1'; i <= n; i++ {
				y++
			}
			x = x*10 + y
		}
	}
	return x * k
}
