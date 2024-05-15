package piscine

func IterativeFactorial(nb int) int {
	rus := 1
	if nb < 0 || nb >= 21 {
		return 0
	}
	for i := nb; i >= 1; i-- {
		rus = rus * i
	}
	return rus
}
