package main

import "fmt"

func main() {
	// Start your code....
	s1 := []int{1, 2, 3, 4, 5}

	a := s1[1:3] //[2 3]
	b := s1[2:]  //[3 4 5]

	fmt.Println(a, "length", len(a), "capacity", cap(a))
	fmt.Println(b, "length", len(b), "capacity", cap(b))
}				