package main

import (
	"fmt"
)

func SortWordArr(a []string) {
	swp := ""
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] <= a[j] {
				swp = a[j]
				a[j] = a[i]
				a[i] = swp
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

/*
$ go run .
[1 2 3 A B C a b c]
$
*/
