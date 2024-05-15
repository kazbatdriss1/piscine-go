package piscine

func Abort(a, b, c, d, e int) int {
	tb := []int{a, b, c, d, e}
	swp := 0
	for i := 0; i < len(tb); i++ {
		for j := 0; j < len(tb); j++ {
			if tb[i] <= tb[j] {
				swp = int(tb[j])
				tb[j] = int(tb[i])
				tb[i] = swp
			}
		}
	}

	return int(tb[(len(tb)-1)/2])
}
