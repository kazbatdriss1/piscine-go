package piscine

func RecursiveFactorial(nb int) int {
	rus := 1
	if nb < 0 || nb >= 21 {
		return 0
	}
	if nb > 0 {
		rus *= nb * RecursiveFactorial(nb-1)
		nb = nb - 1
	}

	return rus
}
