package main

import (
	"factory/products"
	"fmt"
	"time"
)

func main() {
	tundra := products.New()

	fmt.Printf("Product Name: %v\n", tundra.ProductName)
	fmt.Printf("Created at: %v\n", tundra.CreatedAt)
	fmt.Printf("Updated at: %v\n", tundra.UpdatedAt)

	func(t *products.Product) {
		fmt.Println("After 4 seconds, update the product name")
		time.Sleep(4 * time.Second)

		t.ProductName = "Totota Tundra"
		t.UpdatedAt = time.Now()

	}(tundra)

	fmt.Printf("Product Name: %v\n", tundra.ProductName)
	fmt.Printf("Created at: %v\n", tundra.CreatedAt)
	fmt.Printf("Updated at: %v\n", tundra.UpdatedAt)
}
