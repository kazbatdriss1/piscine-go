package main

import "fmt"

func Max(a []int) int {
	swp := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] <= a[j] {
				swp = a[j]
				a[j] = a[i]
				a[i] = swp
			}
		}
	}
	return int(a[len(a)-1])
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
