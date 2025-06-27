package main

import "github.com/kokizen/z01"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

}
