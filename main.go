package main

import (
	"fmt"
	//"simplecalc"
)

func main() {
	i := 6
	j := &i
	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", &i)
	fmt.Println("Value of j :", j)
	fmt.Println("Value pointed to by j:", *j)
	*j = 10
	fmt.Println("New value of i:", i)
}

/*
func main() {
	day := 1 // write switch case statement to print day of the week

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}

/*
// write if else if else statements
func main() {
	a := 10
	if a > 0 {
		fmt.Println(a, "is positive")
	} else if a < 0 {
		fmt.Println(a, "is negative")
	} else {
		fmt.Println(a, "is zero")
	}
}

/*
// write if else statements
func main() {
	a := 10
	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}
}

/*
func main() {      //write for loo
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
*/

/*
func main() {
	fmt.Println("Hello, World!")

	a, b := 5, 3
	fmt.Println(simplecalc.Add(a, b))
	fmt.Println(simplecalc.Subtract(a, b))
	fmt.Println(simplecalc.Multiply(a, b))
	fmt.Println(simplecalc.Divide(a, b))
}
*/
