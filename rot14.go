package piscine

func Rot14(s string) string {
	rus := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if s[i] >= 'A' && s[i] < 'M' {
				rus += string(s[i] + 14)
			} else {
				rus += string(s[i] - 12)
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			if s[i] >= 'a' && s[i] < 'm' {
				rus += string(s[i] + 14)
			} else {
				rus += string(s[i] - 12)
			}
		} else {
			rus += string(s[i])
		}
	}
	return rus
}

func Rot13(s string) string {
	rus := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'm' || s[i] >= 'A' && s[i] <= 'M' {
			rus += string(s[i] + 13)
		} else if s[i] >= 'n' && s[i] <= 'z' || s[i] >= 'N' && s[i] <= 'Z' {
			rus += string(s[i] - 13)
		} else {
			rus += string(s[i])
		}
	}
	return rus
}
