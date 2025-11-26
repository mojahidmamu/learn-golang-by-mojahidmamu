package main

import "fmt"

// variadic_function: 
func variadic(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	// Start your code....
	variadic(10, 20, 30, 40, 50, 40, 60, 50)
}
