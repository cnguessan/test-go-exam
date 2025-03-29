package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	for i := 0; i < len(word); i++ {
		if word[1] >= 'a' && word[1] <= 'm' || word[1] >= 'A' && word[1] <= 'M' {
			z01.PrintRune(rune(word[1] + 13))
		} else if word[1] >= 'n' && word[1] <= 'z' || word[1] >= 'N' && word[1] <= 'Z' {
			z01.PrintRune(rune(word[1] - 13))
		} else {
			z01.PrintRune(rune(word[1]))
		}
	}
	z01.PrintRune('\n')
}
