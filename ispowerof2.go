package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		arg := os.Args[1]
		if int(arg[0])%2 == 0 {
			fmt.Println("true")
			return
		} else {
			fmt.Println("false")
			return
		}
	} else {
		return
	}
}

/*
$ go run . 2 | cat -e
true$
$ go run . 64
true
$ go run . 513
false
$ go run .
$ go run . 64 1024
$
*/
