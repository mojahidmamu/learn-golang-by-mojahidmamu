package main

import "fmt"

func main() {
	// Start your code....
	great := func(name string) { // Anonymous function assigned to variable
		fmt.Println("Hello", name)
	}

	great("Alice") // Call the anonymous function with argument

	// Demonstrate defer and panic
	defer fmt.Println("Clean up") // This will execute when the function exits, even if there's a panic
	panic("Something wrong")      // This will trigger a panic

}
