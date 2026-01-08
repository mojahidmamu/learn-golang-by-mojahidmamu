package main

import "fmt"

func main() {
	// Start your code....
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1)) 
	copy(s2, s1)

	s2[0] = 126

	fmt.Println("s1:", s1) // [1 2 3]
	fmt.Println("s2:", s2) // 
}
