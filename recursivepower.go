package piscine

func RecursivePower(nb int, power int) int {
	rus := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	rus *= nb * RecursivePower(nb, power-1)
	return rus
}
