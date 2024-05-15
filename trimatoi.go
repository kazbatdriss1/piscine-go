package piscine

func TrimAtoi(s string) int {
	rus := 0
	isng := false
	sng := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' || (s[i] == '-') || (s[i] == '+') {
			if s[i] >= '0' && s[i] <= '9' {
				rus *= 10
				rus += int(s[i]) - '0'
				isng = true
			}
			if s[i] == '-' && isng == false {
				sng = -1
			}
		} else {
			continue
		}
	}
	return (rus * sng)
}
