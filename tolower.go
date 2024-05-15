package piscine

func ToLower(s string) string {
	lenS := len(s)
	var rus string
	for i := 0; i < lenS; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			rus += string(s[i] + 32)
		} else {
			rus += string(s[i])
		}
	}
	return rus
}
