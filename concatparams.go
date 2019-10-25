package piscine

func ConcatParams(args []string) string {
	var result string = args[0]
	for _, s := range args[1:] {
		result += "\n" + s
	}

	return result
}
/* CONCATPARAMS.GO

Write a function that takes the arguments
received in parameters and returns them as a string.
The arguments must be separated by a \n.*
