package main

import "fmt"

func multiple(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

func main() {
	// Start your code....
	double := multiple(2)
	times3 := multiple(3)
	result := times3(4)
	fmt.Println(double(5)) // Output: 10
	fmt.Println(result)     // Output: 12
}
