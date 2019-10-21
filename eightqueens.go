package piscine

import (
	"github.com/01-edu/z01"
)

const n = 8

var cells = [n]int{}

var ans = '0'

func RuneTransform(x int) rune {
	t := '0'
	for i := 0; i < x; i++ {
		t++
	}
	return t
}

func PositionCheck(queen, pos int) bool {
	for i := 0; i < queen; i++ { // проверяем всех queen и их позиции, затем продвигаемся
		_pos := cells[i]                                                   // получить позицию другой queen
		if _pos == pos || _pos == pos-(queen-i) || _pos == pos+(queen-i) { // rows, columns, diagonals
			return false
		}
	}
	return true
}

func SetQueens(k int) {
	if k == n {
		for i := 0; i < n; i++ { // генерируем возможности
			ans = RuneTransform(cells[i] + 1)
			z01.PrintRune(ans)
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < n; i++ {
			if PositionCheck(k, i) { // проверяем перед тем, как поставить queen в положение k
				cells[k] = i
				SetQueens(k + 1) // поставить другую queen
			}
		}
	}
}

func EightQueens() {
	SetQueens(0)
}