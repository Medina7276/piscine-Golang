package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 1; j < len(table)+1; j++ {
			if table[j] > table[i] {
				table[i] = table[i] + table[j]
				table[j] = table[i] - table[j]
				table[i] = table[i] - table[j]
			}
		}
	}
}