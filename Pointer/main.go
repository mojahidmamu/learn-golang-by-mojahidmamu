package main

import "fmt"

func main() {
	// Start your code....
	x := 20
	a := &x
	fmt.Println("This is value", x)
	fmt.Println("This is address", a)
	fmt.Println("value of address", *a)
}
