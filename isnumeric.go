package piscine

func IsNumeric(s string) bool {
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return false
		}
	}
	return true
}
