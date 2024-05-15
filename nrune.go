package piscine

func NRune(s string, n int) rune {
	cpt := -1
	for range s {
		cpt++
	}
	if cpt < n-1 || n-1 < 0 {
		return 0
	} else {
		if rune(s[n-1]) == 'â' {
			return '♥'
		}
		return rune(s[n-1])
	}
}
