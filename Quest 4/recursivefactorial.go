package piscine

func RecursiveFactorial(nb int) int {

	if nb < 0 || nb > 12 {  
		return 0  
	} else if nb == 0 {
		return 1
	}

	if nb == 1 { 
		return 1
	}
	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}

	return 0 //возвращается в начало (это не ознначает что он возвращает ноль)
}
/*RECURSIVEFACTORIAL.GO
Write a recursive function that returns the factorial of the int passed as parameter.
Errors (non possible values or overflows) will return 0.
for is forbidden for this exercise.
student@ubuntu:~/piscine-go/test$ ./test
24
*
