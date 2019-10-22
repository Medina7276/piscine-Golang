package main

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		 for b := '1'; b <= '8'; b++ {
			 for c := '2'; c <= '9'; c++ {
				 if a < b && b < c {
					     if a == '7' && b ==''8 && c == '9'{
								 z01.Printrune(a)
								 z01.Printrune(b)
								 z01.Printrune(c)
						  }else{
								 z01.Printrune(a)
								 z01.Printrune(b)
								 z01.Printrune(c)
								 z01.Printrune(',')
								 z01.Printrune(' ')
						 }
				 }
		 } 	
	}
	z01.Printrune('\n') 
}
func main() {
	PrintComb
}
