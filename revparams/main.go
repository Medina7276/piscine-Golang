package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ArrLen(Args []string) int {
	var count int
	for i := range Args {
		count = i + 1
	}
	return count
}

func main() {
	args := os.Args[1:]

	for i := range args {
		for _, j := range os.Args[ArrLen(args)-i] { //cause Args [] starts with 0
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
	}
}

/*REVPARAMS/main.go
Write a program that prints the arguments received in the command line in a reverse order.
Example of output :
go build
./revparams choumi is the best cat

cat
best
the
is
choumi*

