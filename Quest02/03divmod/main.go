package main

import (
	"fmt"
	"piscine-go-test/Quest02/03divmod/piscine"
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
