package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	str := ""
	// rus := ""
	bol := false
	if len(args) == 1 {
		arg := os.Args[1]
		for i := 0; i < len(arg); i++ {
			if arg[i] != ' ' {
				str += string(arg[i])
				bol = false
			} else if !bol && str != "" {
				if str[len(str)-1] != ' ' {
					str += "   "
				}
				bol = true
			}
		}
		bol = false
		/*for i := len(str) - 1; i >= 0; i-- {
			if str[i] != ' ' {
				rus += string(str[i])
				bol = false
			} else if !bol && rus != "" {
				if rus[len(rus)-1] != ' ' {
					rus += "   "
				}
				bol = true
			}
		}*/
		fmt.Println(str)
	} else {
		return
	}
}

/*
$ go run . "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
$ go run . "   only  it's harder   " | cat -e
only   it's   harder$
$ go run . " how funny it is" "did you  hear, Mathilde ?" | cat -e
$ go run .
$
*/
