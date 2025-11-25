package main

import "fmt"

func main() {
	// Start your code....
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println(nums, len(nums), cap(nums))

	nums = append(nums, 4, 100, 200)
	fmt.Println(nums, len(nums), cap(nums))
}
		