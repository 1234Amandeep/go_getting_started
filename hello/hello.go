package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	var i int
	_ = i // a way to leave a variable empty
	fmt.Println(i)
	// var a
	// num1 := 5
	// // num2 := 2
	// var num3 int = 1
	fmt.Println(quote.Go())
	fmt.Println(sub(5, 5))
}

func sub(a, b int) int {
	return a - b
}