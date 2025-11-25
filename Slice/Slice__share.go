package main

import "fmt"

func main() {
	// Start your code....
	arr := [5]int{10, 20, 30, 40, 50}
	s1 := arr[1:3]
	s2 := arr[2:]
	fmt.Println("Before", s1, s2)
	
	s1[1] = 999
	fmt.Println("After", s1, s2)
	fmt.Println("The array is:", arr)
}
