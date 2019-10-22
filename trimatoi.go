package piscine

func IntCheck(s string) bool {
	ok := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			ok = true
			break
		}
	}
	return ok
}

func MinusCheck(s string) bool {
	minus := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			break
		}
		if c == '-' {
			minus = true
		}
	}
	return minus
}

func TrimAtoi(s string) int {
	x := 0
	if IntCheck(s) == true {
		for _, c := range s {
			if c >= '0' && c <= '9' {
				k := 0
				for i := '1'; i <= c; i++ {
					k++
				}
				x = x*10 + k
			}
		}
	}
	if MinusCheck(s) {
		x *= -1
	}
	return x
}
