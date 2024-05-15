package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

/*
   48        0
   49        1
   50        2
   51        3
   52        4
   53        5
   54        6
   55        7
   56        8
   57        9
*/
func printStr(s rune) {
	z01.PrintRune(s)
}

func findint(inp int) {
	ax := '0'
	ay := '0'
	for i := 0; i < inp%10; i++ {
		ax++
	}
	for j := 0; j < inp/10; j++ {
		ay++
	}
	printStr(ay)
	printStr(ax)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	printStr('x')
	printStr(' ')
	printStr('=')
	printStr(' ')
	findint(points.x)
	printStr(',')
	printStr(' ')
	printStr('y')
	printStr(' ')
	printStr('=')
	printStr(' ')
	findint(points.y)
	printStr('\n')
}
