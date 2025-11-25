package main

import "fmt"

func main() {
	// Start your code....
	s := []int{1, 2, 3}

	//    add in last:
	s = append(s, 4) // [1 2 3 4]

	// a lot at a time:
	s = append(s, 5, 6, 7, 8) // [1 2 3 4 5 6 7 8]

	// add another slice: 
	t := []int{10,20}
	s = append(s, t...)

	
}
