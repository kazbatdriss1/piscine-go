package piscine

func MakeRange(min, max int) []int {
	var rus []int
	if max > min {
		rus = make([]int, max-min)
		for i := range rus {
			rus[i] = min + i
		}
	}
	return rus
}
