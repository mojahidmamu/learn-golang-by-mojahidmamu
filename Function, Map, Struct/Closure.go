package main

import "fmt"

func main() {
	// Start your code....
	// Closure example
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
	// End your code....
}
