package piscine

func StrLen(s string) int {
	rus := 0
	for _, i := range s {
		i++
		rus++
	}
	return rus
}
