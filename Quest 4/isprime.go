package piscine

func IsPrime(nb int) bool {
	if nb < 0 {
		return false
	}

	if nb <= 3 {
		return nb > 1
	}

	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i = i + 6
	}

	return true
}
/* ISPRIME.GO
Write a function that returns true if the int passed as parameter is a prime number. 
Otherwise it returns false.
The function must be optimized in order to avoid time-outs with the tester.
student@ubuntu:~/piscine-go/test$ ./test
true
false
*
