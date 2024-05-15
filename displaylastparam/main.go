package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lis := os.Args[1:]
	for i := len(lis) - 1; i < len(lis); i++ {
		for _, pr := range lis[i] {
			z01.PrintRune(rune(pr))
		}
	}
	z01.PrintRune('\n')
}
