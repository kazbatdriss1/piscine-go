package main

import (
	"fmt"
)

func Atoi(s string) int {
	rus := 0
	sng := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' || s[i] == '+' || s[i] == '-' {
			if s[i] >= '0' && s[i] <= '9' {
				rus = rus * 10
				rus = rus + int(s[i]-'0')
			}
			if s[i] == '-' {
				sng = -1
			}
			if (s[i] == '-' && i != 0) || (s[i] == '+' && i != 0) {
				return 0
			}
		} else {
			return 0
		}
	}
	return (rus * sng)
}

func max(tb []int) int {
	//[91, 2, 3, 7, 8, 9,-1,90,777,0]
	sw := 0
	max := 0
	for i := 0; i < len(tb); i++ {
		for j := 0; j < len(tb); j++ {
			if tb[i] < tb[j] {
				sw = tb[j]
				tb[j] = tb[i]
				tb[i] = sw
			}
		}
	}
	max = int(tb[len(tb)-1])
	return max
}

func min(tb []int) int {
	//[91, 2, 3, 7, 8, 9,-1,90,777,0]
	sw := 0
	min := 0
	for i := 0; i < len(tb); i++ {
		for j := 0; j < len(tb); j++ {
			if tb[i] < tb[j] {
				sw = tb[j]
				tb[j] = tb[i]
				tb[i] = sw
			}
		}
	}
	min = int(tb[0])
	return min
}

func Itoa(n int) string {
	rus := ""
	sng := true
	if n == 0 {
		rus = "0"
	} else {
		if n < 0 {
			n = n * -1
			sng = false
		}
		for n != 0 {
			rus = string(n%10+'0') + rus
			n = n / 10
		}
		if !sng {
			rus = "-" + rus
		}
	}
	return rus
}

func main() {
	tb := []int{91, 2, 3, 7, 8, 9, -1, 90, 777, 0}
	fmt.Println(Itoa(12625145))  // 12345
	fmt.Println(Itoa(0))         // 0
	fmt.Println(Itoa(-1234))     //-1234
	fmt.Println(Itoa(987654321)) // 987654321
	fmt.Println(max(tb))
	fmt.Println(min(tb))
}
