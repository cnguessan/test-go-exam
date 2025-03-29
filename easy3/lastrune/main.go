package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	word := []rune(s)
	return word[len(s)-1]
}

func main() {
	z01.PrintRune(FirstRune("HelloP"))
	z01.PrintRune(FirstRune("SalutO"))
	z01.PrintRune(FirstRune("OlN"))
	z01.PrintRune(FirstRune("OlY"))
	z01.PrintRune('\n')
}
