package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lis := os.Args[1:]
	for i := 0; i < len(lis); i++ {
		for _, v := range lis[i] {
			z01.PrintRune(rune(v))
		}
		z01.PrintRune('\n')
	}
}
