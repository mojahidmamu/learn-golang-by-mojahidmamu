package main

import "fmt"

var arr2 = [3]string{"I" , "Love", "You"}

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
