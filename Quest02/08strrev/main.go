package main

import (
	"fmt"
	"piscine-go-test/Quest02/08strrev/piscine"
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
