package main

import (
	"fmt"
	"piscine-go-test/Quest02/02ultimatepointone/piscine"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
