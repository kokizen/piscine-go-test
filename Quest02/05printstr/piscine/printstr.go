package piscine

import "github.com/kokizen/z01"

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
