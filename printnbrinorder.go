package piscine

import (
	"github.com/01-edu/z01"
)

func IntToDigits(n int) (digits []int) {
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n /= 10
	}
	return digits
}

func sortIntegerTable(table []int) []int {
	ln := 0
	for _, i := range table {
		if i == i {
			ln++
		}
	}
	for i := 0; i < ln; i++ {
		for j := 0; j < ln; j++ {
			if table[j] > table[i] {
				table[i] = table[i] + table[j]
				table[j] = table[i] - table[j]
				table[i] = table[i] - table[j]
			}
		}
	}
	return table
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for _, c := range sortIntegerTable(IntToDigits(n)) {
			z01.PrintRune(rune(c) + '0')
		}
	}
}
