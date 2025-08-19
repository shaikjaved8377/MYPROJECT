package main

import "fmt"

// using maps concepts
func main() {

	products := make(map[string]string)

	products["product1"] = "chocolate"
	products["product2"] = "chips"
	products["product3"] = "soda"

	fmt.Println("Products:", products)

	// now im going to Add the product
	products["product4"] = "cookies"
	fmt.Println("after adding new Product:", products)

	//now im going to update the product
	products["product1"] = "chocolate bar"
	fmt.Println("after updating product1:", products)

	//now im going to delete the product
	delete(products, "product2")
	fmt.Println("after deleting product2:", products)

	// now im going to check if the product exists
	if value, exists := products["product3"]; exists {
		fmt.Println("product3 exists with value:", value)
	} else {
		fmt.Println("product3 does not exist")
	}
	// iterating over the map
	for key, value := range products {
		fmt.Println(key, value)
	}

	// pointers with functions stocks
	applePrice := 10

	fmt.Println("Before stock increase:", applePrice)

	StockPriceInc(&applePrice)
	fmt.Println("After stock increase:", applePrice)

	//now im using methods concepts
	P := Product{
		Name:  "Macbook Pro",
		Price: 1000.00,
	}

	fmt.Println("Product Name:", P.Name)
	fmt.Println("Product Actual price:", P.Price)
	fmt.Println("Product After discount:", P.ApplyDiscount())

}

func StockPriceInc(appleprice *int) {
	*appleprice = *appleprice + 20

}

type Product struct {
	Name  string
	Price float64
}

func (p Product) ApplyDiscount() float64 {
	return p.Price * 0.9
}

/*



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
