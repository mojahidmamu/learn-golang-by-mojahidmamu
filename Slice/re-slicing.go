package main

import "fmt"

func main() {
	// Start your code....
	s1 := []int{1, 2, 3, 4, 5}

	a := s1[1:3]
	// b := s1[2:]

	fmt.Println(a, len(a), cap(a))
}
