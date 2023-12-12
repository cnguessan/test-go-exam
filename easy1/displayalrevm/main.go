package main

import "github.com/01-edu/z01"

func main () {
	sentence := "zYxWvUtSrQpOnMlKjIhGfEdCbA"
	for _, mot := range sentence {
		z01.PrintRune(mot)
	}
	z01.PrintRune(10)
}