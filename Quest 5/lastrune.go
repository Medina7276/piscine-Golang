package piscine

func LastRune(s string) rune {
	s1 := []rune(s)
	k := 0
	for _, i := range s1 {
		if i == i {
		}
		k++
	}
	return s1[k-1]
}

// func LastRune(s string) rune {
// 	r := []rune(s)
// 	last := 0
// 	for i := range r {
// 		last = i
// 	}
// 	return r[last]
// }
