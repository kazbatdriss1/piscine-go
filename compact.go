package piscine

func Compact(ptr *[]string) int {
	cpt := 0
	var str []string
	for _, p := range *ptr {
		if p != "" {
			str = append(str, p)
			cpt++
		}
	}
	*ptr = str
	return cpt
}
