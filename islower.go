package piscine

func IsLower(s string) bool {
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') {
			return false
		}
	}
	return true
}
