package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	fille := "quest8.txt"
	txt, _ := ioutil.ReadFile(fille)
	fmt.Print(string(txt))
}
