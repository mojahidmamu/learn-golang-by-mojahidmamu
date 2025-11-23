package main

import "fmt"

func main() {
	// Start your code....
	x := 20 
	a := &x // address of x 
	*a = 50 // reassign value of address a	
	
	fmt.Println("This is value", x)
	fmt.Println("This is address", a)
	fmt.Println("value of address", *a)
	fmt.Println("reassign value", a)
	fmt.Println("reassign value", *a)
}
