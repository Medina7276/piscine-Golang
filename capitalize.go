package piscine

func Capitalize(s string) string {
	s1 := []rune(s)
	k := 0
	for i := range s {
		i = i
		k++
	}
	if s1[0] >= 'a' && s1[0] <= 'z' {
		s1[0] = s1[0] - 32
	}
	for i := 1; i < k; i++ {
		if s1[i] >= 'A' && s1[i] <= 'Z' {
			if s1[i-1] >= 'A' && s1[i-1] <= 'Z' {
				s1[i] = s1[i] + 32
			}
			if s1[i-1] >= 'a' && s1[i-1] <= 'z' {
				s1[i] = s1[i] + 32
			}
			if s1[i-1] >= '0' && s1[i-1] <= '9' {
				s1[i] = s1[i] + 32
			}
		}
		if s1[i] >= 'a' && s1[i] <= 'z' {
			if s1[i-1] >= 'A' && s1[i-1] <= 'Z' {
				continue
			}
			if s1[i-1] >= 'a' && s1[i-1] <= 'z' {
				continue
			}
			if s1[i-1] >= '0' && s1[i-1] <= '9' {
				continue
			} else {
				s1[i] = s1[i] - 32
			}
		}
	}
	s2 := string(s1)
	return s2
}
