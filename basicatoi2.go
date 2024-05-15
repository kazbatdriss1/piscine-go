package piscine

func BasicAtoi2(s string) int {
	cpt := -1
	rus := 0

	for range s {
		cpt++
	}
	for i := 0; i <= cpt; i++ {
		if s[i] >= 48 && s[i] <= 57 {
			rus *= 10
			rus += int(s[i]) - '0'
		} else if !(s[i] >= 48 && s[i] <= 57) {
			return 0
		}
	}
	return rus
}
