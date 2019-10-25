package piscine

func Sqrt(nb int) int {                 //nb=16 
	for i := 0; i <= nb; i++ {      // i=0 i <=16
		if i*i == nb {          // 0*0 != 16 и дальше идет иттерация, то есть уже i= 1, выолняется снова 0*1!=16 и так далее по увеличению i 
			return i
		}
	}

	return 0
}
/*SQRT.GO
Write a function that returns the square root of the int passed as parameter 
if that square root is a whole number. 
Otherwise it returns 0.
student@ubuntu:~/piscine-go/test$ ./test
2
0
*
