package main

import (
	"fmt"
	"piscine-go-test/Quest04/11isnumeric/piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("010203"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
}
