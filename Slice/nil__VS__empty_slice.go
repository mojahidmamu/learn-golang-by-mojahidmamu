package main

import "fmt"

func main() {
	// Start your code....
	var a []int  // nil slice
	b := []int{} // empty slice

	fmt.Println(a == nil)       // true
	fmt.Println(b == nil)       // false
	fmt.Println(len(a), len(b)) // 0 0
	fmt.Println(a, b)           // [] []

}
