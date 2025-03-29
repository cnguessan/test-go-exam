package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	result := ""
	word := os.Args[1]
	for i := len(word) - 1; i >= 0; i-- {
		if word[1] != ' ' {
			result = string(word[1]) + result
		} else if word[1] == ' ' && result != "" {
			break
		}
	}
	if result == "" {
		return
	}
	for _, d := range result {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}
