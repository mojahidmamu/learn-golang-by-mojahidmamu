package main

import "fmt"

func main() {
	// Start your code....

	defer fmt.Println("Clean up") // This will execute when the function exits, even if there's a panic
	panic("Something wrong")      // This will trigger a panic
}
