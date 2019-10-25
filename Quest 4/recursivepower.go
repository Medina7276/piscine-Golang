package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {  //
		return 1
	}

	if power < 0 { //Negative powers will return 0
		return 0
	}
	return nb * RecursivePower(nb, power-1)
}
/*RECURSIVEPOWER.GO
Write an recursive function that returns the power of the int passed as parameter.
Negative powers will return 0. Overflows do not have to be dealt with.
for is forbidden for this exercise.
student@ubuntu:~/piscine-go/test$ ./test
64
*
