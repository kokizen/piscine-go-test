package main

import (
	"fmt"
	"piscine-go-test/Quest04/05alphacount/piscine"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
