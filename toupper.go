package piscine

func ToUpper(s string) string {
	lenS := len(s)
	var rus string
	for i := 0; i < lenS; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			rus += string(s[i] - 32)
		} else {
			rus += string(s[i])
		}
	}
	return rus
}
