package piscine

func Unmatch(a []int) int {
	for i, pr := range a {
		found := 1
		for j, ph := range a {
			if i != j && pr == ph {
				found++
			}
		}
		if found%2 != 0 {
			return pr
		}
	}
	return -1
}
