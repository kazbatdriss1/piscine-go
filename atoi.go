package piscine

func Atoi(s string) int {
	cpt := -1
	rus := 0
	sng := 1
	for range s {
		cpt++
	}
	for i := 0; i <= cpt; i++ {
		if !((s[i] >= 48 && s[i] <= 57) || (s[i] == 43) || (s[i] == 45)) {
			return 0
		} else if (s[i] >= 48 && s[i] <= 57) || (s[i] == 43) || (s[i] == 45) {

			if s[i] >= 48 && s[i] <= 57 {
				rus *= 10
				rus += int(s[i]) - '0'
			}
			if s[i] == 45 && i == 0 {
				sng = -1
			}
			if (s[i] == 45 && i != 0) || (s[i] == 43 && i != 0) {
				return 0
			}
		}
	}
	return (rus * sng)
}
