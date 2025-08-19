package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func main() {
	p := Person{
		Name:    "Javed",
		Age:     22,
		Address: Address{Street: "12355 devon St", City: "Chicago", State: "IL", Zip: "123"},
	}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:", p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)

	//array example
	a := [5]string{"Chicago", "Newyork", "Texas", "colorado", "Florida"}
	for i, v := range a {
		fmt.Println(i, v)
	}

	//slice we have 4 types-example
	//type-1 -So len and cap are same value
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)
	//type -2
	s = make([]int, 3, 5)
	s = append(s, 4, 5)
	s = append(s, 6)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// type -3
	s = make([]int, 3)
	s = append(s, 4)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)

	//type -4 slicing with an array
	arr := []int{1, 2, 3, 4, 5}
	s = arr[1:5]
	fmt.Println(s)

}

/*
// example for pointers
func main() {
	i := 6
	j := &i
	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", &i)
	fmt.Println("Value of j :", j)
	fmt.Println("Value pointed to by j:", *j)
	*j = 10
	fmt.Println("New value of i:", i)

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

	// write if else if else statements

	a := 10
	if a > 0 {
		fmt.Println(a, "is positive")
	} else if a < 0 {
		fmt.Println(a, "is negative")
	} else {
		fmt.Println(a, "is zero")
	}

	// write if else statements

	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}

	//write for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Hello, World!")

	a, b := 5, 3
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	if b != 0 {
		fmt.Println("Division:", a/b)
	} else {
		fmt.Println("Division by zero error")
	}
}
*/
