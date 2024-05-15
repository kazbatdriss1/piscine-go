package piscine

func Compare(a, b string) int {
	lena := len(a) - 1
	lenb := len(b) - 1
	bol := false
	rus := 0
	if a[0] < b[0] {
		return -1
	} else {
		for i := 0; i <= lena; i++ {
			for j := 0; j <= lenb; j++ {
				if a[i] == b[j] && a[lena] == b[lenb] && bol == true {
					bol = true
					rus = 0
				} else if a[i] == b[lenb] && bol == false {
					bol = true
					rus = 1
				} else {
					rus = 1
				}
			}
		}
		return rus
	}
}
