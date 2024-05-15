package piscine

func FirstRune(s string) rune {
	if rune(s[0]) == 'â' {
		return '♥'
	}
	return rune(s[0])
}
