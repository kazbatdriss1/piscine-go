package piscine

func AppendRange(min, max int) []int {
	var rus []int
	if max > min {
		for i := min; i < max; i++ {
			rus = append(rus, i)
		}
	}
	return rus
}
