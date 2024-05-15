package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lis := os.Args[1:]
	var swp string
	for i := 0; i < len(lis); i++ {
		for j := 0; j < len(lis); j++ {
			if lis[i] < lis[j] {
				swp = lis[j]
				lis[j] = lis[i]
				lis[i] = swp
			}
		}
	}
	for i := 0; i < len(lis); i++ {
		for _, v := range lis[i] {
			z01.PrintRune(rune(v))
		}
		z01.PrintRune('\n')
	}
}
