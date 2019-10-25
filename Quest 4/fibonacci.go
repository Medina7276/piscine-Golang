package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}

}
/*FIBONACCI.GO
Write an recursive function that returns the value of fibonacci sequence matching the index passed as parameter.
The first value is at index 0.
The sequence starts this way: 0, 1, 1, 2, 3 etc…
A negative index will return -1.
for is forbidden for this exercise.
student@ubuntu:~/piscine-go/test$ ./test
3
*
