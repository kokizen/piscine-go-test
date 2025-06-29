package main

import (
	"fmt"
	"piscine-go-test/Quest02/07swap/piscine"
)

func main() {
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
