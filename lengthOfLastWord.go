package piscine

import "fmt"

func lengthOfLastWord(s string) {
	lenS := len(s)
	for i := lenS - 1; i >= 0; {
		for i >= 0 && string(s[i]) == " " {
			i--
		}
		if i < 0 {
			return
		}
		for i >= 0 && string(s[i]) != " " {
			fmt.Print(string(s[i]))
			i--
		}
	}
}
