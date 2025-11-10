package main

import "fmt"

func main() {
	// Start your code....
	arr := [5]int{1, 2}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("Sum:", sum)
	fmt.Println(arr)
	// End your code.
}
