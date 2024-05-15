package piscine

func AlphaCount(s string) int {
	max := len(s) - 1
	rus := 0
	for i := 0; i <= max; i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {
			rus++
		}
	}
	return rus
}
