package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if f(a[0], a[1]) >= 0 {
			for i := 1; i < len(a); i++ {
				if f(a[i-1], a[i]) < 0 {
					return false
				}
			}
		}
		if f(a[0], a[1]) <= 0 {
			for i := 1; i < len(a); i++ {
				if f(a[i-1], a[i]) > 0 {
					return false
				}
			}
		}
	}
	return true
}
