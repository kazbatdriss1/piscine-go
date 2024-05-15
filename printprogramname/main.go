package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lis := os.Args[0]
	for i, v := range lis {
		if i > 1 {
			z01.PrintRune(rune(v))
		}
	}
	z01.PrintRune('\n')
}
