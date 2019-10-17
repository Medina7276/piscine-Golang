package piscine

func StrRev(str string) string {
	var ans string
	for _, c := range str {
		ans = string(c) + ans
	}
	return ans
}
