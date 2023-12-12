package main

import (
	"github.com/01-edu/z01"
)
 func main () {
	a := "Hello World!"
	for _, mot := range a {
		z01.PrintRune(mot)
	}
	z01.PrintRune('\n')
 }