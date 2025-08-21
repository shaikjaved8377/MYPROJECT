package example

import "fmt"

func Shop() {

	products := make(map[string]string)

	products["product1"] = "Onion"
	products["product2"] = "Oil"
	products["product3"] = "chicken"

	fmt.Println("products:", products)

	products["Product4"] = "Rice"
	fmt.Println("After adding product4:", products)

	products["product1"] = "potato"
	fmt.Println("After updating product1:", products)

	for i, k := range products {
		fmt.Println(i, k)
	}

}
