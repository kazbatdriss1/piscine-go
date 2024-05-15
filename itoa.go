package piscine

func Itoa(n int) string {
	rus := ""
	sng := ""
	if n == 0 {
		rus = "0"
	} else {
		if n < 0 {
			n *= -1
			sng = "-"
		}
		for n != 0 {
			rus = string(n%10+'0') + rus
			n = n / 10
		}
	}
	return sng + rus
}
