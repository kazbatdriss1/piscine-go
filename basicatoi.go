package piscine

func strle(str string) int {
	rus := -1
	for range str {
		rus++
	}
	return rus
}

func BasicAtoi(s string) int {
	stlen := strle(s)
	rus := 0
	for i := 0; i <= stlen; i++ {
		if s[i] >= 48 && s[i] <= 57 {
			rus *= 10
			rus += int(s[i]) - '0'
		} else if !(s[i] >= 48 && s[i] <= 57) {
			return 0
		}
	}
	return rus
}
