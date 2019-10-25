package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < 2147483647; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
/* FINDNEXTPRIME.GO
Write a function that returns the first prime number 
that is equal or superior to the int passed as parameter.
The function must be optimized in order 
to avoid time-outs with the tester.

student@ubuntu:~/piscine-go/test$ ./test
5
5
*
