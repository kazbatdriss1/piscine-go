package piscine

func Split(s, sep string) []string {
	var rus []string
	max := 0
	for i := 0; i < len(s)-len(sep); {
		for j := i; j < len(s)-len(sep); j++ {
			if s[j:j+len(sep)] == sep {
				rus = append(rus, s[i:j])
				i = j + len(sep)
				max = j
			}
		}
		i++
	}
	rus = append(rus, s[max+len(sep):])
	return rus
}
