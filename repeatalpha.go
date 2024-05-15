package main

import (
	"fmt"
	"os"
)

func findIndex(val rune) int {
	low := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upp := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for i := 0; i < len(low); i++ {
		if val == low[i] {
			return i
		} else if val == upp[i] {
			return i
		}
	}
	return 0
}

func main() {
	args := os.Args[1:]
	index := 0
	if len(args) != 1 {
		return
	} else {
		arg := os.Args[1]
		for i := 0; i < len(arg); i++ {
			index = findIndex(rune(arg[i]))
			for j := 0; j <= index; j++ {
				fmt.Print(string(arg[i]))
			}
		}
	}
}

/*
$ go run . abc | cat -e
abbccc
$ go run . Choumi. | cat -e
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
Chhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.
$ go run . "abacadaba 01!" | cat -e
abbacccaddddabba 01!$
$ go run .
$ go run . "" | cat -e
$
$
*/
