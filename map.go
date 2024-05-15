package piscine

func Map(f func(int) bool, a []int) []bool {
	rus := make([]bool, len(a))
	for h, fr := range a {
		rus[h] = f(fr)
	}
	return rus
}
