package piscine

func UltimateDivMod(a *int, b *int) {
	var x int
	var z int
	x = *a / *b
	z = *a % *b
	*a = x
	*b = z
}
