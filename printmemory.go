package main

import "fmt"

func PrintMemory(arr [10]byte) {
	hexBase := "0123456789abcdef"
	result := ""
	for _, n := range arr {
		for n > 0 {
			mod := n % 16
			n = n / 16
			result = string(hexBase[mod]) + result
		}
	}

	for i := 0; i < len(result); i += 8 {
		// fmt.Printf("%08s ", result[i:i+8])
		fmt.Println(result[i : i+8])
		fmt.Println()
	}

	fmt.Println()

	for _, v := range arr {
		if v > 32 {
			fmt.Printf("%c", v)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
