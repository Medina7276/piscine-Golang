package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {    //Errors (non possible values or overflows) will return 0.
		return 0
	}

	result := 1 // заранее пишем ответ задачи
	for i := 0; i < power; i++ { //в массиве мы умножаем наше число на себя столько раз пока i не будет равен power(3) 
		result *= nb //result = result * nb - это то же самое
	}
	return result
}
/*ITERATIVEPOWER.GO
Write an iterative function that returns the power of the int passed as parameter.
Negative powers will return 0. Overflows do not have to be dealt with.
student@ubuntu:~/piscine-go/test$ ./test
64 
*
