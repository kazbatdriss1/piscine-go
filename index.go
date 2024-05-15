package piscine

func Index(s string, toFind string) int {
	A := []rune(s)
	B := []rune(toFind)
	lnA := len(A)
	lnB := len(B)
	for i := 0; i <= lnA-lnB; i++ {
		if toFind == s[i:i+lnB] {
			return i
		}
	}
	return -1
}
