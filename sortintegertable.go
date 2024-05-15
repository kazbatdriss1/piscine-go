package piscine

func SortIntegerTable(table []int) {
	ln := -1
	swp := 0
	for range table {
		ln++
	}
	for i := 0; i <= ln; i++ {
		for j := 0; j <= ln; j++ {
			if table[i] <= table[j] {
				swp = table[j]
				table[j] = table[i]
				table[i] = swp
			}
		}
	}
}
