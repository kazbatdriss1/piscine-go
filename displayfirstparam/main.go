package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lis := os.Args[1]
	for i := 0; i < len(lis); i++ {
		z01.PrintRune(rune(lis[i]))
	}
	z01.PrintRune('\n')
}
