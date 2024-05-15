package piscine

func Capitalize(s string) string {
	rus := ""
	capnext := true
	for _, c := range s {
		if c >= '0' && c <= '9' {
			rus += string(c)
			capnext = false
		} else if c >= 'A' && c <= 'Z' {
			if capnext {
				rus += string(c)
			} else {
				rus += string(c + 32)
			}
			capnext = false
		} else if c >= 'a' && c <= 'z' {
			if capnext {
				rus += string(c - 32)
			} else {
				rus += string(c)
			}
			capnext = false
		} else {
			rus += string(c)
			capnext = true
		}
	}
	return rus
}

func Capitalizee(s string) string {
	rus := ""
	upper := true
	for _, pr := range s {
		if (pr >= '0' && pr <= '9') || (pr >= 'a' && pr <= 'z') || (pr >= 'A' && pr <= 'Z') {
			if upper {
				if pr >= 'a' && pr <= 'z' {
					rus += string(pr - 32)
				} else {
					rus += string(pr)
				}
			} else {
				if pr >= 'A' && pr <= 'Z' {
					rus += string(pr + 32)
				} else {
					rus += string(pr)
				}
			}
			upper = false
		} else {
			rus += string(pr)
			upper = true
		}
	}
	return rus
}
