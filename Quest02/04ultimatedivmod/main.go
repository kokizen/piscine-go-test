package main

import (
	"fmt"
	"piscine-go-test/Quest02/04ultimatedivmod/piscine"
)

func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
