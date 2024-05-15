package piscine

func IsAlpha(s string) bool {
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9')) {
			return false
		}
	}
	return true
}
