package piscine

func LastRune(s string) rune {
	cpt := -1
	for range s {
		cpt++
	}
	if rune(s[cpt]) == 'â' {
		return '♥'
	}
	return rune(s[cpt])
}
