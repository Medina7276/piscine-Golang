package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	len := 0
	args := os.Args
	for range args {
		len++
	}
	revargs := args
	for i := 1; i < len; i++ {
		revargs[i-1] = args[i]
	}
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		args[i], args[j] = args[j], args[i]
	}
	for i := 1; i < len; i++ {
		arg := revargs[i]
		for _, letter := range arg {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}

}

//klk
