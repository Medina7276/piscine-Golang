package piscine

func Join(strs []string, sep string) string {
	frazi := ""
	first := true
	for _, fraza := range strs {
		if first {
			frazi = fraza
			first = false

		} else {
			frazi += sep + fraza

		}

	}
	return frazi

}
