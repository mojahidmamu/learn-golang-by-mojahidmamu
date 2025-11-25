package main

import "fmt"

func main() {
	// Start your code....
	s := make([]int, 3, 5) // make(type, length, capacity)

	fmt.Println("S1 = ", s)
	fmt.Println("Len of s = ", len(s))
	fmt.Println("cap of s = ", cap(s))
}
	