package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main () {
	if len(os.Args) == 3 {
		for i , j := 0, 0 ; i < len(os.Args[1]) && j < len(os.Args[2]) ; {
			if os.Args[1][i] == os.Args[2][j]{
				i++
				if i == len(os.Args[1]){
					for _, d := range os.Args[1]{
						z01.PrintRune(d)
					}
					z01.PrintRune(10)
				}
				
			}
			j++
			if j == len(os.Args[2]){
				return
			}
		}
	}
}