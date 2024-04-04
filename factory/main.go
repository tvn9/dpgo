package main

import (
	"factory/products"
	"fmt"
	"time"
)

func main() {
	// Create a new product
	tundra := products.New()

	// Print out the product info
	fmt.Printf("Product Name: %v\n", tundra.ProductName)
	fmt.Printf("Created at: %v\n", tundra.CreatedAt)
	fmt.Printf("Updated at: %v\n", tundra.UpdatedAt)

	fmt.Println("After 5 seconds, update the product name")
	time.Sleep(5 * time.Second)

	// update product name
	func(t *products.Product) {
		t.ProductName = "Toyota Tundra"
		t.UpdatedAt = time.Now()
	}(tundra)

	// Print out the tundra object
	fmt.Printf("Product Name: %v\n", tundra.ProductName)
	fmt.Printf("Created at: %v\n", tundra.CreatedAt)
	fmt.Printf("Updated at: %v\n", tundra.UpdatedAt)
}
