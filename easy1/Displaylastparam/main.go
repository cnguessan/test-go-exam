package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main () {
	if len(os.Args) == 1 {
		return
	}
	a := os.Args[len(os.Args) -1]
	for _, mot := range a {
		z01.PrintRune(mot)
	}
	z01.PrintRune('\n')
}