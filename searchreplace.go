package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	rus := ""
	if len(args) != 3 {
		return
	} else {
		a := os.Args[1]
		b := os.Args[2]
		c := os.Args[3]
		for _, val := range a {
			if val == rune(b[0]) {
				rus += string(c[0])
			} else {
				rus += string(val)
			}
		}
		fmt.Println(rus)
	}
}

/*
$ go run . "hella there" "a" "o"
hello there
$ go run . "hallo thara" "a" "e"
hello there
$ go run . "abcd" "z" "l"
abcd
$ go run . "something" "a" "o" "b" "c"
$
*/
