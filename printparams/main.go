package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := 1
	for i := 1; i < len(os.Args); i++ {
		for j := 0; j < len(os.Args[i]); j++ {
			r := []rune(os.Args[i])
			if i > a {
				z01.PrintRune(10)
			}
			z01.PrintRune(r[j])
			a = i
		}
	}
	z01.PrintRune(10)
}
