package main

import (
	"fmt"
	"myproject/simplecalc"
)

func main() {
	fmt.Println("Hello, World!")

	a, b := 5, 3
	fmt.Println(simplecalc.Add(a, b))
	fmt.Println(simplecalc.Subtract(a, b))
	fmt.Println(simplecalc.Multiply(a, b))
	fmt.Println(simplecalc.Divide(a, b))
}
