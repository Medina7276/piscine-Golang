package main

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for p := 0; p < y; p++ {
		if p == 0 || p == y-1 {
			for n := 0; n < x; n++ {
				if n == 0 || n == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune(10)

		} else {
			for p := 0; p < x; p++ {
				if p == 0 || p == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune(10)
		}

	}
}

func main() {
	Raid1a(5,3)
}
