package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	str := ""
	if len(args) == 1 {
		arg := os.Args[1]
		for i := 0; i < len(arg); i++ {
			if arg[i] >= 'a' && arg[i] <= 'z' {
				str += string(arg[i] - 32)
			} else if arg[i] >= 'A' && arg[i] <= 'Z' {
				str += string(arg[i] + 32)
			} else {
				str += string(arg[i])
			}
		}
		for _, pr := range str {
			z01.PrintRune(rune(pr))
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}

/*

$ go run . "SometHingS iS WronG"
sOMEThINGs Is wRONg
$ go run .
$
*/
