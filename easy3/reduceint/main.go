package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	for _, d := range strconv.Itoa(result) {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}
