package piscine

func strlen(str string) int {
	rus := -1
	for range str {
		rus++
	}
	return rus
}

func StrRev(s string) string {
	index := strlen(s)
	var rus string
	for i := index; i >= 0; i-- {
		rus += string(s[i])
	}
	return rus
}
