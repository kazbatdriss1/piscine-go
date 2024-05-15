package piscine

func SplitWhiteSpaces(s string) []string {
	var rus []string
	str := ""
	for _, v := range s {
		if v != ' ' && v != '\n' && v != '\t' {
			str += string(v)
		} else {
			if str != "" {
				rus = append(rus, str)
				str = ""
			}
		}
	}
	rus = append(rus, str)
	return rus
}
