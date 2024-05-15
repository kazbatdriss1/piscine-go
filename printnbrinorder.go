package piscine

import (
	"github.com/01-edu/z01"
)

func remplitb(A []int) {
	lenN := len(A) - 1
	swp := 0
	for i := 0; i <= lenN; i++ {
		for j := 0; j <= lenN; j++ {
			if A[i] <= A[j] {
				swp = A[j]
				A[j] = A[i]
				A[i] = swp
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	A := []int{}
	if n == 0 {
		z01.PrintRune(rune(n) + '0')
	} else {
		for n != 0 {
			A = append(A, n%10)
			n = n / 10
		}
		remplitb(A)
		for _, h := range A {
			z01.PrintRune(rune(h) + '0')
		}
	}
}
