package piscine

func Join(strs []string, sep string) string {
	var rus string
	for i := 0; i < len(strs); i++ {
		rus += strs[i]
		for j := 0; j < len(sep); j++ {
			if i < len(strs)-1 {
				rus += string(sep[j])
			}
		}
	}
	return rus
}
