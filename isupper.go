package piscine

func IsUpper(s string) bool {
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		if !(s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}
	return true
}
