package piscine

func FishAndChips(n int) string {
	rus := ""
	if n < 0 {
		rus = "number is negative"
		return rus
	} else {
		if (n%2 == 0) && (n%3 == 0) {
			rus = "fish and chips"
			return rus
		} else if n%2 == 0 {
			rus = "fish"
			return rus
		} else if n%3 == 0 {
			rus = "chips"
			return rus
		} else if !((n%2 == 0) && (n%3 == 0)) {
			rus = "non divisible"
			return rus
		}
	}
	return rus
}
