package main

import "fmt"

func main() {
	// Start your code....

	// Internal memory example
	type Person struct {
		Name string
		Age  int
	}
	p1 := Person{Name: "Alice", Age: 30}
	p2 := p1 // p2 is a copy of p1
	p2.Age = 31

	fmt.Println("p1 Age:", p1.Age)
	fmt.Println("p2 Age:", p2.Age)

	// End your code....
}
