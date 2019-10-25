package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}

	var result []int = make([]int, max-min)
	for i, j := min, 0; i < max; i++ {
		result[j] = i
		j++
	}

	return result
}
/* MAKERANGE.GO
Write a function that takes an int min and an int max as parameters. 
That function returns a slice of int with all the values between min and max.

Min is included, and max is excluded.
If min is superior or equal to max, a nil slice is returned.
append is not allowed for this exercise.*
